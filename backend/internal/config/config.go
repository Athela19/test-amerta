package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const DevSecret = "dev-only-jwt-secret-change-me-0123456789abcdef"

type Config struct {
	Addr         string
	DatabaseURL  string
	JWTSecret    string
	Origins      []string
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
	CookieSecure bool
}

func Load() Config {
	return Config{
		Addr:         ":" + env("PORT", "8080"),
		DatabaseURL:  env("DATABASE_URL", "postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable"),
		JWTSecret:    env("JWT_SECRET", DevSecret),
		Origins:      envList("CORS_ALLOWED_ORIGIN", "http://localhost:3000,http://localhost:5173"),
		AccessTTL:    envDuration("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTTL:   envDuration("REFRESH_TOKEN_TTL", 7*24*time.Hour),
		CookieSecure: envBool("COOKIE_SECURE", false),
	}
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func envList(key, fallback string) []string {
	parts := strings.Split(env(key, fallback), ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		if value := strings.TrimSpace(part); value != "" {
			values = append(values, value)
		}
	}
	return values
}

func envDuration(key string, fallback time.Duration) time.Duration {
	raw := strings.TrimSpace(os.Getenv(key))
	value, err := time.ParseDuration(raw)
	if err != nil || value <= 0 {
		return fallback
	}
	return value
}

func envBool(key string, fallback bool) bool {
	raw := strings.TrimSpace(os.Getenv(key))
	value, err := strconv.ParseBool(raw)
	if err != nil {
		return fallback
	}
	return value
}
