// Initialize and start our TODO micro-service application using the Conf2 system
// to load configuration and start management port

package main

import (
	"flag"
	"fmt"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
	"os"
	"time"
)

// We load from a local config file for simplicity, but same exact data can come
// over network on management port in any accepted format.
var configFileName = flag.String("config", "", "Configuration file")

func main() {
	flag.Parse()

	if len(*configFileName) == 0 {
		fmt.Fprint(os.Stderr, "Required 'config' parameter missing\n")
		os.Exit(-1)
	}

	// Most applications have a common app service from which you can access
	// all other services and data structures
	app := &App{}
	configFile, err := os.Open(*configFileName)
	if err != nil {
		panic("Error loading config " + err.Error())
	}

	// Read json, but you can implement reader in any format you want
	// your reader will be passed schema to validate node.
	config := node.NewJsonReader(configFile).Node()

	model, err := yang.LoadModule(yang.YangPath(), "todo")
	if err != nil {
		panic("Error loading TODO YANG " + err.Error())
	}

	api := &ApiHandler{}
	b := node.NewBrowser2(model, api.Manage(app))

	// load the config into empty app system.  Well designed api will not
	// distinguish config loading from management calls post operation
	if err := b.Root().UpsertFrom(config).LastErr; err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// start any main thread to keep app from exiting
	api.Restconf.Listen()
}

// APPLICATION
type App struct {
	todos map[string]*Task
}

type Status int

const (
	StatusTodo Status = iota
	StatusDone
)

type Task struct {
	Id          string
	Summary     string
	Status      Status
	Description string
	DueDate     time.Duration
	Keywords    []string
	timer       *time.Timer
}

// MANAGEMENT
type ApiHandler struct {
	Restconf *restconf.Service
}

func (api *ApiHandler) Manage(app *App) node.Node {
	s := &node.MyNode{}
	s.OnSelect = func(r node.ContainerRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "restconf":
			if r.New {
				api.Restconf = restconf.NewService(yang.YangPath(), r.Selection.Browser)
			}
			if api.Restconf != nil {
				return restconf.Api{}.Manage(api.Restconf), nil
			}
		case "todos":
			if r.New {
				app.todos = make(map[string]*Task)
			}
			if app.todos != nil {
				return api.Todos(app.todos), nil
			}
		}
		return nil, nil
	}
	return s
}

func (api *ApiHandler) Todos(todos map[string]*Task) node.Node {
	index := node.NewIndex(todos)
	return &node.MyNode{
		OnNext: func(r node.ListRequest) (node.Node, []*node.Value, error) {
			key := r.Key
			var id string
			var task *Task
			if key != nil {
				id = key[0].Str
			}
			if r.New {
				task = &Task{}
				todos[id] = task
			} else if key != nil {
				task = todos[id]
			} else {
				kv := index.NextKey(r.Row)
				id := kv.String()
				key = node.SetValues(r.Meta.KeyMeta(), id)
				task = todos[id]
			}
			if task != nil {
				return api.Todo(id, todos, task), key, nil
			}
			return nil, nil, nil
		},
		OnEvent: func(s node.Selection, e node.Event) error {
			switch e.Type {
			case node.REMOVE_LIST_ITEM:
				delete(todos, e.Src.Key()[0].Str)
			}
			return nil
		},
	}
}

func (api *ApiHandler) Todo(id string, todos map[string]*Task, task *Task) node.Node {
	return &node.Extend{
		Node: node.MarshalContainer(task),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "id":
				if r.Write {
					delete(todos, task.Id)
					task.Id = hnd.Val.Str
					todos[task.Id] = task
				} else {
					hnd.Val = &node.Value{Str:task.Id}
				}
			case "dueDate":
				if r.Write {
					task.DueDate = time.Duration(hnd.Val.Int64)
					if task.timer != nil {
						task.timer.Reset(task.DueDate)
					}
				} else {
					hnd.Val = &node.Value{Int64: int64(task.DueDate)}
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnEvent: func(p node.Node, s node.Selection, e node.Event) error {
			switch e {
			// This is what i want to change timers after all fields have been updated
			//		case data.UPDATE:
			//
			case node.NEW:
				if task.Status != StatusDone {
					task.timer = time.NewTimer(task.DueDate)
				}
			case node.DELETE:
				if task.timer != nil {
					task.timer.Stop()
				}
			}
			return p.Event(s, e)
		},
	}

}
