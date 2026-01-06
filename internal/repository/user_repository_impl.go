package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepositoryImpl(db *pgxpool.Pool) domain.UserRepository {
	return userRepositoryImpl{db: db}
}

func (u userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := u.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrUserAlreadyExists
		}
		return err
	}

	return nil
}

func (u userRepositoryImpl) Find(ctx context.Context, id int) (model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	var user model.User
	err := u.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, domain.ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func (u userRepositoryImpl) Update(ctx context.Context, updateUser *model.UpdateUser) error {
	query := "UPDATE users SET name = COALESCE($1, name), email = COALESCE($2, email), password = COALESCE($3, password) WHERE id = $4 RETURNING id, name, email"
	err := u.db.QueryRow(
		ctx,
		query,
		updateUser.Name,
		updateUser.Email,
		updateUser.Password,
		updateUser.ID,
	).Scan(
		&updateUser.ID,
		&updateUser.Name,
		&updateUser.Email,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (u userRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	cmd, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
