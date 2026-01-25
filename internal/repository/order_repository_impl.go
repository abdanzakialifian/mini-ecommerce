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

func NewOrder(tx *helper.Transaction) order.Repository {
	return &orderRepositoryImpl{tx: tx}
}

func (o *orderRepositoryImpl) Create(ctx context.Context, data *order.Data) error {
	db := o.tx.GetTx(ctx)
	query := "INSET INTO orders (user_id, total_price, status) VALUES ($1, $2, $3) RETURNING id"
	return db.QueryRow(
		ctx,
		query,
		data.UserID,
		data.TotalPrice,
		data.Status,
	).Scan(&data.ID)
}

func (o *orderRepositoryImpl) FindById(ctx context.Context, id int) (order.Data, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, user_id, total_price, status FROM orders WHERE id = $1"
	var orderData order.Data
	if err := db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&orderData.ID,
		&orderData.UserID,
		&orderData.TotalPrice,
		&orderData.Status,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order.Data{}, helper.ErrOrderNotFound
		}
		return order.Data{}, err
	}

	return orderData, nil
}

func (o *orderRepositoryImpl) FindByUserId(ctx context.Context, userId int) ([]order.Data, error) {
	db := o.tx.GetTx(ctx)
	query := "SELECT id, user_id, total_price, status FROM orders WHERE user_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order.Data
	for rows.Next() {
		var orderData order.Data
		if err := rows.Scan(
			&orderData.ID,
			&orderData.UserID,
			&orderData.TotalPrice,
			&orderData.Status,
		); err != nil {
			return nil, err
		}
		orders = append(orders, orderData)
	}

	return orders, rows.Err()
}

func (o *orderRepositoryImpl) Update(ctx context.Context, update *order.Update) error {
	db := o.tx.GetTx(ctx)
	query := "UPDATE orders SET total_price = COALESCE($1, total_price), status = COALESCE($2, status), updated_at = NOW() WHERE id = $3"
	cmd, err := db.Exec(
		ctx,
		query,
		update.TotalPrice,
		update.Status,
		update.ID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrOrderNotFound
	}

	return nil
}

func (o *orderRepositoryImpl) UpdateStatus(ctx context.Context, id int, status order.Status) error {
	db := o.tx.GetTx(ctx)
	query := "UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2"
	cmd, err := db.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrOrderNotFound
	}

	return nil
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
