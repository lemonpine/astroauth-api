package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
)

func AppCreateValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Name string `json:"name" gorm:"uniqueIndex"`
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		err := validation.ValidateStruct(&r,
			validation.Field(&r.Name, validation.Required),
		)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			c.Abort()
			return
		}
	}
}
