// +build ignore

package main

// Initialize and start our Todo micro-service application with C2Stack for
// RESTful based management

// To run:
//   cp todo{-sample,}.json && \
//     YANGPATH=.:../../../../../../../etc/yang/ \
//     go run ./main.go -config todo-sample.json
//
import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
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

	b := node.NewBrowser(model, ManageNode(app))

	// load the config into empty app system.  Well designed api will not
	// distinguish config loading from management calls post operation
	if err := b.Root().UpsertFrom(config).LastErr; err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// sleep forever
	select {}
}

// APPLICATION
type App struct {
	todos    map[string]*Task
	Restconf *restconf.Management
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

func (task *Task) Init() {
	if task.Status != StatusDone {
		task.timer = time.NewTimer(task.DueDate)
	}
}

func (task *Task) Deinit() {
	if task.timer != nil {
		task.timer.Stop()
	}
}

// MANAGEMENT
func ManageNode(app *App) node.Node {
	s := &nodes.Basic{}
	s.OnChild = func(r node.ChildRequest) (node.Node, error) {
		switch r.Meta.GetIdent() {
		case "restconf":
			if r.New {
				app.Restconf = restconf.NewManagement(yang.YangPath(), r.Selection.Browser)
			}
			if app.Restconf != nil {
				return restconf.ServiceNode(app.Restconf), nil
			}
		case "todos":
			if r.New {
				app.todos = make(map[string]*Task)
			}
			if app.todos != nil {
				return TodosNode(app.todos), nil
			}
		}
		return nil, nil
	}
	return s
}

func TodosNode(todos map[string]*Task) node.Node {
	index := node.NewIndex(todos)
	return &nodes.Basic{
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
				if r.Delete {
					delete(todos, id)
				} else {
					task = todos[id]
				}
			} else {
				kv := index.NextKey(r.Row)
				id := kv.String()
				key = node.SetValues(r.Meta.KeyMeta(), id)
				task = todos[id]
			}
			if task != nil {
				return TodoNode(task, todos), key, nil
			}
			return nil, nil, nil
		},
	}
}

func TodoNode(task *Task, todos map[string]*Task) node.Node {
	originalDueDate := task.DueDate
	return &nodes.Extend{
		Node: nodes.Reflect(task),
		OnField: func(p node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.GetIdent() {
			case "dueDate":
				if r.Write {
					task.DueDate = time.Duration(hnd.Val.Int64)
				} else {
					hnd.Val = &node.Value{Int64: int64(task.DueDate)}
				}
			case "id":
				if r.Write {
					old := task.Id
					task.Id = hnd.Val.String
					delete(old, todos)
					todos[task.Id] = task
				} else {
					hnd.Val = &node.Value{String: task.Id}
				}
			default:
				return p.Field(r, hnd)
			}
			return nil
		},
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if r.New {
				task.Init()
			} else if originalDueDate != task.DueDate {
				// This is where you can "watch" for changes to specific properties that
				// special handling
				task.timer.Reset(task.DueDate)
			}
			return p.EndEdit(r)
		},
		OnDelete: func(p node.Node, r node.NodeRequest) error {
			task.Deinit()
			return p.Delete(r)
		},
	}
}
