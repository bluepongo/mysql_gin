package install

import (
	"github.com/romberli/log"
)

const (
	RootPath     = "/usr/local/"
	DataDirPath  = "/usr/local/mysql/data/"
	BaseDirPath  = "/usr/local/mysql/"
	MySQLDPath   = "/usr/local/mysql/bin/mysqld"
	EtcPath      = "/etc/"
	LimitsPath   = "/etc/security/"
	BashFilePath = "/home/mysql/"

	ShareFilePath = "/mnt/hgfs/share/"

	MySQLTarName  = "mysql-5.7.31-linux-glibc2.12-x86_64.tar.gz"
	MySQLFileName = "mysql-5.7.31-linux-glibc2.12-x86_64"
	MyCnfFileName = "my.cnf"
	LimitsFile    = "limits.conf"
	BashFile      = ".bash_profile"
	AutoCnfFile   = "auto.cnf"

	GroupName = "mysql"
	UserName  = "mysql"
	MySQL     = "mysql"

	// MySQLMulti
	MySQLMulti_3306       = "/usr/local/mysql/3306/"
	MySQLMulti_3307       = "/usr/local/mysql/3307"
	MySQLMultiBinlog_3306 = "/usr/local/mysql/3306/binlog"
	MySQLMultiData_3306   = "/usr/local/mysql/3306/data"
	MySQLMultiData_3307   = "/usr/local/mysql/3307/data/"
	MySQLMultiRoot        = "/usr/local/mysql"
	MySQLMultiBin         = "/usr/local/mysql/bin"
	MultiDataDir          = "/usr/local/mysql/3306/data"

	RelatedPath = "./related/"

	BinPath = "/usr/bin"

	PortNum_3306 = "3306"
	PortNum_3307 = "3307"

	// Need to change

)

// Install multiple instances of mysql
func InstallMySQLMul() {
	// 1、Create mysql user and user group
	_, stdErr, err := AddGroup(GroupName)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = AddUser(GroupName, UserName)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========Create the mysql user and group success=========")

	// 2、Create the catalogue
	_, stdErr, err = Mkdir(MySQLMultiBinlog_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = Mkdir(MySQLMultiData_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = Chown(GroupName, UserName, MySQLMultiRoot)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 3、Alter the file limits.cnf
	_, stdErr, err = Cp(RelatedPath+LimitsFile, LimitsPath+LimitsFile)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========Alter the flie limits.cnf success=========")

	// 4、Unzip the file
	stdErr, err = UnTarGz(RelatedPath+MySQLTarName, RelatedPath)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========UnTarGz the flie success=========")

	// 5、Move the mysql file
	_, stdErr, err = Mv(RelatedPath+MySQLFileName, RelatedPath+MySQL)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = Cp(RelatedPath+MySQL+"/*", MySQLMultiRoot)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========Move the mysql file success=========")

	// 6、Alter the my_basic.cnf file
	_, stdErr, err = Cp(RelatedPath+MyCnfFileName, EtcPath+MyCnfFileName)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========Alter the my_basic.cnf file success=========")

	// 7、Create the folder data/tmp/log/pid/sock to exam
	stdErr, err = CreateFolder(PortNum_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 8、Batch execute cp command
	stdErr, err = BatchCpBin()
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 9、Alter the .bash_profile
	_, stdErr, err = Cp(RelatedPath+BashFile, BashFilePath+BashFile)
	stdErr, err = BatchCpBin()
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 10、Initializes the mysql instance
	_, stdErr, err = MultiInitMysql(MySQLDPath, UserName, BaseDirPath, MultiDataDir)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	log.Info("=========Initializes the mysql instance success=========")

	// 11、Start the example
	_, stdErr, err = MultiStartMysql(PortNum_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
}

// Build master-slave replication
func BuildMS() {
	// 1、Stop the 3306 instance
	_, stdErr, err := MultiStopMysql(PortNum_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 2、Copy data file from 3306 to 3307
	_, stdErr, err = Cp(MySQLMulti_3306, MySQLMulti_3307)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = Chown(GroupName, UserName, MySQLMulti_3307)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = Rm(MySQLMultiData_3307 + AutoCnfFile)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}

	// 3、Start the 3306、3307
	_, stdErr, err = MultiStartMysql(PortNum_3306)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	_, stdErr, err = MultiStartMysql(PortNum_3307)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
		return
	}
	// Initialize the database
	db, err := InitDB("root", "", "3307", "mysql")
	if err != nil {
		log.Warnf("Initialize the database failed: %v", err)
		return
	}
	// change the master
	db, err = ExecMysql(db, "change master to master_host='localhost', master_port=3306, master_user='replication', master_password='admin', master_auto_position=1")
	if err != nil {
		log.Warnf("Change the master failed: %v", err)
		return
	}
	db, err = ExecMysql(db, "start slave")
	if err != nil {
		log.Warnf("Start the slave failed: %v", err)
		return
	}
}
