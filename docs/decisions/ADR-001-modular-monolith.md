# ADR-001: Modular Monolith Architecture

## Status
**Decided** (2026-11-17)

## Context
WorkSphere is an all-in-one business operations SaaS platform covering 9 functional domains (Organizations, Projects, Workforce, CRM, Finance, Analytics, Notifications, Subscriptions, Email Marketing).
An early architectural question is whether to build microservices or a monolith. Microservices introduce substantial operational overhead: network latency, distributed transactions, complex CI/CD pipelines, higher cloud hosting costs on AWS, and orchestration overhead.

## Decision
We adopt a **Modular Monolith** architecture:
- All domain modules reside in a single Go application codebase (`apps/api`).
- Each module has strict domain boundaries encapsulated inside `apps/api/internal/<module>`.
- Cross-module interaction happens via Go interfaces and in-memory function calls, or domain events via an in-process dispatcher / River background jobs.
- Direct database cross-table joins across distinct domain tables are minimized or prohibited; each domain owns its entities.

## Consequences
- **Positive**: Rapid feature development, simplified local development and debugging, atomic transactions where needed, zero distributed network latency between domains, lower cloud costs (single API container pool).
- **Negative / Risks**: Requires discipline to prevent tight coupling across packages. Code reviews must enforce module encapsulation boundaries.
