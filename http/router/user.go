package router

import (
	"todo-list-app/domain"
	"todo-list-app/http/handler/user"

	"github.com/gin-gonic/gin"
)

func NewUserRoutes(g *gin.RouterGroup, ucUser domain.UserUsecase) {
	g.GET("/:id", user.FindByID(ucUser))
}
