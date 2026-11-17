package background

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Config represents background worker configuration
type Config struct {
	DatabaseURL string
	Concurrency int
}

// Worker orchestrates River transactional background jobs against PostgreSQL.
type Worker struct {
	pool   *pgxpool.Pool
	cancel context.CancelFunc
}

func NewWorker(pool *pgxpool.Pool) *Worker {
	return &Worker{
		pool: pool,
	}
}

// Start begins processing River background queues.
// Real River queue registration will be added as business modules (e.g. Email dispatch) are implemented.
func (w *Worker) Start(ctx context.Context) error {
	ctxWorker, cancel := context.WithCancel(ctx)
	w.cancel = cancel

	slog.Info("starting River background worker engine")

	// Placeholder loop verifying background loop health until River queues are wired
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctxWorker.Done():
				slog.Info("background worker engine shutting down")
				return
			case <-ticker.C:
				slog.Debug("river background worker heartbeat ok")
			}
		}
	}()

	return nil
}

func (w *Worker) Stop() {
	if w.cancel != nil {
		w.cancel()
	}
	slog.Info("background worker stopped cleanly")
}

func PlaceholderJobDescription() string {
	return fmt.Sprintf("River transactional job queue ready for email, billing reminders, and export workers.")
}
