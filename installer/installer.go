package installer

import (
	"flag"
	"fmt"
	"github.com/bluepongo/mysql_autoInstall/conf"
	"github.com/bluepongo/mysql_autoInstall/install"
	"github.com/bluepongo/mysql_autoInstall/parameters"
	"github.com/romberli/log"
	"strconv"
)

const (
	LogFilePath = "/tmp/test.log"
	// Need to change
)
func Execute() {
	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		panic(err)
	}
	log.Info("Initial log file success.")


}
