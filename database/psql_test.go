package database

import (
	"context"
	"fmt"
	"testing"
	"todo-list-app/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/stretchr/testify/require"
)

func TestDbSuccess(t *testing.T) {
	assert := require.New(t)
	conn, err := InitDatabase()
	assert.Nil(err)
	assert.NotNil(conn)
}

func TestQuery(t *testing.T) {
	conn, _ := InitDatabase()
	users := []*domain.User{}
	err := pgxscan.Select(context.Background(), conn, &users, `SELECT * FROM users`)
	require.Nil(t, err)
	for _, v := range users {
		fmt.Printf("user: %v\n", v)
	}
}

func TestDbFailed(t *testing.T) {
	asssert := require.New(t)
	conn, err := InitDatabase()
	asssert.NotNil(err)
	asssert.Nil(conn)
}
