package billing

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SubscriptionPlan string

const (
	PlanStarter    SubscriptionPlan = "starter"
	PlanGrowth     SubscriptionPlan = "growth"
	PlanEnterprise SubscriptionPlan = "enterprise"
)

type Subscription struct {
	ID               uuid.UUID        `json:"id"`
	OrganizationID   uuid.UUID        `json:"organization_id"`
	Plan             SubscriptionPlan `json:"plan"`
	Status           string           `json:"status"`        // "trialing", "active", "past_due", "canceled"
	Currency         string           `json:"currency"`      // "LKR"
	PriceMonthly     string           `json:"price_monthly"` // NUMERIC string
	TrialEndsAt      *time.Time       `json:"trial_ends_at,omitempty"`
	CurrentPeriodEnd time.Time        `json:"current_period_end"`
}

// PaymentGateway abstracts external payment providers (PayHere candidate, Stripe, etc.)
type PaymentGateway interface {
	InitiateCheckout(ctx context.Context, orgID uuid.UUID, amountLKR string, returnURL string) (*CheckoutSession, error)
	VerifyWebhookSignature(payload []byte, signature string) bool
}

type CheckoutSession struct {
	SessionID   string `json:"session_id"`
	RedirectURL string `json:"redirect_url"`
}
