package helpers

import (
	"net"
	"strings"
)

func LocalInterfaceIPs() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := make([]net.IP, 0)
	for _, addr := range addrs {
		addr1 := addr.String()
		switch addr.Network() {
		case "ip+net":
			addr1 = strings.Split(addr1, "/")[0]
		}
		if ip := net.ParseIP(addr1); ip != nil {
			ips = append(ips, ip)
		}
	}

	return ips, nil
}
