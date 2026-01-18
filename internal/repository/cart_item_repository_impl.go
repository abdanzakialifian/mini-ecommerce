package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/cart"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
)

type cartItemRepositoryImpl struct {
	tx *helper.Transaction
}

func NewCartItemRepositoryImpl(tx *helper.Transaction) cart.CartItemRepository {
	return &cartItemRepositoryImpl{tx: tx}
}

func (c *cartItemRepositoryImpl) Create(ctx context.Context, cartItem *cart.CartItem) error {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(ctx, query, cartItem.CartID, cartItem.ProductID, cartItem.Quantity).Scan(&cartItem.ID)
	return err
}

func (c *cartItemRepositoryImpl) FindById(ctx context.Context, id int) (cart.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE id = $1"
	var result cart.CartItem
	err := db.QueryRow(ctx, query, id).Scan(&result.ID, &result.CartID, &result.ProductID, &result.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return cart.CartItem{}, helper.ErrCartItemNotFound
		}
		return cart.CartItem{}, err
	}

	return result, nil
}

func (c *cartItemRepositoryImpl) FindAllByCartId(ctx context.Context, cartId int) ([]cart.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 ORDER BY id"
	rows, err := db.Query(ctx, query, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []cart.CartItem
	for rows.Next() {
		var cartItem cart.CartItem
		if err := rows.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
			return nil, err
		}
		results = append(results, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *cartItemRepositoryImpl) FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*cart.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 AND product_id = $2"
	result := &cart.CartItem{}
	err := db.QueryRow(ctx, query, cartId, productId).Scan(&result.ID, &result.CartID, &result.ProductID, &result.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (c *cartItemRepositoryImpl) Update(ctx context.Context, updateCartItem cart.UpdateCartItem) error {
	db := c.tx.GetTx(ctx)
	query := "UPDATE cart_items SET quantity = $1 WHERE id = $2"
	cmd, err := db.Exec(ctx, query, updateCartItem.Quantity, updateCartItem.ID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrCartItemNotFound
	}

	return nil
}

func (c *cartItemRepositoryImpl) Delete(ctx context.Context, id int) error {
	db := c.tx.GetTx(ctx)
	query := "DELETE FROM cart_items WHERE id = $1"
	cmd, err := db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrCartItemNotFound
	}

	return nil
}
