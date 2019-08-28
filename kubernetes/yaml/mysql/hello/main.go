package main

import (
	"database/sql" // sql 官方接口
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
)

// CreateDB 创建数据库连接
func CreateDB(DbDriver string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DbDriver)
	if err != nil {
		return nil, err
	}
	return db, err
}

func main() {
	host := os.Getenv("MYSQL_SERVICE_HOST")
	port := os.Getenv("MYSQL_SERVICE_PORT")
	dbDriver := fmt.Sprintf("root:123456@tcp(%s:%s)/mysql", host, port)
	db, err := CreateDB(dbDriver)
	if err != nil {

	}

	rows, err := db.Query("show tables")
	if err != nil {

	}

	for rows.Next() {
		var table string
		rows.Scan(&table)
		fmt.Println(table)
	}
	fmt.Println("end - success")
}
