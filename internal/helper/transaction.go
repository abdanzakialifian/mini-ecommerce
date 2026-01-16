package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txKey struct{}

type Transactor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Transaction struct {
	db *pgxpool.Pool
}

func NewTransaction(db *pgxpool.Pool) *Transaction {
	return &Transaction{db: db}
}

func (s *Transaction) ExecTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	ctxWithTx := context.WithValue(ctx, txKey{}, tx)

	err = fn(ctxWithTx)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit(ctx)
}

func (s *Transaction) GetTx(ctx context.Context) Transactor {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return s.db
}
