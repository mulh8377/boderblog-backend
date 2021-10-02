package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UserRegistration)
	router.POST("/login", UserLogin)
}

func UserRegistration(c *gin.Context) {

}

func UserLogin(c *gin.Context) {

}