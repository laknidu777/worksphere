package organizations

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Organization represents a multi-tenant business workspace.
type Organization struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Currency  string     `json:"currency"` // Default "LKR"
	Timezone  string     `json:"timezone"` // Default "Asia/Colombo"
	Locale    string     `json:"locale"`   // Default "en-LK"
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateOrganizationRequest struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Currency string `json:"currency"`
	Timezone string `json:"timezone"`
}

type Service interface {
	CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*Organization, error)
	GetOrganization(ctx context.Context, id uuid.UUID) (*Organization, error)
}
