package iutils

import (
	"net"
	"os"
)

func GetLocalIP4() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	ips := map[string]string{}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ips[netInterfaces[i].Name] = ipnet.IP.String()
					}
				}
			}
		}
	}
	if ips["eth0"] != "" {
		return ips["eth0"]
	}
	for _, v := range ips {
		return v
	}
	return ""
}

func GetLocalHost() string {
	host, _ := os.Hostname()
	return host
}
