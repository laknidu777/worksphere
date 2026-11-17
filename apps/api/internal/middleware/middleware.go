package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const (
	RequestIDKey     contextKey = "request_id"
	TenantContextKey contextKey = "tenant_context"
)

type TenantContext struct {
	OrganizationID uuid.UUID
	UserID         uuid.UUID
	Role           string
	Permissions    []string
}

// RequestID attaches a unique request ID to each incoming HTTP request.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		w.Header().Set("X-Request-ID", reqID)
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// StructuredLogger logs HTTP request metadata using log/slog.
func StructuredLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		reqID, _ := r.Context().Value(RequestIDKey).(string)

		slog.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", ww.statusCode),
			slog.Duration("duration_ms", duration),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("request_id", reqID),
		)
	})
}

// Recoverer catches panics, logs stack traces, and returns a clean 500 error.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				reqID, _ := r.Context().Value(RequestIDKey).(string)
				slog.Error("panic recovered",
					slog.Any("error", rvr),
					slog.String("stack", string(debug.Stack())),
					slog.String("request_id", reqID),
				)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]any{
					"type":       "https://api.worksphere.lk/errors/internal",
					"title":      "Internal Server Error",
					"status":     http.StatusInternalServerError,
					"detail":     "An unexpected internal error occurred.",
					"request_id": reqID,
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TenantContextExtractor extracts tenant and user identifiers from request headers or auth context.
func TenantContextExtractor(mockAuthEnabled bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			orgHeader := r.Header.Get("X-Organization-ID")
			userHeader := r.Header.Get("X-User-ID")

			var orgID, userID uuid.UUID
			var err error

			if orgHeader != "" {
				orgID, err = uuid.Parse(orgHeader)
				if err != nil {
					writeProblem(w, http.StatusBadRequest, "Invalid Organization ID", "X-Organization-ID must be a valid UUID")
					return
				}
			}

			if userHeader != "" {
				userID, err = uuid.Parse(userHeader)
				if err != nil {
					writeProblem(w, http.StatusBadRequest, "Invalid User ID", "X-User-ID must be a valid UUID")
					return
				}
			}

			// In local development / mock auth mode, provide fallback identifiers if absent
			if mockAuthEnabled && orgID == uuid.Nil {
				orgID = uuid.MustParse("01933a5e-38c2-7472-8f92-ec64152e259b")
				userID = uuid.MustParse("01933a5e-38c2-7472-8f92-ec64152e259c")
			}

			tc := &TenantContext{
				OrganizationID: orgID,
				UserID:         userID,
				Role:           "admin",
				Permissions:    []string{"*"},
			}

			ctx := context.WithValue(r.Context(), TenantContextKey, tc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTenantContext(ctx context.Context) (*TenantContext, error) {
	tc, ok := ctx.Value(TenantContextKey).(*TenantContext)
	if !ok || tc == nil || tc.OrganizationID == uuid.Nil {
		return nil, fmt.Errorf("tenant context missing from request")
	}
	return tc, nil
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func writeProblem(w http.ResponseWriter, status int, title, detail string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"title":  title,
		"status": status,
		"detail": detail,
	})
}
