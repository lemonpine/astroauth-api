package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func AppUserRegisterValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			HWID     string `json:"hwid"`
			AppID    string `json:"app_id" `
			License  string `json:"license" `
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		err := validation.ValidateStruct(&r,
			validation.Field(&r.Email, validation.Required, is.Email),
			validation.Field(&r.Username, validation.Required),
			validation.Field(&r.Password, validation.Required),
			validation.Field(&r.HWID, validation.Required),
			validation.Field(&r.AppID, validation.Required),
			validation.Field(&r.License, validation.Required),
		)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			c.Abort()
			return
		}
	}
}
