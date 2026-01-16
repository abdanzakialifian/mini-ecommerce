package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
)

type cartRepositoryImpl struct {
	tx *helper.Transaction
}

func NewCartRepositoryImpl(tx *helper.Transaction) domain.CartRepository {
	return &cartRepositoryImpl{tx: tx}
}

func (c *cartRepositoryImpl) FindByUserId(ctx context.Context, userId int) (model.Cart, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, user_id FROM carts WHERE user_id = $1"
	var cart model.Cart
	err := db.QueryRow(ctx, query, userId).Scan(&cart.ID, &cart.UserID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Cart{}, domain.ErrCartNotFound
		}
		return model.Cart{}, err
	}

	return cart, nil
}

func (c *cartRepositoryImpl) FindOrCreateByUserId(ctx context.Context, userId int) (model.Cart, error) {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO carts (user_id) VALUES ($1) ON CONFLICT (user_id) DO UPDATE SET user_id = EXCLUDED.user_id RETURNING id, user_id"
	var cart model.Cart
	err := db.QueryRow(ctx, query, userId).Scan(&cart.ID, &cart.UserID)
	if err != nil {
		return model.Cart{}, err
	}

	return cart, nil
}
