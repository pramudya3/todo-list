package user

import (
	"context"
	"time"
	"todo-list-app/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

// CreateOrUpdate implements domain.UserRepository.
func (u *userRepository) CreateOrUpdate(ctx context.Context, user *domain.User) error {
	if _, err := u.db.Exec(ctx, `INSERT INTO users (username, name, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`, user.Username, user.Name, user.Password, time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(ctx context.Context, id uint64) error {
	if _, err := u.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}

// FindByID implements domain.UserRepository.
func (u *userRepository) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	user := &domain.User{}
	if err := pgxscan.Get(ctx, u.db, user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return user, nil
}

// FindByUsername implements domain.UserRepository.
func (u *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	if err := pgxscan.Get(ctx, u.db, user, `SELECT * FROM users WHERE username = $1`, username); err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
