package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	h := NewHandler(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	h.Health(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp LivenessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode json response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", resp.Status)
	}

	if resp.Service != "worksphere-api" {
		t.Errorf("expected service 'worksphere-api', got '%s'", resp.Service)
	}
}

func TestReadyEndpoint_UninitializedDB(t *testing.T) {
	h := NewHandler(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rr := httptest.NewRecorder()

	h.Ready(rr, req)

	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503 for uninitialized db, got %d", rr.Code)
	}

	var resp ReadinessResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode json response: %v", err)
	}

	if resp.Status != "not_ready" {
		t.Errorf("expected status 'not_ready', got '%s'", resp.Status)
	}

	if resp.Checks["database"] != "uninitialized" {
		t.Errorf("expected database check 'uninitialized', got '%s'", resp.Checks["database"])
	}
}
