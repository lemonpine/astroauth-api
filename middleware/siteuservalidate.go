package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func SiteRegisterValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Email    string `json:"email" gorm:"uniqueIndex"`
			Password string `json:"password"`
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		err := validation.ValidateStruct(&r,
			validation.Field(&r.Email, validation.Required, is.Email),
			validation.Field(&r.Password, validation.Required),
		)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			c.Abort()
			return
		}
	}
}
