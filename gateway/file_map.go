package gateway

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/freeconf/gconf/c2"
	"github.com/freeconf/gconf/meta/yang"

	"github.com/freeconf/gconf/device"
	"github.com/freeconf/gconf/meta"

	"github.com/freeconf/gconf/node"
	"github.com/freeconf/gconf/nodes"
)

// Store all data in simple files.  Normally you would save this to a highly
// available, distributed service like a database.
type FileStore struct {
	VarDir     string
	ids        []string
	locations  Registrar
	listeners  *list.List
	southbound []device.ProtocolHandler
}

func NewFileStore(registrar Registrar, varDir string) *FileStore {
	return &FileStore{
		VarDir:    varDir,
		locations: registrar,
		listeners: list.New(),
	}
}

func (self *FileStore) AddProtocolHandler(h device.ProtocolHandler) {
	self.southbound = append(self.southbound, h)
}

func (self *FileStore) fname(deviceId string, module string) (string, string) {
	dir := self.deviceDir(deviceId)
	return fmt.Sprintf("%s/%s.json", dir, module), dir
}

func (self *FileStore) deviceDir(deviceId string) string {
	return fmt.Sprintf("%s/%s", self.configDir(), deviceId)
}

func (self *FileStore) configDir() string {
	return fmt.Sprintf("%s/%s", self.VarDir, "config")
}

func (self *FileStore) schemaDir() string {
	return fmt.Sprintf("%s/yang", self.VarDir)
}

func (self *FileStore) uiDir() string {
	return fmt.Sprintf("%s/web", self.VarDir)
}

func (self *FileStore) mkdir(s string) string {
	if err := os.MkdirAll(s, 0755); err != nil {
		panic(err)
	}
	return s
}

func (self *FileStore) NthDeviceId(index int) string {
	return self.deviceIds()[index]
}

func (self *FileStore) deviceIds() []string {
	if self.ids == nil {
		files, err := ioutil.ReadDir(self.configDir())
		if err != nil {
			return []string{}
		}
		ids := make([]string, len(files))
		for i, f := range files {
			ids[i] = f.Name()
		}
		self.ids = ids
	}
	return self.ids
}

func (self *FileStore) Device(deviceId string) (device.Device, error) {
	operDevice, err := self.operational(deviceId)
	if err != nil {
		return nil, err
	}
	var ypath, uipath meta.StreamSource
	if operDevice == nil {
		if !self.exists(deviceId) {
			return nil, nil
		}
		ypath = &meta.FileStreamSource{Root: self.schemaDir()}
		uipath = &meta.FileStreamSource{Root: self.uiDir()}
	} else {
		self.mkdir(self.schemaDir())
		self.mkdir(self.uiDir())
		ypath = meta.CacheSource{
			Local:    &meta.FileStreamSource{Root: self.schemaDir()},
			Upstream: operDevice.SchemaSource(),
		}
		uipath = meta.CacheSource{
			Local:    &meta.FileStreamSource{Root: self.uiDir()},
			Upstream: operDevice.SchemaSource(),
		}
	}
	d := device.NewWithUi(ypath, uipath)
	for _, moduleName := range self.modules(deviceId, operDevice) {
		m, err := yang.LoadModule(ypath, moduleName)
		if err != nil {
			panic(moduleName)
			return nil, err
		}
		var oper node.Node
		if operDevice != nil {
			o, err := operDevice.Browser(moduleName)
			if err != nil {
				return nil, err
			}
			oper = o.Root().Node
		}
		fname, dirname := self.fname(deviceId, moduleName)
		if err := os.MkdirAll(dirname, 0755); err != nil {
			return nil, err
		}
		b, err := self.newBrowser(fname, m, oper)
		if err != nil {
			return nil, err
		}
		d.AddBrowser(b)
	}
	return d, nil
}

func (self *FileStore) operational(deviceId string) (device.Device, error) {
	if reg, found := self.locations.LookupRegistration(deviceId); found {
		for _, p := range self.southbound {
			return p(reg.Address)
		}
	}
	return nil, nil
}

func (self *FileStore) OnUpdate(l device.ChangeListener) c2.Subscription {
	return c2.NewSubscription(self.listeners, self.listeners.PushBack(l))
}

func (self *FileStore) OnModuleUpdate(module string, l device.ChangeListener) c2.Subscription {
	return self.OnUpdate(func(d device.Device, id string, c device.Change) {
		if hnd := d.Modules()[module]; hnd != nil {
			l(d, id, c)
		}
	})
}

func (self *FileStore) updateListeners(d device.Device, id string, c device.Change) {
	p := self.listeners.Front()
	for p != nil {
		p.Value.(device.ChangeListener)(d, id, c)
		p = p.Next()
	}
}

func (self *FileStore) Add(id string, d device.Device) {
	// download yang and web?
	self.mkdir(self.deviceDir(id))
	self.updateListeners(d, id, device.Added)
}

func (self *FileStore) exists(deviceId string) bool {
	_, err := os.Stat(self.deviceDir(deviceId))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func (self *FileStore) Len() int {
	return len(self.deviceIds())
}

func mergeStrings(a []string, b []string) []string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	uniq := make(map[string]struct{})
	for _, s := range a {
		uniq[s] = struct{}{}
	}
	for _, s := range b {
		uniq[s] = struct{}{}
	}
	merged := make([]string, len(uniq))
	i := 0
	for s := range uniq {
		merged[i] = s
		i++
	}
	sort.Strings(merged)
	return merged
}

func (self *FileStore) Modules(deviceId string) []string {
	o, _ := self.operational(deviceId)
	return self.modules(deviceId, o)
}

func (self *FileStore) modules(deviceId string, o device.Device) []string {
	c := self.configModules(deviceId)
	if o == nil {
		return c
	}
	return mergeStrings(c, moduleNames(o.Modules()))
}

func moduleNames(m map[string]*meta.Module) []string {
	names := make([]string, len(m))
	i := 0
	for name := range m {
		names[i] = name
		i++
	}
	return names
}

func (self *FileStore) configModules(deviceId string) []string {
	files, err := ioutil.ReadDir(self.deviceDir(deviceId))
	if err != nil {
		return []string{}
	}
	modules := make([]string, len(files))
	j := 0
	for _, f := range files {
		fname := f.Name()
		if strings.HasSuffix(fname, ".json") {
			modules[j] = fname[:len(fname)-5]
			j++
		}
	}
	return modules[:j]
}

func (self *FileStore) newBrowser(fname string, m *meta.Module, oper node.Node) (*node.Browser, error) {
	data := make(map[string]interface{})
	dataNode := nodes.ReflectChild(data)

	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			if oper == nil {
				return nil, nil
			}
		} else {
			return nil, err
		}
	} else {
		rdr, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		defer rdr.Close()
		readOnly := nodes.ReadJSONIO(rdr)
		node.NewBrowser(m, dataNode).Root().InsertFrom(readOnly)
	}

	var n node.Node
	n = &nodes.Extend{
		Base: dataNode,
		OnEndEdit: func(p node.Node, r node.NodeRequest) error {
			if err := p.EndEdit(r); err != nil {
				return err
			}
			if r.EditRoot {
				wtr, err := os.Create(fname)
				if err != nil {
					return err
				}
				defer wtr.Close()

				// We only want to look at config and only config that isn't set to the default values
				params := "content=config&with-defaults=trim"

				// this walks data for device's data for this module (a device might have multiple
				// modules) and sends it to json
				jwtr := &nodes.JSONWtr{Out: wtr, Pretty: true}
				if err := r.Selection.Constrain(params).InsertInto(jwtr.Node()).LastErr; err != nil {
					return err
				}
			}
			return nil
		},
	}
	if oper != nil {
		n = nodes.ConfigProxy{}.Node(n, oper)
	}
	return node.NewBrowser(m, n), nil
}
