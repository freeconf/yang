package main

import (
	"time"
	"data"
	"schema"
)


func (todos *Todos) Manage() data.Node {
	s := &data.MyNode{}
	var index int
	s.OnNext = func(sel *data.Selection, meta *schema.List, new bool, key []*schema.Value, first bool) (data.Node, error) {
		var task *Task
		if len(key) > 0 {
			index = key[0].Int
			task = todos.Tasks[index]
		} else if new {
			task = &Task{}
			index = len(todos.Tasks)
			if todos.Tasks == nil {
				todos.Tasks = []*Task{task}
			} else {
				todos.Tasks = append(todos.Tasks, task)
			}
		} else {
			if first {
				index = 0
			} else {
				index++
			}
			if index < len(todos.Tasks) {
				task = todos.Tasks[index]
				keyMeta := meta.KeyMeta()[0]
				sel.State.SetKey([]*schema.Value{&schema.Value{Type:keyMeta.GetDataType(), Int:index}})
			}
		}
		if task != nil {
			return task.Select(index), nil
		}
		return nil, nil
	}
	return s
}


func (task *Task) Select(id int) data.Node {
	s := &data.MyNode{}
	s.OnRead = func(sel *data.Selection, meta schema.HasDataType) (*schema.Value, error) {
		switch meta.GetIdent() {
		case "id":
			return  &schema.Value{Int:id}, nil
		case "dueDate":
			return &schema.Value{Int64:int64(task.DueDate)}, nil
		}
		return data.ReadField(meta, task)
	}
	s.OnWrite = func(sel *data.Selection, meta schema.HasDataType, v *schema.Value) error {
		switch meta.GetIdent() {
		case "id":
			// Not allowed
		case "dueDate":
			task.DueDate = time.Duration(v.Int64)
			if task.timer != nil {
				task.timer.Reset(task.DueDate)
			}
		default:
			return data.WriteField(meta, task, v)
		}
		return nil
	}
	s.OnEvent = func(sel *data.Selection, e data.Event) error {
		switch e {
// This is what i want to change timers after all fields have been updated
//		case data.UPDATE:
//
		case data.NEW:
			if task.Status != StatusDone {
				task.timer = time.NewTimer(task.DueDate)
			}
		case data.DELETE:
			if task.timer != nil {
				task.timer.Stop()
			}
		}
		return nil
	}

	return s
}