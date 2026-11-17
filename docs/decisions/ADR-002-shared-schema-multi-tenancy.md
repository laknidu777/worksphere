# ADR-002: Shared-Schema Multi-Tenancy with Row-Level Security

## Status
**Decided** (2026-11-17)

## Context
A multi-tenant SaaS requires rigorous tenant data isolation. The primary models are:
1. Database-per-tenant (costly, hard to scale across hundreds of small businesses).
2. Schema-per-tenant (complex schema migrations, RDS connection limits).
3. Shared-database, shared-schema with tenant discriminator column (`organization_id`).

## Decision
We adopt **Shared-Database, Shared-Schema Multi-Tenancy**:
- Every organization-owned table has `organization_id UUID NOT NULL`.
- Foreign keys where possible include `organization_id` to enforce intra-tenant referential integrity.
- PostgreSQL Row-Level Security (RLS) is configured on all tenant tables as a defense-in-depth safety net.
- Transactions set `SET LOCAL app.current_tenant_id = '<tenant-uuid>'` upon checkout from the pgx connection pool.
- Application-level authorization checks are always the first line of defense; RLS is the database-level safety net.

## Consequences
- **Positive**: Extremely cost effective on AWS RDS, simple connection pooling via PgBouncer / RDS Proxy, single schema migration run per deployment.
- **Negative / Risks**: A bug in query construction or RLS bypass could risk data leakage; mitigated by strict `sqlc` query generation, RLS policies, and automated cross-tenant security integration tests.
