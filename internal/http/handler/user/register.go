package user

import (
	"net/http"
	"todo-list-app/domain"
	"todo-list-app/internal/utils"

	"github.com/gin-gonic/gin"
)

func Register(ucUser domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := &domain.User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseFailed{Message: utils.ValidationError(err)})
			return
		}

		if err := ucUser.CreateOrUpdate(ctx, user); err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseFailed{Message: err.Error()})
		}

		c.Status(http.StatusCreated)
	}
}
