package users

import (
	"errors"
	_ "errors"
	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID			uint `gorm:"primary_key"`
	UserName 	string `gorm:"column:username"`
	Email		string `gorm:"column:email;unique_index"`
	Bio			string `gorm:"column:bio;size:1024"`
	Image		*string `gorm:"column:image"`
	Password	string `gorm:"column:password;not null"`
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password cannot be empty")
	}
	bytes := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytes := []byte(password)
	check := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(bytes, check)
}