package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
)

func Open(dbHost, dbPort, dbName, username, password string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		username, password, dbHost, dbPort, dbName, url.QueryEscape("Asia/Shanghai"))
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return
}
