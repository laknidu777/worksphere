# WorkSphere Database Design & Standards

## 1. Core Principles

WorkSphere uses PostgreSQL 16+ as its single source of truth for business data, transactional events, and background jobs.

1. **UUIDv7 Primary Keys**:
   All entities use time-sortable UUIDv7 identifiers (`uuid` type in PostgreSQL) to provide uniform distributed key generation without auto-increment hotspot bottlenecks.
2. **Strict Financial Types**:
   All currency, price, rate, and fee values MUST be stored as `NUMERIC(14, 2)` (or higher precision where unit fractional calculations require). **Floating-point types (`REAL`, `DOUBLE PRECISION`, `FLOAT`) are strictly prohibited** for money.
3. **UTC Timestamps Everywhere**:
   All datetime columns must be declared as `TIMESTAMPTZ` (timestamp with time zone) and stored strictly in UTC. Local conversions to Sri Lanka Standard Time (`Asia/Colombo`) occur strictly in the presentation or client layer.
4. **Tenant Isolation Columns**:
   Every organization-scoped table includes an `organization_id UUID NOT NULL` column with appropriate indexes and composite foreign keys.
5. **Soft Deletion Policy**:
   Operational tables requiring retention or audit compliance use `deleted_at TIMESTAMPTZ NULL`. Hard deletes are reserved for GDPR/privacy compliance or explicit tenant purging.

---

## 2. Migration Tooling: Goose

Database migrations are managed using **Goose** under `apps/api/db/migrations`.
- File naming format: `YYYYMMDDHHMMSS_description.sql`
- Each migration file MUST define both `-- +goose Up` and `-- +goose Down` blocks.
- Migrations must be idempotent and non-blocking in production environments (e.g. `CREATE INDEX CONCURRENTLY`).

---

## 3. Query Generation: sqlc

SQL queries are written in raw SQL under `apps/api/db/queries` and compiled into type-safe, idiomatic Go code using **sqlc**.
- ORMs (such as GORM) are explicitly avoided to guarantee optimal query performance, deterministic indexing, and transparent query auditing.
- `sqlc.yaml` is configured to target PostgreSQL, generate structs with JSON tags, and output to `apps/api/internal/database`.

---

## 4. Sri Lankan Business Schema Standards

- **Primary Currency**: Sri Lankan Rupee (`LKR`).
- **Standard Tax Identifiers**:
  - `tin` (Taxpayer Identification Number)
  - `vat_number` (Value Added Tax registration number)
  - `sscl_registered` (Social Security Contribution Levy indicator)
- **Formatting Conventions**:
  - Currency display: `LKR 125,000.00`
  - National Identity Card (NIC): Validated against both old (9 digits + V/X) and new (12 digits) formats.
  - Phone numbers: Validated against E.164 convention with +94 default country code.

---

## 5. Background Jobs: River Schema

WorkSphere uses **River** for robust, transactional background task queues directly inside PostgreSQL.
- River shares the same database transaction as the business operation (`tx`), preventing dual-write failures (e.g., an invoice is created, and the job to send the email notification is enqueued in the exact same transaction).
- River migration tables are initialized via River CLI / migrations alongside application migrations.
