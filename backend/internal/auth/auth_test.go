package auth

import (
	"testing"
	"time"
)

func TestPassword(t *testing.T) {
	hash, err := Hash("admin123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}
	if !Verify(hash, "admin123") {
		t.Fatal("Verify() = false for valid password")
	}
	if Verify(hash, "wrong") {
		t.Fatal("Verify() = true for wrong password")
	}
}

// Hash di bawah disalin dari db/init/001_init.sql. Gagal = password demo di
// README/docs tidak lagi cocok dengan seed.
func TestSeedPasswords(t *testing.T) {
	cases := []struct {
		username string
		password string
		hash     string
	}{
		{
			username: "admin",
			password: "admin123",
			hash:     "$2a$12$5chNCxFg0zvnneYqbXd8w.9hRhT48ekpq0WFHIXvQOaBOuamq/3c6",
		},
		{
			username: "user",
			password: "user123",
			hash:     "$2a$12$zrJxSpywBOzUX.y3sqUntOIXS4jv9kSrs65VQq0SWS6zs9qvdOXHG",
		},
	}

	for _, tc := range cases {
		if !Verify(tc.hash, tc.password) {
			t.Fatalf("seed password for %s no longer matches db/init/001_init.sql", tc.username)
		}
		if Verify(tc.hash, "wrong") {
			t.Fatalf("seed hash for %s accepted a wrong password", tc.username)
		}
	}
}

func TestIssueAndParse(t *testing.T) {
	m := New("test-secret-at-least-32-characters", time.Minute)

	token, err := m.Issue(42)
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	claims, err := m.Parse(token)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if claims.Subject != "42" {
		t.Fatalf("Subject = %q, want %q", claims.Subject, "42")
	}
}

func TestWrongSecret(t *testing.T) {
	token, err := New("first-secret-at-least-32-characters", time.Minute).Issue(1)
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	if _, err := New("second-secret-at-least-32-characters", time.Minute).Parse(token); err == nil {
		t.Fatal("Parse() accepted a token signed with another secret")
	}
}

func TestExpired(t *testing.T) {
	m := New("test-secret-at-least-32-characters", -time.Minute)

	token, err := m.Issue(1)
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	if _, err := m.Parse(token); err == nil {
		t.Fatal("Parse() accepted an expired token")
	}
}

func TestRefresh(t *testing.T) {
	first, err := NewRefresh()
	if err != nil {
		t.Fatalf("NewRefresh() error = %v", err)
	}
	second, err := NewRefresh()
	if err != nil {
		t.Fatalf("NewRefresh() error = %v", err)
	}

	if first == "" || first == second {
		t.Fatal("NewRefresh() did not return a unique token")
	}
	if len(HashRefresh(first)) != 64 {
		t.Fatal("HashRefresh() returned an unexpected length")
	}
	if HashRefresh(first) == HashRefresh(second) {
		t.Fatal("HashRefresh() returned the same hash for different tokens")
	}
}
