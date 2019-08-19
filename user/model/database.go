package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// gorm require
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateConn 创建连接
func CreateConn() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(host, name, user, password)
	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			name,
		),
	)
}
