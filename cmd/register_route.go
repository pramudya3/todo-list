package cmd

import (
	"todo-list-app/http/router"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginRouter *gin.Engine, params Params) *gin.Engine {
	// initiate routers, if there is new
	// router's group, just add in this line
	router.NewHealthcheckRoutes(ginRouter.Group("/healthcheck"))
	router.NewUserRoutes(ginRouter.Group("/users"), params.UserUsecase)

	return ginRouter
}
