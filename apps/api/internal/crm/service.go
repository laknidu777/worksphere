package crm

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ClientCompany struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	ContactPerson  string    `json:"contact_person"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"` // +94 convention
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Service interface {
	GetClient(ctx context.Context, orgID, clientID uuid.UUID) (*ClientCompany, error)
}
