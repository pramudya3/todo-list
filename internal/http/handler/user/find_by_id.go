package user

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-list-app/domain"

	"github.com/gin-gonic/gin"
)

type response struct {
	User *domain.User `json:"user"`
}

func FindByID(uc domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		fmt.Println(ctx)

		idString := c.Param("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := uc.FindByID(ctx, uint64(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, response{user})
	}
}
