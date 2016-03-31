package node
import (
	"github.com/c2g/meta"
	"errors"
)

type Config struct {
	OnExtend ExtendConfigFunc
}

func ConfigNode(operational Node, config Node) Node {
	return Config{}.Node(operational, config)
}

type ExtendConfigFunc func(e Config, sel *Selection, m meta.MetaList, operational Node, config Node) (Node, error)

func (self Config) Node(operational Node, config Node) Node {
	return &MyNode{
		Label: "Persist",
		OnSelect: func(r ContainerRequest) (Node, error) {
			operChild, err := operational.Select(r)
			if err != nil || operChild == nil {
				return nil, err
			}
			if ! r.Selection.IsConfig(r.Meta) {
				return operChild, nil
			}
			configChild, storeErr := config.Select(r)
			if storeErr != nil {
				return nil, err
			}
			if configChild == nil && ! r.New {
				r.New = true
				configChild, storeErr = config.Select(r)
				if storeErr != nil {
					return nil, err
				}
				if configChild == nil {
					return nil, errors.New("Could not create storage node for " + r.Selection.String())
				}
			}
			if self.OnExtend != nil {
				if n, err := self.OnExtend(self, r.Selection, r.Meta, operChild, configChild); n != nil || err != nil {
					return n, err
				}
			}
			return self.Node(operChild, configChild), nil
		},
		OnNext: func(r ListRequest) (next Node, key []*Value, err error) {
			var operChild, configChild Node
			if operChild, key, err = operational.Next(r); err != nil || operChild == nil {
				return
			}
			if ! r.Selection.IsConfig(r.Meta) {
				return operChild, key, nil
			}
			r.Key = key
			if configChild, _, err = config.Next(r); err != nil {
				return
			}
			if configChild == nil && ! r.New {
				r.New = true
				if configChild, _, err = config.Next(r); err != nil {
					return
				}
				if configChild == nil {
					return nil, nil, errors.New("Could not create storage node for " + r.Selection.String())
				}
			}
			if self.OnExtend != nil {
				if n, err := self.OnExtend(self, r.Selection, r.Meta, operChild, configChild); n != nil || err != nil{
					return n, key, err
				}
			}
			return self.Node(operChild, configChild), key, nil
		},
		OnWrite: func(r FieldRequest, val *Value) error {
			if err := operational.Write(r, val); err != nil {
				return err
			}
			return config.Write(r, val)
		},
		OnEvent: func(sel *Selection, e Event) error {
			if err := operational.Event(sel, e); err != nil {
				return err
			}
			if err := config.Event(sel, e); err != nil {
				return err
			}
			return nil
		},
		OnRead : func(r FieldRequest) (*Value, error) {
			if r.Meta.(meta.HasDetails).Details().Config(r.Selection.Path()) {
				return config.Read(r)
			}
			return operational.Read(r)
		},
		OnChoose : operational.Choose,
		OnAction : operational.Action,
		OnPeek: operational.Peek,
	}
}

//func Config(operational Node, config Node) Node {
//	return ConfigNode{}.Node(operational, config)
//}
