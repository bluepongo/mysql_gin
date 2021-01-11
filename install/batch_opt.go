package install

// Batch create the folder to example
func CreateFolder(examName string) (stdErr string, err error) {
	_, stdErr, err = Mkdir(MySQLMultiRoot + "/" + examName + "/data")
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Mkdir(MySQLMultiRoot + "/" + examName + "/tmp")
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Mkdir(MySQLMultiRoot + "/" + examName + "/sock")
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Mkdir(MySQLMultiRoot + "/" + examName + "/log")
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Mkdir(MySQLMultiRoot + "/" + examName + "/pid")
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Chmod(MySQLMultiRoot)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Chown(GroupName, UserName, MySQLMultiRoot)
	if err != nil {
		return stdErr, err
	}
	return stdErr, err
}

// Batch execute the cp command
func BatchCpBin() (stdErr string, err error) {
	_, stdErr, err = Cp(MySQLMultiBin+"/mysql", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqld", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqld_safe", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqld_multi", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqldump", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqlbinlog", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysql_config_editor", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/my_print_defaults", BinPath)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = Cp(MySQLMultiBin+"/mysqladmin", BinPath)
	if err != nil {
		return stdErr, err
	}
	return stdErr, err
}
