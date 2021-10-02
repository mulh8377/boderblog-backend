package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mulh8377/boderblog-backend/users"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	db, err := gorm.Open("sqlite3", "./boder.db")
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	db.DB().SetMaxIdleConns(10)

	Migrate(db)

	defer db.Close()

	r := gin.Default()

	r.Run()
}