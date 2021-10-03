package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mulh8377/boderblog-backend/common"
)

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
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.UserName = self.User.Username
	self.userModel.Email = self.User.Email
	self.userModel.Bio = self.User.Bio

	if self.User.Password != common.NBRandomPassword {
		self.userModel.setPassword(self.User.Password)
	}
	if self.User.Image != "" {
		self.userModel.Image = &self.User.Image
	}

	return nil

}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.Email = self.User.Email
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

	user.User.Username = userModel.UserName
	user.User.Email = userModel.Email
	user.User.Bio = userModel.Bio
	user.User.Password = common.NBRandomPassword
	// fill out the rest

	if userModel.Image != nil {
		user.User.Image = *userModel.Image
	}

	return user
}