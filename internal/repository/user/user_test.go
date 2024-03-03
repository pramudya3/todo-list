package user

import (
	"context"
	"testing"
	"todo-list-app/database"
	"todo-list-app/domain"
	"todo-list-app/domain/password"
	"todo-list-app/internal/utils"

	"github.com/stretchr/testify/require"
)

func TestQueryFindByIDSuccess(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	assert.Nil(err)
	assert.NotNil(db)

	u := NewUserRepository(db)
	user, err := u.FindByID(context.Background(), 1)
	assert.Nil(err)
	assert.NotNil(user)
}

func TestQueryFindByIDFailed(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	assert.Nil(err)
	assert.NotNil(db)

	u := NewUserRepository(db)
	user, err := u.FindByID(context.Background(), 0)
	assert.NotNil(err)
	assert.Nil(user)
}

func TestQueryInsertSuccess(t *testing.T) {
	assert := require.New(t)

	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	// var userRepo domain.UserRepository
	// uc := user.NewUserUsecase(userRepo)
	db, err := database.InitDatabase(cfg)

	u := NewUserRepository(db)
	passHashed, err := password.HashAndSalt("password")
	assert.Nil(err)
	err = u.CreateOrUpdate(context.Background(), &domain.User{
		Username: "joko123",
		Name:     "Joko Thinker",
		Password: passHashed,
	})

	assert.Nil(err)
}

func TestQueryInsertfailed(t *testing.T) {
	assert := require.New(t)

	cfg, err := utils.LoadConfig()
	assert.Nil(err)

	db, err := database.InitDatabase(cfg)
	u := NewUserRepository(db)

	err = u.CreateOrUpdate(context.Background(), &domain.User{
		Username: "joko123",
		Name:     "Joko Thinker",
		Password: "jokothinker123",
	})
	assert.NotNil(err)
}

func TestQueryDeleteSuccess(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	u := NewUserRepository(db)
	err = u.Delete(context.Background(), 3)
	assert.Nil(err)
}

func TestQueryDeleteFailed(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	u := NewUserRepository(db)
	err = u.Delete(context.Background(), 0)
	assert.NotNil(err)
}

func TestQueryFindByUsernameSuccess(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	u := NewUserRepository(db)
	user, err := u.FindByUsername(context.Background(), "joko123")
	assert.Nil(err)
	assert.NotNil(user)
}

func TestQueryFindByUsernameFailed(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := database.InitDatabase(cfg)
	u := NewUserRepository(db)
	user, err := u.FindByUsername(context.Background(), "budi123")
	assert.NotNil(err)
	assert.Nil(user)
}
