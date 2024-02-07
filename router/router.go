package router

import (
	"todo-list-app/domain"
	"todo-list-app/handler/healtcheck"
	"todo-list-app/handler/user"

	"github.com/gin-gonic/gin"
)

func NewRoutes(
	ucUser domain.UserUsecase,
) *gin.Engine {
	r := gin.Default()

	hc := r.Group("/healthcheck")
	hc.GET("/", healtcheck.Healtcheck())

	u := r.Group("/users")
	u.GET("/:id", user.FindByID(ucUser))
	return r
}
