package parameters

import (
	"flag"
	"fmt"
	"strconv"
	"testing"
)

func TestGetParameters(t *testing.T) {
	Ips := "192.168.59.02:3306"
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	// 解析传入参数
	IpPorts := ExtractIP(Ips)
	for _, IpPort := range IpPorts {
		CurrentIP := IpPort.Ip
		CurrentPort, _ := strconv.Atoi(IpPort.Port)
		fmt.Printf("Current:Ip:%s, Port:%d", CurrentIP, CurrentPort)
	}
}
