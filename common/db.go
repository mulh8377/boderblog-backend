package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Create() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../boder.db")
	if err != nil {
		fmt.Println("db err: (Create) ", err)
	}

	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func Get() *gorm.DB {
	return DB
}