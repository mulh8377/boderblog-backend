package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UserRegistration)
	router.POST("/login", UserLogin)
}

func UserRegistration(c *gin.Context) {

}

func UserLogin(c *gin.Context) {

}