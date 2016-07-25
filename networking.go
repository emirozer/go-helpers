package helpers

import (
	"os"
	"fmt"
	"net"
)

// GetLocalIpv4 returns a string of the host machine local ipv4
func GetLocalIpv4() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return fmt.Sprintf("%s", ipv4)
		}
	}
	return "localhost"
}
