package repository

import (
	"context"
	"mini-ecommerce/internal/domain/order"
	"mini-ecommerce/internal/helper"
)

type orderItemRepositoryImpl struct {
	tx *helper.Transaction
}

func NewOrderItem(tx *helper.Transaction) order.ItemRepository {
	return &orderItemRepositoryImpl{tx: tx}
}

func (o *orderItemRepositoryImpl) CreateItems(ctx context.Context, items []order.Item) error {
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

func (o *orderItemRepositoryImpl) FindItems(ctx context.Context, orderId int) ([]order.Item, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, order_id, product_id, price, quantity FROM order_items WHERE order_id = $1"
	rows, err := db.Query(ctx, query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []order.Item
	for rows.Next() {
		var orderItem order.Item
		if err := rows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.Price,
			&orderItem.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, orderItem)
	}

	return items, rows.Err()
}
