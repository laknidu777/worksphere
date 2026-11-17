package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/worksphere/worksphere/apps/api/internal/config"
	"github.com/worksphere/worksphere/apps/api/internal/server"
)

func setupTestServer() *server.Server {
	cfg := &config.Config{
		Port:                   "8080",
		Environment:            "test",
		EnableMockAuth:         true,
		CorsAllowedOrigins:     []string{"*"},
		ShutdownTimeoutSeconds: 5,
	}
	return server.New(cfg, nil, nil)
}

func TestHealthCheck(t *testing.T) {
	srv := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Router().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", body["status"])
	}
	if body["service"] != "worksphere-api" {
		t.Errorf("expected service 'worksphere-api', got %v", body["service"])
	}
}

func TestReadinessWithoutDatabase(t *testing.T) {
	srv := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()

	srv.Router().ServeHTTP(rec, req)

	// When database is not initialized, readiness probe MUST return 503
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["status"] != "not_ready" {
		t.Errorf("expected status 'not_ready', got %v", body["status"])
	}
}

func TestModulesEndpointSriLankanDefaults(t *testing.T) {
	srv := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/modules", nil)
	rec := httptest.NewRecorder()

	srv.Router().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["currency_default"] != "LKR" {
		t.Errorf("expected currency default 'LKR', got %v", body["currency_default"])
	}
	if body["timezone_default"] != "Asia/Colombo" {
		t.Errorf("expected timezone default 'Asia/Colombo', got %v", body["timezone_default"])
	}
}
