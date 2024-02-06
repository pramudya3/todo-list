package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase() (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dsn := "postgres://postgres:postgres@localhost:5433/postgres"

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	conn.Config().MaxConns = 30
	conn.Config().MinConns = 5
	conn.Config().MaxConnIdleTime = time.Duration(time.Minute)
	conn.Config().MaxConnLifetime = time.Duration(30 * time.Minute)

	return conn, nil
}
