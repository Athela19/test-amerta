package httpapi

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"

	"amerta/backend/internal/auth"
	"amerta/backend/internal/store"
)

var usernamePattern = regexp.MustCompile(`^[a-z0-9._-]{3,32}$`)

type createUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type roleInput struct {
	Role string `json:"role"`
}

func (s *server) listUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	users, err := s.store.Users(ctx)
	if err != nil {
		s.logger.Error("list users", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memuat pengguna.")
		return
	}

	list := make([]userView, 0, len(users))
	for _, user := range users {
		list = append(list, toView(user))
	}
	writeJSON(w, http.StatusOK, map[string]any{"users": list})
}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	var in createUserInput
	if err := decode(w, r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "Permintaan tidak valid.")
		return
	}

	username := strings.ToLower(strings.TrimSpace(in.Username))
	if !usernamePattern.MatchString(username) {
		writeError(w, http.StatusUnprocessableEntity, "Username harus 3-32 karakter: huruf, angka, titik, garis bawah, atau strip.")
		return
	}
	if len(in.Password) < 6 || len(in.Password) > 72 {
		writeError(w, http.StatusUnprocessableEntity, "Kata sandi harus 6-72 karakter.")
		return
	}
	if !validRole(in.Role) {
		writeError(w, http.StatusUnprocessableEntity, "Peran tidak valid.")
		return
	}

	hash, err := auth.Hash(in.Password)
	if err != nil {
		s.logger.Error("hash password", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal menambah pengguna.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.CreateUser(ctx, username, hash, in.Role)
	if err != nil {
		if isUnique(err) {
			writeError(w, http.StatusConflict, "Username sudah digunakan.")
			return
		}
		s.logger.Error("create user", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal menambah pengguna.")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{"user": toView(user)})
}

func (s *server) updateRole(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "ID pengguna tidak valid.")
		return
	}

	var in roleInput
	if err := decode(w, r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "Permintaan tidak valid.")
		return
	}
	if !validRole(in.Role) {
		writeError(w, http.StatusUnprocessableEntity, "Peran tidak valid.")
		return
	}

	current, _ := currentUser(r)
	if current.ID == id {
		writeError(w, http.StatusUnprocessableEntity, "Tidak dapat mengubah peran sendiri.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.UpdateRole(ctx, id, in.Role)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Pengguna tidak ditemukan.")
			return
		}
		s.logger.Error("update role", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal mengubah peran.")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"user": toView(user)})
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "ID pengguna tidak valid.")
		return
	}

	current, _ := currentUser(r)
	if current.ID == id {
		writeError(w, http.StatusUnprocessableEntity, "Tidak dapat menghapus akun sendiri.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := s.store.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Pengguna tidak ditemukan.")
			return
		}
		s.logger.Error("delete user", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal menghapus pengguna.")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func pathID(r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return id, err == nil && id > 0
}

func validRole(role string) bool {
	return role == "admin" || role == "user"
}

func isUnique(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
