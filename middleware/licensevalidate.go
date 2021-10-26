package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
)

func AddLicenseValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			AppID  string `json:"app_id" `
			Length uint   `json:"length" `
			Level  uint   `json:"level" `
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		err := validation.ValidateStruct(&r,
			validation.Field(&r.AppID, validation.Required),
			validation.Field(&r.Length, validation.Required),
			validation.Field(&r.Level, validation.Required),
		)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			c.Abort()
			return
		}
	}
}
