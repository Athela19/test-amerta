package httpapi

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"amerta/backend/internal/auth"
	"amerta/backend/internal/store"
)

type Deps struct {
	Pool       *pgxpool.Pool
	Tokens     *auth.Manager
	Logger     *slog.Logger
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Secure     bool
	Origins    []string
}

type server struct {
	pool       *pgxpool.Pool
	store      *store.Store
	tokens     *auth.Manager
	logger     *slog.Logger
	accessTTL  time.Duration
	refreshTTL time.Duration
	secure     bool
}

func NewRouter(d Deps) http.Handler {
	s := &server{
		pool:       d.Pool,
		store:      store.New(d.Pool),
		tokens:     d.Tokens,
		logger:     d.Logger,
		accessTTL:  d.AccessTTL,
		refreshTTL: d.RefreshTTL,
		secure:     d.Secure,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/auth/login", s.login)
	mux.HandleFunc("POST /api/auth/refresh", s.refresh)
	mux.HandleFunc("POST /api/auth/logout", s.logout)
	mux.HandleFunc("GET /api/auth/me", s.requireAuth(s.me))
	mux.HandleFunc("GET /api/health", s.health)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusNotFound, "Rute tidak ditemukan.")
	})

	var handler http.Handler = mux
	handler = withRecovery(d.Logger, handler)
	handler = withLogging(d.Logger, handler)
	return withCORS(d.Origins, handler)
}
