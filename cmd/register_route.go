package cmd

import (
	"todo-list-app/internal/http/router"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginRouter *gin.Engine, params Params) *gin.Engine {

	// initiate routers, if there is new
	// router's group, just add in this line
	group := ginRouter.Group("api/v1")

	router.NewHealthcheckRoutes(group.Group("/healthz"))
	router.NewUserRoutes(group.Group("/users"), params.UserUsecase)

	return ginRouter
}
