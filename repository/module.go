package repository

import (
	"todo-list-app/repository/user"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewUserRepository,
)
