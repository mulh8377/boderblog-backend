package users

import "github.com/gin-gonic/gin"
import "github.com/mulh8377/boderblog-backend/common"

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Bio string `json:"bio"`
	Image *string `json:"image"`
	Token string `json:"token"`
}

func (self * UserSerializer) Response() UserResponse {
	myModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myModel.UserName,
		Email: myModel.Email,
		Bio: myModel.Bio,
		Image: myModel.Image,
		Token: common.GenToken(myModel.ID),
	}
	return user
}