package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/order"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
)

type orderRepositoryImpl struct {
	tx *helper.Transaction
}

func NewOrderRepositoryImpl(tx *helper.Transaction) order.OrderRepository {
	return &orderRepositoryImpl{tx: tx}
}

func (o *orderRepositoryImpl) Create(ctx context.Context, order *order.Order) error {
	db := o.tx.GetTx(ctx)
	query := "INSET INTO orders (user_id, total_price, status) VALUES ($1, $2, $3) RETURNING id"
	return db.QueryRow(
		ctx,
		query,
		order.UserID,
		order.TotalPrice,
		order.Status,
	).Scan(&order.ID)
}

func (o *orderRepositoryImpl) FindById(ctx context.Context, id int) (order.Order, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, user_id, total_price, status FROM orders WHERE id = $1"
	var result order.Order
	if err := db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.TotalPrice,
		&result.Status,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order.Order{}, helper.ErrOrderNotFound
		}
		return order.Order{}, err
	}

	return result, nil
}

func (o *orderRepositoryImpl) FindByUserId(ctx context.Context, userId int) ([]order.Order, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, user_id, total_price, status FROM orders WHERE user_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []order.Order
	for rows.Next() {
		var order order.Order
		if err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
		); err != nil {
			return nil, err
		}
		results = append(results, order)
	}

	return results, rows.Err()
}

func (o *orderRepositoryImpl) Update(ctx context.Context, updateOrder *order.UpdateOrder) error {
	db := o.tx.GetTx(ctx)
	query := "UPDATE orders SET total_price = COALESCE($1, total_price), status = COALESCE($2, status), updated_at = NOW() WHERE id = $3"
	cmd, err := db.Exec(
		ctx,
		query,
		updateOrder.TotalPrice,
		updateOrder.Status,
		updateOrder.ID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrOrderNotFound
	}

	return nil
}

func (o *orderRepositoryImpl) UpdateStatus(ctx context.Context, id int, status order.Status) (order.Order, error) {
	db := o.tx.GetTx(ctx)
	query := "UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2 RETURNING id, user_id, total_price, status"
	var result order.Order
	if err := db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.TotalPrice,
		&result.Status,
	); err != nil {
		return order.Order{}, err
	}

	return result, nil
}

func (o *orderRepositoryImpl) Delete(ctx context.Context, id int) error {
	db := o.tx.GetTx(ctx)
	query := "DELETE FROM orders WHERE id = $1"
	cmd, err := db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrOrderNotFound
	}

	return nil
}
