package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	dsn := "postgres://postgres:postgres123@localhost:3030/mini_ecommerce"
	if dsn == "" {
		return nil, fmt.Errorf("Database url is not set")
	}

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect database : %w", err)
	}

	return conn, nil
}
