package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/cart"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
)

type cartRepositoryImpl struct {
	tx *helper.Transaction
}

func NewCartRepositoryImpl(tx *helper.Transaction) cart.CartRepository {
	return &cartRepositoryImpl{tx: tx}
}

func (c *cartRepositoryImpl) FindByUserId(ctx context.Context, userId int) (cart.Cart, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, user_id FROM carts WHERE user_id = $1"
	var result cart.Cart
	err := db.QueryRow(ctx, query, userId).Scan(&result.ID, &result.UserID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return cart.Cart{}, helper.ErrCartNotFound
		}
		return cart.Cart{}, err
	}

	return result, nil
}

func (c *cartRepositoryImpl) FindOrCreateByUserId(ctx context.Context, userId int) (cart.Cart, error) {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO carts (user_id) VALUES ($1) ON CONFLICT (user_id) DO UPDATE SET user_id = EXCLUDED.user_id RETURNING id, user_id"
	var result cart.Cart
	err := db.QueryRow(ctx, query, userId).Scan(&result.ID, &result.UserID)
	if err != nil {
		return cart.Cart{}, err
	}

	return result, nil
}
