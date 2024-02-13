package database

import (
	"context"
	"fmt"
	"testing"
	"todo-list-app/domain"
	"todo-list-app/internal/utils"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDbSuccess(t *testing.T) {
	assert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(err)
	db, err := InitDatabase(cfg)
	assert.Nil(err)
	assert.NotNil(db)
}

func TestQuery(t *testing.T) {
	cfg, err := utils.LoadConfig()
	require.Nil(t, err)
	db, _ := InitDatabase(cfg)
	users := []*domain.User{}
	err = pgxscan.Select(context.Background(), db, &users, `SELECT * FROM users`)
	require.Nil(t, err)
	for _, v := range users {
		fmt.Printf("user: %v\n", v)
	}
}

func TestDbFailed(t *testing.T) {
	asssert := require.New(t)
	cfg, err := utils.LoadConfig()
	assert.Nil(t, err)
	db, err := InitDatabase(cfg)
	asssert.NotNil(err)
	asssert.Nil(db)
}
