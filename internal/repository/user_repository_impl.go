package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

func (u userRepositoryImpl) FindByEmail(ctx context.Context, login model.LoginUser) (model.User, string, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	var user model.User
	err := u.db.QueryRow(
		ctx,
		query,
		login.Email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, "", domain.ErrUserInvalid
		}
		return model.User{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return model.User{}, "", domain.ErrUserInvalid
	}

	accessToken, err := helper.GenerateAccessToken(user.ID, user.Name, user.Email)

	if err != nil {
		return model.User{}, "", err
	}

	return user, accessToken, nil
}

func (u userRepositoryImpl) FindById(ctx context.Context, id int) (model.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	var user model.User
	err := u.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
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
		updateUser.NewPassword,
		updateUser.ID,
	).Scan(
		&updateUser.ID,
		&updateUser.Name,
		&updateUser.Email,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrUserAlreadyExists
		}

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
