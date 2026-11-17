package audit

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	UserID         *uuid.UUID `json:"user_id,omitempty"`
	Action         string     `json:"action"`
	Entity         string     `json:"entity"`
	EntityID       uuid.UUID  `json:"entity_id"`
	Metadata       string     `json:"metadata,omitempty"` // JSON string
	IPAddress      string     `json:"ip_address,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type Logger interface {
	Log(ctx context.Context, log AuditLog) error
}
