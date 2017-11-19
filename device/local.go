package device

import (
	"encoding/json"
	"io"
	"os"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

type Local struct {
	browsers     map[string]*node.Browser
	schemaSource meta.StreamSource
	uiSource     meta.StreamSource
}

func New(schemaSource meta.StreamSource) *Local {
	return &Local{
		schemaSource: schemaSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func NewWithUi(schemaSource meta.StreamSource, uiSource meta.StreamSource) *Local {
	return &Local{
		schemaSource: schemaSource,
		uiSource:     uiSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func (self *Local) SchemaSource() meta.StreamSource {
	return self.schemaSource
}

func (self *Local) UiSource() meta.StreamSource {
	return self.uiSource
}

func (self *Local) Modules() map[string]*meta.Module {
	mods := make(map[string]*meta.Module)
	for _, b := range self.browsers {
		mods[b.Meta.Ident()] = b.Meta
	}
	return mods
}

func (self *Local) Browser(module string) (*node.Browser, error) {
	return self.browsers[module], nil
}

func (self *Local) Close() {
}

func (self *Local) Add(module string, n node.Node) error {
	m, err := yang.LoadModule(self.schemaSource, module)
	if err != nil {
		return err
	}
	self.browsers[module] = node.NewBrowser(m, n)
	return nil
}

func (self *Local) AddBrowser(b *node.Browser) {
	self.browsers[b.Meta.Ident()] = b
}

func (self *Local) ApplyStartupConfig(config io.Reader) error {
	var cfg map[string]interface{}
	if err := json.NewDecoder(config).Decode(&cfg); err != nil {
		return err
	}
	return self.ApplyStartupConfigData(cfg)
}

func (self *Local) ApplyStartupConfigData(config map[string]interface{}) error {
	for module, data := range config {
		b, err := self.Browser(module)
		if err != nil {
			return err
		}
		if b == nil {
			return c2.NewErrC("browser not found:"+module, 404)
		}
		moduleCfg := data.(map[string]interface{})
		if err := b.Root().UpsertFromSetDefaults(nodes.ReflectChild(moduleCfg)).LastErr; err != nil {
			return err
		}
	}
	return nil
}

func (self *Local) ApplyStartupConfigFile(fname string) error {
	cfgRdr, err := os.OpenFile(fname, os.O_RDWR, os.ModeExclusive)
	defer cfgRdr.Close()
	if err != nil {
		panic(err)
	}
	return self.ApplyStartupConfig(cfgRdr)
}
