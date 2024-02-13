package database

import (
	"context"
	"fmt"
	"time"
	"todo-list-app/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase(cfg *domain.Config) (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUserame, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	conn.Config().MaxConns = 30
	conn.Config().MinConns = 5
	conn.Config().MaxConnIdleTime = time.Duration(time.Minute)
	conn.Config().MaxConnLifetime = time.Duration(30 * time.Minute)

	return conn, nil
}
