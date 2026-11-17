# WorkSphere

> **Sri Lanka-First, Subscription-Based, Multi-Tenant Business Operations SaaS Platform**

WorkSphere is a unified operational backbone engineered for growing businesses in Sri Lanka and regional emerging markets. It eliminates fragmented spreadsheets and disjointed point solutions by integrating organization management, workforce timekeeping, client relationships, project delivery, billing, and financial operations into a secure, high-performance platform.

---

## Architecture Blueprint

WorkSphere is designed as a **Modular Monolith** prioritizing operational simplicity, strict domain isolation, and rapid product iteration:

```
worksphere/
├── apps/
│   ├── web/                    # Next.js 15+ App Router, Tailwind CSS, TanStack Query, TypeScript
│   └── api/                    # Go 1.23+ with Chi HTTP router, pgx/v5, and sqlc
│       ├── cmd/server/         # Main backend application entrypoint
│       ├── internal/           # Encapsulated domain packages (strict module boundaries)
│       │   ├── identity/       # OIDC authentication & Cognito subject mapping
│       │   ├── organizations/  # Multi-tenant workspace management
│       │   ├── authorization/  # RBAC & permission evaluation engine
│       │   ├── projects/       # Project delivery lifecycles
│       │   ├── tasks/          # Tasks, Kanban views, dependencies
│       │   ├── workforce/      # Digital clock-in/out & timesheets
│       │   ├── leave/          # Leave requests & approval workflows
│       │   ├── crm/            # Client directories & sales pipelines
│       │   ├── finance/        # LKR invoicing, expenses, profitability
│       │   ├── billing/        # SaaS subscriptions & PayHere IPG abstraction
│       │   ├── notifications/  # Transactional alerts & in-app feeds
│       │   ├── email/          # Outbound marketing & SES campaign delivery
│       │   ├── audit/          # Compliance & operational audit logging
│       │   ├── config/         # Environment & runtime configuration
│       │   ├── database/       # pgx connection pool & transaction utilities
│       │   ├── health/         # Shallow liveness and deep readiness probes
│       │   ├── middleware/     # Request ID, structured slog, recovery, tenant context
│       │   └── server/         # Chi router assembly and HTTP server lifecycle
│       ├── db/
│       │   ├── migrations/     # Goose SQL schema migrations
│       │   └── queries/        # sqlc raw SQL query definitions
│       └── tests/              # Backend integration and security tests
├── workers/
│   └── background/             # River transactional background job runners
├── packages/
│   └── api-contracts/          # OpenAPI 3.1 specifications & contract schemas
├── infrastructure/             # Infrastructure as Code
│   ├── terraform/              # AWS VPC, RDS, and EKS modules
│   ├── ansible/                # Auxiliary host provisioning
│   └── aws/                    # IAM policies and IRSA definitions
├── deployment/                 # Deployment configurations
│   ├── docker/                 # Multi-stage Dockerfiles for API and Web
│   ├── compose/                # Local container definitions
│   ├── kubernetes/             # Declarative K8s manifests & Kustomize overlays
│   ├── helm/                   # Parameterized Helm charts
│   └── argocd/                 # GitOps application manifests
├── monitoring/                 # Observability stack
│   ├── prometheus/             # Metrics collection & scrape configs
│   ├── grafana/                # Operational dashboards
│   ├── loki/                   # Centralized log collection
│   └── alerts/                 # Alertmanager notification rules
├── docs/                       # Architecture & engineering documentation
│   ├── architecture/           # System overview, modules, multi-tenancy model
│   ├── database/               # Database standards, NUMERIC conventions, RLS
│   ├── api/                    # REST conventions, RFC 7807 error formats
│   ├── security/               # Security baseline & secret hygiene
│   ├── deployment/             # Local development & cloud deployment roadmap
│   └── decisions/              # Architecture Decision Records (ADRs)
├── tools/
│   └── github-bot/             # Preserved legacy GitHub + OpenAI starter integration
├── docker-compose.yml          # Local multi-container development environment
├── Makefile                    # Developer workflow automation commands
├── .editorconfig               # Uniform formatting rules
├── .gitignore                  # Strict repository hygiene and secret exclusions
├── .env.example                # Safe environment variable configuration template
└── README.md
```

---

## 9 Core Product Modules

1. **Organization & User Management**: Company registration, workspace provisioning, employee invitations, departments, and organization settings.
2. **Project & Task Management**: Kanban boards, task lists, milestones, dependencies, priorities, and deadlines.
3. **Workforce & Time Management**: Digital clock-in/out, timesheets, working hour reports, and leave approvals.
4. **Client & CRM Management**: Client directories, customer contacts, interaction histories, and sales pipelines.
5. **Finance & Expense Management**: Invoices, expense claims, payment tracking, project budgets, and profitability metrics.
6. **Business Analytics**: Executive revenue summaries, employee utilization, overdue receivables, and business KPIs.
7. **Notifications & Automation**: Transactional email/in-app notifications, task reminders, and automated workflows.
8. **SaaS Subscription Management**: Multi-tier plans, free trials, usage limit enforcement, and payment provider abstraction.
9. **Email Marketing & Client Communication**: Outbound campaigns, contact lists, templates, unsubscribe suppression, and delivery telemetry (*outbound-only; no inbound mailbox*).

---

## Sri Lankan Business Defaults

WorkSphere ships with pre-configured Sri Lankan business defaults:
- **Currency**: `LKR` (Sri Lankan Rupee) — financial values stored strictly as PostgreSQL `NUMERIC(14, 2)`.
- **Timezone**: `Asia/Colombo` (UTC+05:30) — timestamps stored in UTC and converted at presentation layer.
- **Locale**: `en-LK` (formatting dates as `DD/MM/YYYY`).
- **Telephone Convention**: `+94` E.164 parsing.
- **Candidate Integrations**: PayHere (payment gateway abstraction) and Amazon SES (outbound email). Both are proposed candidates pending validation.

---

## Technology Stack

- **Frontend**: Next.js 15+ (App Router), TypeScript, Tailwind CSS, TanStack Query, Zod, Vitest.
- **Backend**: Go 1.23+, Chi router, `pgx/v5`, `sqlc`, `goose`, `slog` structured logging.
- **Data & Queue**: PostgreSQL 16+ (Row-Level Security), Redis 7+, River transactional queue.
- **DevOps & Cloud**: Docker, Docker Compose, Kubernetes, Helm, Argo CD, Terraform, AWS.

---

## Local Development Quickstart

### Prerequisites
- Node.js 20+ or 24+
- Go 1.23+
- Docker & Docker Compose

### 1. Configure Environment
```bash
cp .env.example .env
```

### 2. Start Core Database & Redis
```bash
docker compose up -d postgres redis
```

### 3. Run Backend (Go)
```bash
cd apps/api
go test ./...       # Run tests
go run cmd/server/main.go
```
- API Base: `http://localhost:8080`
- Liveness Probe: `http://localhost:8080/health`
- Readiness Probe: `http://localhost:8080/ready` (checks DB & Redis; returns 503 if unavailable)

### 4. Run Frontend (Next.js)
```bash
cd apps/web
npm install
npm run test        # Run vitest
npm run dev         # Start Next.js dev server
```
- Web Application: `http://localhost:3000`

---

## Automated Verification & CI

Run repository checks via `make`:

```bash
make lint           # Go formatting + Next.js ESLint
make test           # Go unit/integration tests + Frontend vitest
make build          # Go binary build + Next.js production build
```

---

## Preserved Legacy Work

The earlier GitHub webhook + OpenAI starter integration was preserved and relocated to:
- [`tools/github-bot/`](file:///c:/Users/admin/Documents/GitHub/Plymouth%20Projects/worksphere/tools/github-bot)
