package leave

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LeaveType string

const (
	LeaveTypeAnnual     LeaveType = "annual"
	LeaveTypeCasual     LeaveType = "casual"
	LeaveTypeMedical    LeaveType = "medical"
	LeaveTypeMercantile LeaveType = "mercantile"
)

type LeaveRequest struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	UserID         uuid.UUID  `json:"user_id"`
	LeaveType      LeaveType  `json:"leave_type"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        time.Time  `json:"end_date"`
	DaysCount      float32    `json:"days_count"`
	Status         string     `json:"status"` // "pending", "approved", "rejected"
	ApprovedBy     *uuid.UUID `json:"approved_by,omitempty"`
}

type Service interface {
	SubmitLeaveRequest(ctx context.Context, req LeaveRequest) (*LeaveRequest, error)
}
