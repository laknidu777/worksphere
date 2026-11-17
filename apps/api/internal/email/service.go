package email

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusSending   CampaignStatus = "sending"
	CampaignStatusSent      CampaignStatus = "sent"
)

type EmailCampaign struct {
	ID             uuid.UUID      `json:"id"`
	OrganizationID uuid.UUID      `json:"organization_id"`
	Title          string         `json:"title"`
	Subject        string         `json:"subject"`
	ContentHTML    string         `json:"content_html"`
	Status         CampaignStatus `json:"status"`
	SentCount      int            `json:"sent_count"`
	BounceCount    int            `json:"bounce_count"`
	ScheduledFor   *time.Time     `json:"scheduled_for,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
}

// EmailSender abstracts outbound delivery providers (Amazon SES candidate)
type EmailSender interface {
	SendTransactional(ctx context.Context, recipientEmail, subject, htmlBody string) error
	IsSuppressed(ctx context.Context, orgID uuid.UUID, recipientEmail string) (bool, error)
}
