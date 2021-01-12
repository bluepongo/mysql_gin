package main

import (
	// "github.com/bluepongo/mysql_autoInstall/installer"
	"fmt"
	"github.com/bluepongo/mysql_autoInstall/conf"
	"github.com/bluepongo/mysql_autoInstall/install"
	"github.com/gin-gonic/gin"
	"github.com/romberli/log"
	"strings"
)

const (
	LogFilePath = "/tmp/test.log"
	// Need to change
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		IpPorts := c.QueryArray("ip")
		// Initialize a logger
		fileName := LogFilePath
		_, _, err := log.InitLoggerWithDefaultConfig(fileName)
		if err != nil {
			panic(err)
		}
		log.Info("Initial log file success.")
		for _, IpPort := range IpPorts {
			ip := strings.Split(IpPort, ":")[0]
			port := strings.Split(IpPort, ":")[1]
			fmt.Printf("Current:Ip:%s, Port:%s\n", ip, port)
			log.Infof("Current:Ip:%s, Port:%d", ip, port)
			fmt.Println("=========Prepare to generate mycnf=========")
			err = conf.GenerateMyCnf(ip, port)
			if err != nil {
				log.Warnf("%v", err)
				return
			}
			fmt.Println("=========Prepare to create ssh connection=========")
			install.InstallMysqlSSH(ip, port)
			fmt.Println("=========Finish=========")
		}
	})
	router.Run(":8080")
}