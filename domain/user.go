package domain

import (
	"context"
	"database/sql"
)

type (
	UserUsecase interface {
		CreateOrUpdate(ctx context.Context, user *User) error
		FindByUsername(ctx context.Context, username string) (*User, error)
		FindByID(ctx context.Context, id uint64) (*User, error)
		Delete(ctx context.Context, id uint64) error
	}

	UserRepository interface {
		CreateOrUpdate(ctx context.Context, user *User) error
		FindByUsername(ctx context.Context, username string) (*User, error)
		FindByID(ctx context.Context, id uint64) (*User, error)
		Delete(ctx context.Context, id uint64) error
	}

	User struct {
		ID        int          `json:"id" db:"id"`
		Username  string       `json:"username" db:"username" binding:"required,alphanum"`
		Name      string       `json:"name" db:"name" binding:"required,alpha"`
		Password  string       `json:"password" db:"password" binding:"required"`
		Email     string       `json:"email" db:"email" binding:"required,email"`
		CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
		UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	}
)
