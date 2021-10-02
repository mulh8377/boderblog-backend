package users

import "github.com/gin-gonic/gin"

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username"`
		Email string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
		Bio string `form:"bio" json:"bio"`
		Image string `form:"image" json:"image"`

	} `json:"user"`
	userModel UserModel `json:"-"`
}

type LoginValidator struct {
	User struct {
		Email string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {

	return nil

}

func (self *LoginValidator) Bind(c *gin.Context) error {

	return nil

}

func emptyUserModelInvalidator() UserModelValidator {
	temp := UserModelValidator{}
	return temp
}

func emptyLoginInvalidator() LoginValidator {
	temp := LoginValidator{}
	return temp
}

func initUserModelInvalidator(userModel UserModel) UserModelValidator {
	user := emptyUserModelInvalidator()

	// fill out the rest

	return user
}