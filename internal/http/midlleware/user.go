package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExtractUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
