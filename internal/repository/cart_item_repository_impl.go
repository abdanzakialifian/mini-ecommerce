package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
)

type cartItemRepositoryImpl struct {
	tx *helper.Transaction
}

func NewCartItemRepositoryImpl(tx *helper.Transaction) domain.CartItemRepository {
	return &cartItemRepositoryImpl{tx: tx}
}

func (c *cartItemRepositoryImpl) Create(ctx context.Context, cartItem *model.CartItem) error {
	db := c.tx.GetTx(ctx)
	query := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(ctx, query, cartItem.CartID, cartItem.ProductID, cartItem.Quantity).Scan(&cartItem.ID)
	return err
}

func (c *cartItemRepositoryImpl) FindById(ctx context.Context, id int) (model.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE id = $1"
	var cartItem model.CartItem
	err := db.QueryRow(ctx, query, id).Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.CartItem{}, domain.ErrCartItemNotFound
		}
		return model.CartItem{}, err
	}

	return cartItem, nil
}

func (c *cartItemRepositoryImpl) FindAllByCartId(ctx context.Context, cartId int) ([]model.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 ORDER BY id"
	rows, err := db.Query(ctx, query, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []model.CartItem
	for rows.Next() {
		var cartItem model.CartItem
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

func (c *cartItemRepositoryImpl) FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*model.CartItem, error) {
	db := c.tx.GetTx(ctx)
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = $1 AND product_id = $2"
	cartItem := &model.CartItem{}
	err := db.QueryRow(ctx, query, cartId, productId).Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return cartItem, nil
}

func (c *cartItemRepositoryImpl) Update(ctx context.Context, updateCartItem model.UpdateCartItem) error {
	db := c.tx.GetTx(ctx)
	query := "UPDATE cart_items SET quantity = $1 WHERE id = $2"
	cmd, err := db.Exec(ctx, query, updateCartItem.Quantity, updateCartItem.ID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return domain.ErrCartItemNotFound
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
		return domain.ErrCartItemNotFound
	}

	return nil
}
