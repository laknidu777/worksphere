package tasks

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
	PriorityUrgent TaskPriority = "urgent"
)

type Task struct {
	ID             uuid.UUID    `json:"id"`
	OrganizationID uuid.UUID    `json:"organization_id"`
	ProjectID      uuid.UUID    `json:"project_id"`
	AssigneeID     *uuid.UUID   `json:"assignee_id,omitempty"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Priority       TaskPriority `json:"priority"`
	DueDate        *time.Time   `json:"due_date,omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type Service interface {
	GetTask(ctx context.Context, orgID, taskID uuid.UUID) (*Task, error)
}
