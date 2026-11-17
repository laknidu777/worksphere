# ADR-005: Payment Provider Abstraction Layer (PayHere Candidate)

## Status
**Proposed / Validation Required** (2026-11-17)

## Context
WorkSphere is a Sri Lanka-first SaaS platform targeting local businesses that primarily transact in Sri Lankan Rupees (LKR) via local payment mechanisms (LankaPay, eZ Cash, mCash, local debit/credit cards), but may also expand to international billing (Stripe).
PayHere is widely used in Sri Lanka as an internet payment gateway aggregator. However, hardcoding PayHere into domain models risks vendor lock-in and creates rigidity if another provider (e.g. Commercial Bank IPG, Sampath Bank IPG, Stripe) is preferred.

## Decision
We establish a **Payment Provider Abstraction Layer** within `apps/api/internal/billing`:
- A clean domain interface `PaymentGateway` defines operations:
  - `InitiateCheckout(ctx, req) (*CheckoutSession, error)`
  - `ProcessWebhook(ctx, payload, signature) (*WebhookEvent, error)`
  - `VerifyPayment(ctx, transactionID) (*PaymentStatus, error)`
  - `ManageSubscription(ctx, subscriptionID) error`
- PayHere is recognized as the primary candidate implementation for LKR processing (`payhere.Provider`), subject to sandbox integration testing and legal agreement validation.
- Mock payment gateway implementation (`mock.Provider`) is provided for local automated testing and development without live credentials.
- All monetary transactions are recorded using PostgreSQL `NUMERIC` types with explicit ISO 4217 currency codes (`LKR` by default).

## Consequences
- **Positive**: Complete isolation of third-party payment mechanics from core subscription billing workflows; ability to switch or run multiple IPGs simultaneously.
- **Validation Needed**: PayHere recurring subscription tokenization API stability, merchant account KYC requirements, and refund webhook handling.
