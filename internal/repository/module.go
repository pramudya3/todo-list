package repository

import (
	"todo-list-app/internal/repository/user"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewUserRepository,
)
