package router

import (
	"todo-list-app/http/handler/healtcheck"

	"github.com/gin-gonic/gin"
)

func NewHealthcheckRoutes(g *gin.RouterGroup) {
	g.GET("/", healtcheck.Healtcheck())
}
