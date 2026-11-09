package store

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	pool *pgxpool.Pool
}

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}
