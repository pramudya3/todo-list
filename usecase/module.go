package usecase

import (
	"todo-list-app/usecase/user"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewUserUsecase,
)
