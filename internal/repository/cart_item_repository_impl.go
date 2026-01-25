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

func NewCartItem(tx *helper.Transaction) cart.ItemRepository {
	return &cartItemRepositoryImpl{tx: tx}
}

func (c *cartItemRepositoryImpl) Create(ctx context.Context, item *cart.Item) error {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(ctx, query, item.CartID, item.ProductID, item.Quantity).Scan(&item.ID)
	return err
}

func (c *cartItemRepositoryImpl) FindById(ctx context.Context, itemId int) (cart.Item, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE id = $1"
	var cartItem cart.Item
	err := db.QueryRow(ctx, query, itemId).Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return cart.Item{}, helper.ErrCartItemNotFound
		}
		return cart.Item{}, err
	}

	return cartItem, nil
}

func (c *cartItemRepositoryImpl) FindAllByCartId(ctx context.Context, cartId int) ([]cart.Item, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 ORDER BY id"
	rows, err := db.Query(ctx, query, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []cart.Item
	for rows.Next() {
		var cartItem cart.Item
		if err := rows.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (c *cartItemRepositoryImpl) FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*cart.Item, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 AND product_id = $2"
	cartItem := &cart.Item{}
	err := db.QueryRow(ctx, query, cartId, productId).Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return cartItem, nil
}

func (c *cartItemRepositoryImpl) Update(ctx context.Context, updateItem cart.UpdateItem) error {
	db := c.tx.GetTx(ctx)
	query := "UPDATE cart_items SET quantity = $1, updated_at = NOW() WHERE id = $2"
	cmd, err := db.Exec(ctx, query, updateItem.Quantity, updateItem.ID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrCartItemNotFound
	}

	return nil
}

func (c *cartItemRepositoryImpl) Delete(ctx context.Context, itemId int) error {
	db := c.tx.GetTx(ctx)
	query := "DELETE FROM cart_items WHERE id = $1"
	cmd, err := db.Exec(ctx, query, itemId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrCartItemNotFound
	}

	return nil
}
