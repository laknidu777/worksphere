package identity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User represents an internal user record linked to an external OIDC identity provider (Cognito).
type User struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	ExternalSub    string    `json:"external_sub"` // Cognito Subject Identifier
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	Status         string    `json:"status"` // "active", "invited", "suspended"
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type IdentityProvider interface {
	VerifyToken(ctx context.Context, tokenString string) (*UserClaims, error)
}

type UserClaims struct {
	Subject       string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
