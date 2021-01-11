package install

import (
	"github.com/romberli/log"
)

const (
	HostIP      = "192.168.59.2"
	PortNum     = 22
	SSHUserName = "root"
	SSHPassWord = "root"
	SSHData     = "/mysqldata/"
	SSHDatadir  = "/mysqldata/mysql3306/data"
	SSHLog      = "/mysqldata/mysql3306/log"
	SSHTmp      = "/mysqldata/mysql3306/tmp"
	SSHBinlog   = "/mysqllog/mysql3306/binlog/"
	SSHRealylog = "/mysqllog/mysql3306/relaylog/"
	SSHLogData  = "/mysqllog/mysql3306/data/"
	SSHLogDir   = "/mysqllog/"
)

// Install mysql remotely via SSH connection
func InstallMysqlSSH(ip, port string) {

	// Establish the ssh connection
	log.Info("==========Install mysql remotely started==========")
	sshConn, err := EstablishSSHConnect(ip, PortNum, SSHUserName, SSHPassWord)
	if err != nil {
		log.Warnf("Can't establish the ssh connection: %v", err)
		return
	}
	log.Info("==========Install mysql remotely completed==========")

	// Execute remote shell command
	// Create the user and group
	result, stdOut, err := AddUserGroupSSH(sshConn, UserName, GroupName)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	log.Info("==========Add user and group completed==========")
	// UnTarGz the file
	//stdErr, err := UnTarLocal(RelatedPath, MySQLTarName, MySQLFileName, MySQL)
	//if err != nil {
	//	log.Warnf("%v: %s", err, stdErr)
	//	return
	//}

	// Move the folder to remote
	err = CopyMysqlToRemote(sshConn, RelatedPath, RootPath)
	if err != nil {
		log.Warnf("Can't copy the file to remote: %v", err)
		return
	}
	err = CopyMyCnfToRemote(sshConn, RelatedPath+MyCnfFileName, EtcPath)
	if err != nil {
		log.Warnf("Can't copy the file to remote: %v", err)
		return
	}
	log.Info("Create the my.cnf and move to remote complete.")

	result, stdOut, err = MkdirSSH(sshConn, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHLog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHTmp)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHBinlog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHRealylog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHLogData)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = ChownSSH(sshConn, UserName, GroupName, SSHData)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = ChmodSSH(sshConn, SSHData)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}

	result, stdOut, err = ChownSSH(sshConn, UserName, GroupName, MySQLMultiRoot)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = ChmodSSH(sshConn, MySQLMultiRoot)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = ChownSSH(sshConn, UserName, GroupName, SSHLogDir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = ChmodSSH(sshConn, SSHLogDir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}

	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysql", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqld", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqld_safe", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqld_multi", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqldump", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqlbinlog", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysql_config_editor", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/my_print_defaults", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	result, stdOut, err = CpSSH(sshConn, MySQLMultiBin+"/mysqladmin", BinPath)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}

	log.Info("Prepare to initial the mysqld...")
	result, stdOut, err = InitMysqlSSH(sshConn, MySQLDPath, UserName, BaseDirPath, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	// /usr/local/mysql/bin/mysqld --initialize-insecure --user=mysql --basedir=/usr/local/mysql --datadir=/mysqldata/mysql3306/data
	log.Info("==========Initial mysqld complete==========")

	result, stdOut, err = MysqldStartSSH(sshConn, port)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
		return
	}
	log.Info("==========Finish==========")

}
