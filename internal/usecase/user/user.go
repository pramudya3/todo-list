package user

import (
	"context"
	"errors"
	"time"
	"todo-list-app/domain"
	"todo-list-app/internal/utils"

	"github.com/jackc/pgx/v5"
)

type userUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func (u *userUsecase) CreateOrUpdate(ctx context.Context, user *domain.User) error {
	// validation start here:
	_, err := u.userRepository.FindByUsername(ctx, user.Username)
	if !errors.Is(err, pgx.ErrNoRows) {
		return errors.New("username already taken")
	}

	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.CreateOrUpdate(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.Delete(ctx, id)
}

func (u *userUsecase) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.FindByID(ctx, id)
}

func (u *userUsecase) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepository.FindByUsername(ctx, username)
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		timeout:        time.Second * time.Duration(utils.TimeoutContext),
	}
}
