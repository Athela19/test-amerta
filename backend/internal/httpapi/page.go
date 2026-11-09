package httpapi

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type pageQuery struct {
	Query  string
	Page   int
	Limit  int
	Offset int
}

func readPage(r *http.Request) (pageQuery, string) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if utf8.RuneCountInString(q) > 100 {
		return pageQuery{}, "Kata kunci pencarian maksimal 100 karakter."
	}

	page := 1
	if raw := r.URL.Query().Get("page"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > 10000 {
			return pageQuery{}, "Parameter halaman tidak valid."
		}
		page = n
	}

	limit := 10
	if raw := r.URL.Query().Get("limit"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > 100 {
			return pageQuery{}, "Parameter limit tidak valid."
		}
		limit = n
	}

	return pageQuery{
		Query:  strings.ToLower(q),
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}, ""
}
