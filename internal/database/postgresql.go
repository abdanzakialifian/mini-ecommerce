package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dsn := "postgres://postgres:postgres123@localhost:3030/mini_ecommerce"
	if dsn == "" {
		return nil, fmt.Errorf("Database url is not set")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect database : %w", err)
	}

	return pool, pool.Ping(context.Background())
}
