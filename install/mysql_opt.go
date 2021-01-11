package install

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Initialize a database function
func InitDB(Username, Password, Port, Database string) (db *sql.DB, err error) {
	// DSN:Data Source Name
	dsn := Username + ":" + Password + "@tcp(localhost:" + Port + ")/" + Database + "?charset=utf8"
	// Check the password
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// attempt connect to the databse
	err = db.Ping()
	if err != nil {
		fmt.Printf("Initialize failed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("Database initialization succeeded.")
	return db, nil
}

// Exec the sql string
func ExecMysql(db *sql.DB, sqlStr string) (dbNew *sql.DB, err error) {
	_, err = db.Exec(sqlStr)
	fmt.Println("mysql >", sqlStr)
	if err != nil {
		return db, err
	}
	return db, err
}

type test struct {
	id   int
	name string
}

// Query many rows
func QueryMany(db *sql.DB, sqlStr string) (dbNew *sql.DB, err error) {
	rows, err := db.Query(sqlStr)
	fmt.Println("mysql >", sqlStr)
	if err != nil {
		return db, err
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var t test
		err := rows.Scan(&t.id, &t.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return db, err
		}
		fmt.Printf("id:%d  name:%s\n", t.id, t.name)
	}
	return db, err
}
