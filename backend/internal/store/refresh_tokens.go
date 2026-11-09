package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *Store) SaveRefresh(ctx context.Context, userID int64, hash string, expires time.Time) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`, userID, hash, expires)
	return err
}

func (s *Store) RotateRefresh(ctx context.Context, current, next string, expires time.Time) (User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return User{}, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var id int64
	err = tx.QueryRow(ctx, `
		UPDATE refresh_tokens
		SET revoked_at = now()
		WHERE token_hash = $1
		  AND revoked_at IS NULL
		  AND expires_at > now()
		RETURNING user_id
	`, current).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, err
	}

	user, err := scanUser(tx.QueryRow(ctx, `
		SELECT id, email, password_hash, role
		FROM users
		WHERE id = $1
	`, id))
	if err != nil {
		return User{}, err
	}

	if _, err = tx.Exec(ctx, `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`, id, next, expires); err != nil {
		return User{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Store) RevokeRefresh(ctx context.Context, hash string) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE refresh_tokens
		SET revoked_at = now()
		WHERE token_hash = $1
		  AND revoked_at IS NULL
	`, hash)
	return err
}
