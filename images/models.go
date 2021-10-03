package images

import (
	"github.com/jinzhu/gorm"
	"github.com/mulh8377/boderblog-backend/users"
)

type ImageModel struct {
	gorm.Model
	Slug		string `gorm:unique_index`
	Title		string
	Description	string 	`gorm:"size2048"`
	Body		string 	`gorm:"size:2048"`
	User		ImageUserModel
	UserID		uint
	Tags		[]TagModel	`gorm:"many2many:article_tags;"`
	Comments	[]CommentModel	`gorm:"ForeignKey:ArticleID"`
}

type ImageUserModel struct {
	gorm.Model
	UserModel users.UserModel
	UserModelID	uint
	ImageModels		[]ImageModel `gorm:"ForeignKey:AuthorID"`
	FavoriteModels	[]FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}

type FavoriteModel struct {
	gorm.Model
	Favorite	ImageModel
	FavoriteID	uint
	FavoriteBy	ImageUserModel
	FavoriteByID uint
}

type TagModel struct {
	gorm.Model
	Tag string `gorm:"unique_index"`
	ImageModels []ImageModel `gorm:"many2many:article_tags;"`
}

type CommentModel struct {
	gorm.Model
	Image ImageModel
	ImageID uint
	User	ImageUserModel
	UserID  uint
	Body	string `gorm:"size:2048"`
}