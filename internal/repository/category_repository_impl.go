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

func NewCategoryRepositoryImpl(db *pgxpool.Pool) category.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (c *categoryRepositoryImpl) Create(ctx context.Context, category *category.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	err := c.db.QueryRow(
		ctx,
		query,
		category.Name,
	).Scan(&category.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrCategoryAlreadyExists
		}
		return err
	}

	return nil
}

func (c *categoryRepositoryImpl) Find(ctx context.Context, id string) (category.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = $1"
	var result category.Category
	err := c.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.Name,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return category.Category{}, helper.ErrCategoryNotFound
		}
		return category.Category{}, err
	}

	return result, nil
}

func (c *categoryRepositoryImpl) FindAll(ctx context.Context) ([]category.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []category.Category

	for rows.Next() {
		var category category.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}
		results = append(results, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *categoryRepositoryImpl) Update(ctx context.Context, updateCategory *category.UpdateCategory) error {
	query := "UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING id, name"
	err := c.db.QueryRow(
		ctx,
		query,
		updateCategory.Name,
		updateCategory.ID,
	).Scan(
		&updateCategory.ID,
		&updateCategory.Name,
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
