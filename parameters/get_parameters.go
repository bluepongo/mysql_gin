package parameters

import (
	"flag"
	"strings"
)

type IpPort struct {
	Ip   string
	Port string
}

var Ip = flag.String("ip", "192.168.59.4:3306", "Input The IP, Use ',' separated")

func ExtractIP(Ips string) (res []IpPort) {
	ipPorts := strings.Split(Ips, ",")
	for _, ipPort := range ipPorts {
		ip := strings.Split(ipPort, ":")[0]
		port := strings.Split(ipPort, ":")[1]
		ipport := IpPort{ip, port}
		res = append(res, ipport)
	}
	return res
}
