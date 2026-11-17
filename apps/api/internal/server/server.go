package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/worksphere/worksphere/apps/api/internal/config"
	"github.com/worksphere/worksphere/apps/api/internal/database"
	"github.com/worksphere/worksphere/apps/api/internal/health"
	appmw "github.com/worksphere/worksphere/apps/api/internal/middleware"
)

type Server struct {
	cfg         *config.Config
	db          *database.DB
	redisClient *database.RedisClient
	router      *chi.Mux
}

func New(cfg *config.Config, db *database.DB, redisClient *database.RedisClient) *Server {
	s := &Server{
		cfg:         cfg,
		db:          db,
		redisClient: redisClient,
		router:      chi.NewRouter(),
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.router.Use(chimw.RealIP)
	s.router.Use(appmw.RequestID)
	s.router.Use(appmw.StructuredLogger)
	s.router.Use(appmw.Recoverer)

	// Configure CORS
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.cfg.CorsAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID", "X-Organization-ID", "X-User-ID"},
		ExposedHeaders:   []string{"Link", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func (s *Server) setupRoutes() {
	healthHandler := health.NewHandler(s.db, s.redisClient)

	// Liveness & Readiness Probes
	s.router.Get("/health", healthHandler.Health)
	s.router.Get("/ready", healthHandler.Ready)

	// Root status
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"name":         "WorkSphere API",
			"description":  "Sri Lanka-first multi-tenant business operations platform",
			"status":       "operational",
			"architecture": "modular-monolith",
			"version":      "0.1.0",
			"docs":         "/docs",
		})
	})

	// API v1 Namespace
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Use(appmw.TenantContextExtractor(s.cfg.EnableMockAuth))

		// Module Route Anchors (Placeholder handlers illustrating domain routing)
		r.Get("/modules", func(w http.ResponseWriter, r *http.Request) {
			tc, _ := appmw.GetTenantContext(r.Context())
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"organization_id": tc.OrganizationID,
				"modules": []string{
					"organizations",
					"projects",
					"tasks",
					"workforce",
					"leave",
					"crm",
					"finance",
					"billing",
					"notifications",
					"email",
					"audit",
				},
				"currency_default": "LKR",
				"timezone_default": "Asia/Colombo",
			})
		})
	})
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func (s *Server) HTTPServer() *http.Server {
	return &http.Server{
		Addr:         ":" + s.cfg.Port,
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
