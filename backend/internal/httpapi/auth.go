package httpapi

import (
	"context"
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
	Username string `json:"username"`
	Password string `json:"password"`
}

type userView struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func toView(user store.User) userView {
	return userView{ID: user.ID, Username: user.Username, Role: user.Role}
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	var in loginInput
	if err := decode(w, r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "Permintaan tidak valid.")
		return
	}

	username := strings.ToLower(strings.TrimSpace(in.Username))
	if username == "" || in.Password == "" {
		writeError(w, http.StatusUnprocessableEntity, "Username dan kata sandi wajib diisi.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := s.store.UserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusUnauthorized, "Username atau kata sandi salah.")
			return
		}
		s.logger.Error("lookup user", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memproses login.")
		return
	}

	if !auth.Verify(user.PasswordHash, in.Password) {
		writeError(w, http.StatusUnauthorized, "Username atau kata sandi salah.")
		return
	}

	access, err := s.tokens.Issue(user.ID)
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

	access, err := s.tokens.Issue(user.ID)
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
	user, ok := currentUser(r)
	if !ok {
		writeError(w, http.StatusUnauthorized, "Akses ditolak.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"user": toView(user)})
}

type userKey struct{}

func currentUser(r *http.Request) (store.User, bool) {
	user, ok := r.Context().Value(userKey{}).(store.User)
	return user, ok
}

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

		next(w, r.WithContext(context.WithValue(r.Context(), userKey{}, user)))
	}
}

func (s *server) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return s.requireAuth(func(w http.ResponseWriter, r *http.Request) {
		user, ok := currentUser(r)
		if !ok || user.Role != "admin" {
			writeError(w, http.StatusForbidden, "Akses ditolak.")
			return
		}
		next(w, r)
	})
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
