package net

import (
	"fmt"
	"net"
	"testing"
)

func TestGetLocalAddress(t *testing.T) {

	addresses, err := net.InterfaceAddrs()

	if err != nil {
		t.Fail()
	}

	for _, add := range addresses {

		if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
