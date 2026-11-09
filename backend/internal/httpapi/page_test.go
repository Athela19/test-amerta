package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadPageDefaults(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	got, msg := readPage(req)
	if msg != "" {
		t.Fatal(msg)
	}
	if got.Query != "" || got.Page != 1 || got.Limit != 10 || got.Offset != 0 {
		t.Fatalf("defaults: %+v", got)
	}
}

func TestReadPageQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/users?q=%20EnG%20&page=3&limit=25", nil)
	got, msg := readPage(req)
	if msg != "" {
		t.Fatal(msg)
	}
	if got.Query != "eng" || got.Page != 3 || got.Limit != 25 || got.Offset != 50 {
		t.Fatalf("parsed: %+v", got)
	}
}

func TestReadPageRejects(t *testing.T) {
	cases := []string{
		"/api/users?page=0",
		"/api/users?page=x",
		"/api/users?limit=0",
		"/api/users?limit=101",
		"/api/users?q=" + strings.Repeat("a", 101),
	}
	for _, raw := range cases {
		req := httptest.NewRequest(http.MethodGet, raw, nil)
		if _, msg := readPage(req); msg == "" {
			t.Fatalf("expected error for %s", raw)
		}
	}
}
