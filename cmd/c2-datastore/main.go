package main

import (
	"flag"
	"log"
	"os"

	"context"

	"strings"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/conf"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/restconf"
)

var portParam = flag.String("port", ":8080", "restconf port")

func main() {
	flag.Parse()
	yangPath := yang.YangPath()
	d := conf.NewLocalDevice(yangPath)
	d.Add("restconf", restconf.Node(restconf.NewManagement(d, *portParam)))
	d.Add("ietf-yang-lib", conf.LocalDeviceYangLibNode(d))
	d.Add("proxy", conf.ProxyNode(conf.NewProxy(yangPath, restconf.NewDeviceByHostAndPort)))
	if err := startupConfigs(d, os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	select {}
}

func startupConfigs(d *conf.LocalDevice, configs []string) error {
	for _, config := range configs {
		moduleFilePair := strings.Split(config, "=")
		b, err := d.Browser(moduleFilePair[0])
		if err != nil {
			return err
		} else if b == nil {
			return c2.NewErr("Module not found: " + moduleFilePair[0])
		}
		if err = configure(b, moduleFilePair[1]); err != nil {
			return err
		}
	}
	return nil
}

func configure(b *node.Browser, filename string) error {
	configFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer configFile.Close()
	config := node.NewJsonReader(configFile).Node()
	if err = b.Root().UpsertFrom(context.Background(), config).LastErr; err != nil {
		return err
	}
	return err
}
