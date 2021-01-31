package nets

import (
	"errors"
	"fmt"
	"net"
)

func GetIntranetIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("net.InterfaceAddrs: %v", err)
	}

	for _, addr := range addrs {
		if IPNet, ok := addr.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if IPv4 := IPNet.IP.To4(); IPv4 != nil {
				return IPv4.String(), nil
			}
		}
	}

	return "", errors.New("unable to get the intranet IP")
}
