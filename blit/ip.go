package blit

import "net"

func GetIpForIface(ifaceName string) string {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		panic(err.Error())
	}
	addrs, addrErr := iface.Addrs()
	if addrErr != nil {
		panic(addrErr.Error())
	}
	if len(addrs) == 0 {
		panic("No addresses available")
	}
	switch v := addrs[0].(type) {
	case *net.IPNet:
		return v.IP.String()
	case *net.IPAddr:
		return v.IP.String()
	}
	panic("Unable to get IP address")
}