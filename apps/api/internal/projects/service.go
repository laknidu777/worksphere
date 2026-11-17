package projects

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ProjectStatus string

const (
	ProjectStatusPlanning  ProjectStatus = "planning"
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusOnHold    ProjectStatus = "on_hold"
	ProjectStatusCompleted ProjectStatus = "completed"
)

type Project struct {
	ID             uuid.UUID     `json:"id"`
	OrganizationID uuid.UUID     `json:"organization_id"`
	ClientID       *uuid.UUID    `json:"client_id,omitempty"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Status         ProjectStatus `json:"status"`
	BudgetLKR      *string       `json:"budget_lkr,omitempty"` // Stored as NUMERIC string
	StartDate      *time.Time    `json:"start_date,omitempty"`
	TargetEndDate  *time.Time    `json:"target_end_date,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type Service interface {
	GetProject(ctx context.Context, orgID, projectID uuid.UUID) (*Project, error)
}
