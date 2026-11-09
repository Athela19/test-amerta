package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func scanUser(row pgx.Row) (User, error) {
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	return user, err
}

func (s *Store) UserByEmail(ctx context.Context, email string) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		SELECT id, email, password_hash, role
		FROM users
		WHERE email = $1
	`, email))
}

func (s *Store) UserByID(ctx context.Context, id int64) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		SELECT id, email, password_hash, role
		FROM users
		WHERE id = $1
	`, id))
}
