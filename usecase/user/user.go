package user

import (
	"context"
	"time"
	"todo-list-app/domain"
	"todo-list-app/utils"
)

type userUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

// CreateOrUpdate implements domain.UserUsecase.
func (u *userUsecase) CreateOrUpdate(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.CreateOrUpdate(ctx, user)
}

// Delete implements domain.UserUsecase.
func (u *userUsecase) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.Delete(ctx, id)
}

// FindByID implements domain.UserUsecase.
func (u *userUsecase) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.FindByID(ctx, id)
}

// FindByUsername implements domain.UserUsecase.
func (u *userUsecase) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.FindByUsername(ctx, username)
}

func NewUserUsecase(userRepository domain.UserRepository, timeout int) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		timeout:        time.Duration(utils.TimeoutContext),
	}
}
