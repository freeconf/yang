package automate

import (
	"fmt"

	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
)

func Garage(sys System, speed int) error {
	for i := 0; i < 10; i++ {
		if hnd, err := sys.New("car"); err != nil {
			return err
		} else {
			if b, err := hnd.Device.Browser("car"); err != nil {
				return err
			} else {
				if err := configureCar(b, speed); err != nil {
					return err
				}
			}
		}
	}
	if _, err := sys.New("garage"); err != nil {
		return err
	}
	return nil
}

func configureCar(b *node.Browser, speed int) error {
	cfg := fmt.Sprintf(`{"speed":%d}`, speed)
	return b.Root().UpsertFrom(nodes.ReadJSON(cfg)).LastErr
}
