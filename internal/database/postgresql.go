package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := "postgres://postgres:postgres123@localhost:3030/mini_ecommerce"
	if dsn == "" {
		return nil, fmt.Errorf("Database url is not set")
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("Create pool : %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Ping db : %w", err)
	}

	return pool, nil
}
