package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"amerta/backend/internal/auth"
	"amerta/backend/internal/store"
)

const (
	accessName  = "amerta_access"
	refreshName = "amerta_refresh"
	refreshPath = "/api/auth"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userView struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func toView(user store.User) userView {
	return userView{ID: user.ID, Email: user.Email, Role: user.Role}
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	var in loginInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "Permintaan tidak valid.")
		return
	}

	email := strings.ToLower(strings.TrimSpace(in.Email))
	if email == "" || in.Password == "" {
		writeError(w, http.StatusUnprocessableEntity, "Email dan kata sandi wajib diisi.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.UserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusUnauthorized, "Email atau kata sandi salah.")
			return
		}
		s.logger.Error("lookup user", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memproses login.")
		return
	}

	if !auth.Verify(user.PasswordHash, in.Password) {
		writeError(w, http.StatusUnauthorized, "Email atau kata sandi salah.")
		return
	}

	access, err := s.tokens.Issue(user.ID, user.Role)
	if err != nil {
		s.logger.Error("issue access", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal membuat sesi.")
		return
	}

	refresh, err := auth.NewRefresh()
	if err != nil {
		s.logger.Error("create refresh", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal membuat sesi.")
		return
	}

	if err := s.store.SaveRefresh(ctx, user.ID, auth.HashRefresh(refresh), time.Now().Add(s.refreshTTL)); err != nil {
		s.logger.Error("save refresh", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal membuat sesi.")
		return
	}

	s.setCookies(w, access, refresh)
	writeJSON(w, http.StatusOK, map[string]any{"user": toView(user)})
}

func (s *server) refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(refreshName)
	if err != nil || cookie.Value == "" {
		s.setCookies(w, "", "")
		writeError(w, http.StatusUnauthorized, "Sesi tidak valid.")
		return
	}

	next, err := auth.NewRefresh()
	if err != nil {
		s.logger.Error("create refresh", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui sesi.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.RotateRefresh(ctx, auth.HashRefresh(cookie.Value), auth.HashRefresh(next), time.Now().Add(s.refreshTTL))
	if err != nil {
		s.setCookies(w, "", "")
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusUnauthorized, "Sesi tidak valid.")
			return
		}
		s.logger.Error("rotate refresh", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui sesi.")
		return
	}

	access, err := s.tokens.Issue(user.ID, user.Role)
	if err != nil {
		s.logger.Error("issue access", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui sesi.")
		return
	}

	s.setCookies(w, access, next)
	writeJSON(w, http.StatusOK, map[string]any{"user": toView(user)})
}

func (s *server) logout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(refreshName); err == nil && cookie.Value != "" {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		if err := s.store.RevokeRefresh(ctx, auth.HashRefresh(cookie.Value)); err != nil {
			s.logger.Error("revoke refresh", "error", err)
		}
		cancel()
	}

	s.setCookies(w, "", "")
	w.WriteHeader(http.StatusNoContent)
}

func (s *server) me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(claimsKey{}).(*auth.Claims)
	if !ok {
		writeError(w, http.StatusUnauthorized, "Akses ditolak.")
		return
	}

	id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "Akses ditolak.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.UserByID(ctx, id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusUnauthorized, "Akses ditolak.")
			return
		}
		s.logger.Error("lookup user", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memuat sesi.")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"user": toView(user)})
}

type claimsKey struct{}

func (s *server) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(accessName)
		if err != nil || cookie.Value == "" {
			writeError(w, http.StatusUnauthorized, "Akses ditolak.")
			return
		}

		claims, err := s.tokens.Parse(cookie.Value)
		if err != nil {
			writeError(w, http.StatusUnauthorized, "Akses ditolak.")
			return
		}

		next(w, r.WithContext(context.WithValue(r.Context(), claimsKey{}, claims)))
	}
}

func (s *server) setCookies(w http.ResponseWriter, access, refresh string) {
	writeCookie(w, accessName, access, "/", s.accessTTL, s.secure)
	writeCookie(w, refreshName, refresh, refreshPath, s.refreshTTL, s.secure)
}

func writeCookie(w http.ResponseWriter, name, value, path string, ttl time.Duration, secure bool) {
	maxAge := int(ttl.Seconds())
	expires := time.Now().Add(ttl)
	if value == "" {
		maxAge = -1
		expires = time.Unix(1, 0)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		MaxAge:   maxAge,
		Expires:  expires,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
