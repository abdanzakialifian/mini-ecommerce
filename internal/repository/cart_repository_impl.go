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

func NewCart(tx *helper.Transaction) cart.Repository {
	return &cartRepositoryImpl{tx: tx}
}

func (c *cartRepositoryImpl) FindByUserId(ctx context.Context, userId int) (cart.Data, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, user_id FROM carts WHERE user_id = $1"
	var cartData cart.Data
	err := db.QueryRow(ctx, query, userId).Scan(&cartData.ID, &cartData.UserID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return cart.Data{}, helper.ErrCartNotFound
		}
		return cart.Data{}, err
	}

	return cartData, nil
}

func (c *cartRepositoryImpl) FindOrCreateByUserId(ctx context.Context, userId int) (cart.Data, error) {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO carts (user_id) VALUES ($1) ON CONFLICT (user_id) DO UPDATE SET user_id = EXCLUDED.user_id RETURNING id, user_id"
	var cartData cart.Data
	err := db.QueryRow(ctx, query, userId).Scan(&cartData.ID, &cartData.UserID)
	if err != nil {
		return cart.Data{}, err
	}

	return cartData, nil
}
