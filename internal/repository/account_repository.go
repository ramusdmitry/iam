package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	StoreAccount(ctx context.Context, email, password string) (string, error)
}

type AccountRepositoryImpl struct {
	db *sqlx.DB
}

func (a *AccountRepositoryImpl) StoreAccount(ctx context.Context, email, password string) (string, error) {
	a.db.Get()
}

func NewAccountRepository(db *sqlx.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{db: db}
}
