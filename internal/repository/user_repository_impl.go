package repository

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/user"
	"mini-ecommerce/internal/helper"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) user.Repository {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) Create(ctx context.Context, data *user.Data) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := u.db.QueryRow(
		ctx,
		query,
		data.Name,
		data.Email,
		data.Password,
	).Scan(&data.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrUserAlreadyExists
		}
		return err
	}

	return nil
}

func (u *userRepositoryImpl) FindByEmail(ctx context.Context, login user.Login) (user.Data, string, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	var userData user.Data
	err := u.db.QueryRow(
		ctx,
		query,
		login.Email,
	).Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.Data{}, "", helper.ErrUserInvalid
		}
		return user.Data{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(login.Password)); err != nil {
		return user.Data{}, "", helper.ErrUserInvalid
	}

	accessToken, err := helper.GenerateAccessToken(userData.ID, userData.Name, userData.Email)

	if err != nil {
		return user.Data{}, "", err
	}

	return userData, accessToken, nil
}

func (u *userRepositoryImpl) FindById(ctx context.Context, id int) (user.Data, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	var userData user.Data
	err := u.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.Data{}, helper.ErrUserNotFound
		}
		return user.Data{}, err
	}

	return userData, nil
}

func (u *userRepositoryImpl) Update(ctx context.Context, update *user.Update) error {
	query := "UPDATE users SET name = COALESCE($1, name), email = COALESCE($2, email), password = COALESCE($3, password) WHERE id = $4 RETURNING id, name, email"
	err := u.db.QueryRow(
		ctx,
		query,
		update.Name,
		update.Email,
		update.NewPassword,
		update.ID,
	).Scan(
		&update.ID,
		&update.Name,
		&update.Email,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return helper.ErrUserAlreadyExists
		}

		if errors.Is(err, pgx.ErrNoRows) {
			return helper.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (u *userRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	cmd, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return helper.ErrUserNotFound
	}

	return nil
}
