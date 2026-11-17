package health

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/worksphere/worksphere/apps/api/internal/database"
)

type LivenessResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type ReadinessResponse struct {
	Status    string            `json:"status"`
	Service   string            `json:"service"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
}

type Handler struct {
	db    *database.DB
	redis *database.RedisClient
}

func NewHandler(db *database.DB, redisClient *database.RedisClient) *Handler {
	return &Handler{
		db:    db,
		redis: redisClient,
	}
}

// Health handles GET /health (shallow liveness check).
// Does not falsely report database status.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	resp := LivenessResponse{
		Status:    "ok",
		Service:   "worksphere-api",
		Version:   "0.1.0",
		Timestamp: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// Ready handles GET /ready (deep readiness probe).
// Explicitly checks the database and cache connectivity.
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	checks := make(map[string]string)
	isReady := true

	// Verify database connectivity
	if h.db != nil {
		if err := h.db.Ping(ctx); err != nil {
			checks["database"] = "unavailable: " + err.Error()
			isReady = false
		} else {
			checks["database"] = "connected"
		}
	} else {
		checks["database"] = "uninitialized"
		isReady = false
	}

	// Verify redis connectivity if configured
	if h.redis != nil {
		if err := h.redis.Ping(ctx); err != nil {
			checks["redis"] = "unavailable: " + err.Error()
			// Redis can be degraded or non-blocking depending on configuration, but report status
		} else {
			checks["redis"] = "connected"
		}
	} else {
		checks["redis"] = "not_configured"
	}

	statusText := "ready"
	statusCode := http.StatusOK
	if !isReady {
		statusText = "not_ready"
		statusCode = http.StatusServiceUnavailable
	}

	resp := ReadinessResponse{
		Status:    statusText,
		Service:   "worksphere-api",
		Timestamp: time.Now().UTC(),
		Checks:    checks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}
