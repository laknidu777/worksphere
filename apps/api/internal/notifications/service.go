package notifications

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Read           bool      `json:"read"`
	CreatedAt      time.Time `json:"created_at"`
}

type Notifier interface {
	SendInApp(ctx context.Context, orgID, userID uuid.UUID, title, body string) (*Notification, error)
}
