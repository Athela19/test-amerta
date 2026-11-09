package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func scanUser(row pgx.Row) (User, error) {
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	return user, err
}

func (s *Store) UserByUsername(ctx context.Context, username string) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		SELECT id, username, password_hash, role
		FROM users
		WHERE username = $1
	`, username))
}

func (s *Store) UserByID(ctx context.Context, id int64) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		SELECT id, username, password_hash, role
		FROM users
		WHERE id = $1
	`, id))
}

const userMatch = `
	$1 = ''
	OR position($1 in lower(username)) > 0
	OR position($1 in lower(role)) > 0
	OR position($1 in id::text) > 0`

func (s *Store) Users(ctx context.Context, query string, limit, offset int) ([]User, int, error) {
	var total int
	err := s.pool.QueryRow(ctx, `
		SELECT COUNT(*)::int
		FROM users
		WHERE `+userMatch, query).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT id, username, password_hash, role
		FROM users
		WHERE `+userMatch+`
		ORDER BY id
		LIMIT $2 OFFSET $3
	`, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}
	return users, total, rows.Err()
}

func (s *Store) CreateUser(ctx context.Context, username, hash, role string) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		INSERT INTO users (username, password_hash, role)
		VALUES ($1, $2, $3)
		RETURNING id, username, password_hash, role
	`, username, hash, role))
}

func (s *Store) UpdateRole(ctx context.Context, id int64, role string) (User, error) {
	return scanUser(s.pool.QueryRow(ctx, `
		UPDATE users
		SET role = $2
		WHERE id = $1
		RETURNING id, username, password_hash, role
	`, id, role))
}

func (s *Store) DeleteUser(ctx context.Context, id int64) error {
	tag, err := s.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
