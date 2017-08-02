package db

import (
	"fmt"
	"os"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

// Store all data in simple files.  Normally you would save this to a highly
// available service like a database.
type FileStore struct {
	VarDir string
}

func (self FileStore) fname(deviceId string, module string) string {
	return fmt.Sprintf("%s/%s:%s.json", self.VarDir, deviceId, module)
}

// DbRead implements device.DbIO interface to load data
func (self FileStore) DbRead(deviceId string, module string, b *node.Browser) error {
	fname := self.fname(deviceId, module)
	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	rdr, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer rdr.Close()
	// this walks data for device's data for this module (a device might have multiple
	// modules) and sends it to json
	if err := b.Root().InsertFrom(nodes.ReadJSONIO(rdr)).LastErr; err != nil {
		return err
	}
	return nil
}

// DbWrite implements device.DbIO interface to save data
func (self FileStore) DbWrite(deviceId string, module string, b *node.Browser) error {
	fname := self.fname(deviceId, module)
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
	if err := b.Root().Constrain(params).InsertInto(jwtr.Node()).LastErr; err != nil {
		return err
	}
	return nil
}
