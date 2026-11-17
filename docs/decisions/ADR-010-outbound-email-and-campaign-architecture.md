# ADR-010: Outbound Email Marketing and Client Communication Architecture

## Status
**Decided** (2026-11-17)

## Context
WorkSphere requires client communication and email marketing capabilities to enable tenant businesses to engage leads, issue announcements, and send transactional notifications.
Email systems can easily become overly complex if shared multi-tenant inbound mailboxes, IMAP sync, and spam filtering are introduced prematurely. Additionally, sender reputation and deliverability are paramount to avoid domain blacklisting.

## Decision
1. **Outbound-Only Initial Scope**:
   The email module is strictly **outbound-only**. WorkSphere does NOT implement an inbound shared mailbox, IMAP server, or incoming mail parsing engine in this phase.
2. **Transactional & Campaign Separation**:
   - Transactional emails (invoices, password resets, system alerts) receive dedicated high-priority River job queues and high-reputation sender pools.
   - Bulk campaign emails (marketing blasts, newsletters) are processed via throttled, rate-limited River workers to protect IP/sender reputation.
3. **Candidate Delivery Provider: Amazon SES**:
   Amazon Simple Email Service (SES) is selected as the primary candidate provider due to AWS ecosystem alignment, cost-efficiency, and bounce/complaint SNS feedback loops. It remains a proposed provider subject to production domain DKIM/SPF verification.
4. **Mandatory Unsubscribe and Suppression Enforcement**:
   - Every marketing email MUST include a standardized `List-Unsubscribe` header and one-click unsubscribe footer link.
   - The database maintains a tenant-scoped `email_suppressions` table. Any contact that unsubscribes, hard-bounces, or logs a spam complaint is automatically suppressed from future campaign dispatches.

## Consequences
- **Positive**: Focused scope, high reliability, low infrastructure cost, robust anti-spam and deliverability protection.
- **Negative / Risks**: Tenants needing shared team inbox / inbound email ticketing must integrate external ticketing tools until WorkSphere introduces inbound email capabilities in a later roadmap phase.
