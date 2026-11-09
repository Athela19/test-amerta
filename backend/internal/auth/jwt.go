package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "amerta-api"

type Manager struct {
	key []byte
	ttl time.Duration
}

func New(secret string, ttl time.Duration) *Manager {
	return &Manager{key: []byte(secret), ttl: ttl}
}

func (m *Manager) Issue(id int64) (string, error) {
	now := time.Now()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(id, 10),
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
	}).SignedString(m.key)
}

func (m *Manager) Parse(raw string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(
		raw,
		claims,
		func(*jwt.Token) (any, error) { return m.key, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
