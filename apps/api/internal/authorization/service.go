package authorization

import (
	"context"

	"github.com/google/uuid"
)

// Role defines role definitions (e.g. Owner, Admin, Manager, Member, Client).
type Role struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	Permissions    []string  `json:"permissions"`
}

type Authorizer interface {
	HasPermission(ctx context.Context, orgID, userID uuid.UUID, permission string) (bool, error)
}
