package dataRelated

import (
	"fmt"
	"github.com/bluepongo/mysql_autoInstall/install"
)

func CreateTest() (err error) {

	// Start the Mysql operations
	// Initialize the database
	db, err := install.InitDB("root", "", "3306", "mysql")
	if err != nil {
		return err
	}
	// Create a admin user
	db, err = install.ExecMysql(db, "grant shutdown on *.* to 'admin'@'localhost' identified by 'admin'")
	if err != nil {
		return err
	}
	// Create a replica account
	db, err = install.ExecMysql(db, "grant replication slave, replication client on *.* to 'replication'@'%' identified by 'admin'")
	if err != nil {
		return err
	}
	// Create a database and use it
	db, err = install.ExecMysql(db, "drop database spdb")
	if err != nil {
		return err
	}
	db, err = install.ExecMysql(db, "create database spdb")
	if err != nil {
		return err
	}
	db, err = install.ExecMysql(db, "use spdb")
	if err != nil {
		return err
	}
	// Create the test data
	db, err = install.ExecMysql(db, fmt.Sprintf(
		`create table t01(
			id int(10) primary key auto_increment comment 'pid',
			name varchar(100) comment 'name'
		) engine= innodb default charset=utf8mb4`))
	if err != nil {
		return err
	}
	db, err = install.ExecMysql(db, "insert into t01(name) values('a'), ('b'), ('c')")
	if err != nil {
		return err
	}
	db, err = install.QueryMany(db, "select * from t01")
	if err != nil {
		return err
	}
	return err
}
