package repository

import (
	"context"
	"mini-ecommerce/internal/domain/order"
	"mini-ecommerce/internal/helper"
)

type orderItemRepositoryImpl struct {
	tx *helper.Transaction
}

func NewOrderItemRepositoryImpl(tx *helper.Transaction) order.OrderItemRepository {
	return &orderItemRepositoryImpl{tx: tx}
}

func (o *orderItemRepositoryImpl) CreateOrderItems(ctx context.Context, items []order.OrderItem) error {
	db := o.tx.GetTx(ctx)
	query := "INSERT INTO order_items (order_id, product_id, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id"

	for i := range items {
		if err := db.QueryRow(
			ctx,
			query,
			items[i].OrderID,
			items[i].ProductID,
			items[i].Price,
			items[i].Quantity,
		).Scan(&items[i].ID); err != nil {
			return err
		}
	}

	return nil
}

func (o *orderItemRepositoryImpl) FindOrderItems(ctx context.Context, orderId int) ([]order.OrderItem, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, order_id, product_id, price, quantity FROM order_items WHERE order_id = $1"
	rows, err := db.Query(ctx, query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []order.OrderItem

	for rows.Next() {
		var orderItem order.OrderItem
		if err := rows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.Price,
			&orderItem.Quantity,
		); err != nil {
			return nil, err
		}
		results = append(results, orderItem)
	}

	return results, rows.Err()
}
