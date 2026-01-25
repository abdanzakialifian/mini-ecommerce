package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/category"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCategory(db *pgxpool.Pool) category.Repository {
	return &categoryRepositoryImpl{db: db}
}

func (c *categoryRepositoryImpl) Create(ctx context.Context, data *category.Data) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	err := c.db.QueryRow(
		ctx,
		query,
		data.Name,
	).Scan(&data.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrCategoryAlreadyExists
		}
		return err
	}

	return nil
}

func (c *categoryRepositoryImpl) Find(ctx context.Context, id string) (category.Data, error) {
	query := "SELECT id, name FROM categories WHERE id = $1"
	var categoryData category.Data
	err := c.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&categoryData.ID,
		&categoryData.Name,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return category.Data{}, helper.ErrCategoryNotFound
		}
		return category.Data{}, err
	}

	return categoryData, nil
}

func (c *categoryRepositoryImpl) FindAll(ctx context.Context) ([]category.Data, error) {
	query := "SELECT id, name FROM categories"
	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category.Data

	for rows.Next() {
		var categoryData category.Data
		if err := rows.Scan(
			&categoryData.ID,
			&categoryData.Name,
		); err != nil {
			return nil, err
		}
		categories = append(categories, categoryData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryRepositoryImpl) Update(ctx context.Context, update *category.Update) error {
	query := "UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING id, name"
	err := c.db.QueryRow(
		ctx,
		query,
		update.Name,
		update.ID,
	).Scan(
		&update.ID,
		&update.Name,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return helper.ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func (c *categoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	cmd, err := c.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrCategoryNotFound
	}

	return nil
}
