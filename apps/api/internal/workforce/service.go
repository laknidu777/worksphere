package workforce

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AttendanceRecord struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	UserID         uuid.UUID  `json:"user_id"`
	ClockInAt      time.Time  `json:"clock_in_at"`  // UTC
	ClockOutAt     *time.Time `json:"clock_out_at"` // UTC
	Notes          string     `json:"notes,omitempty"`
}

type Service interface {
	RecordClockIn(ctx context.Context, orgID, userID uuid.UUID) (*AttendanceRecord, error)
	RecordClockOut(ctx context.Context, orgID, userID uuid.UUID) (*AttendanceRecord, error)
}
