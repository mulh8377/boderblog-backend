module github.com/mulh8377/boderblog-backend/m

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.7.4 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/mulh8377/boderblog-backend/common v0.0.0-20211003000243-7de95f752a61 // indirect
	github.com/mulh8377/boderblog-backend/images v0.0.0-20211003000243-7de95f752a61 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace github.com/mulh8377/boderblog-backend/users => ./users

replace github.com/mulh8377/boderblog-backend/images => ./images

replace github.com/mulh8377/boderblog-backend/common => ./common
