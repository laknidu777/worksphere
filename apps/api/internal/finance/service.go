package finance

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InvoiceStatus string

const (
	InvoiceStatusDraft     InvoiceStatus = "draft"
	InvoiceStatusIssued    InvoiceStatus = "issued"
	InvoiceStatusPaid      InvoiceStatus = "paid"
	InvoiceStatusOverdue   InvoiceStatus = "overdue"
	InvoiceStatusCancelled InvoiceStatus = "cancelled"
)

type Invoice struct {
	ID             uuid.UUID     `json:"id"`
	OrganizationID uuid.UUID     `json:"organization_id"`
	ClientID       uuid.UUID     `json:"client_id"`
	InvoiceNumber  string        `json:"invoice_number"`
	Currency       string        `json:"currency"`        // "LKR"
	SubtotalAmount string        `json:"subtotal_amount"` // PostgreSQL NUMERIC string
	TaxAmount      string        `json:"tax_amount"`      // PostgreSQL NUMERIC string
	TotalAmount    string        `json:"total_amount"`    // PostgreSQL NUMERIC string
	Status         InvoiceStatus `json:"status"`
	IssueDate      time.Time     `json:"issue_date"`
	DueDate        time.Time     `json:"due_date"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type Service interface {
	GetInvoice(ctx context.Context, orgID, invoiceID uuid.UUID) (*Invoice, error)
}
