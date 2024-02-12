package healtcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healtcheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	}
}
