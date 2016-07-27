package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db gorm.DB
var dberr error

func init() {
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, dberr = gorm.Open("mysql", "heiyo:heiyo@tcp(127.0.0.1:3306)/heiyo?charset=utf8&parseTime=True&loc=Local")
	if dberr != nil {
		panic(dberr)
	}
}
