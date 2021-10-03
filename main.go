package main

import (
	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mulh8377/boderblog-backend/common"
	_ "github.com/mulh8377/boderblog-backend/common"
	"github.com/mulh8377/boderblog-backend/users"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	db := common.Create()

	Migrate(db)

	defer db.Close()

	r := gin.Default()

	r.Run()
}