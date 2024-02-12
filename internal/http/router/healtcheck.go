package router

import (
	"todo-list-app/internal/http/handler/healtcheck"

	"github.com/gin-gonic/gin"
)

func NewHealthcheckRoutes(g *gin.RouterGroup) {
	g.GET("/", healtcheck.Healtcheck())
}
