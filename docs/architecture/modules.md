# WorkSphere Product Modules & Domain Boundaries

WorkSphere is composed of nine primary business modules. Although deployed within a single Go modular monolith (`apps/api`), each module is treated as an isolated domain package with explicit boundary rules:

- **Independent Domain State**: Modules directly manage only their own tables and state.
- **Explicit Interfaces**: Communication across module boundaries must go through exported Go interfaces or asynchronous domain events.
- **Tenant Context First**: All business operations require an authenticated tenant context (`TenantContext` containing `OrganizationID` and `UserID`).

---

## Module Index

### 1. Organization & User Management (`internal/organizations`, `internal/identity`)
Responsible for multi-tenant workspace setup and workforce directory.
- **Key Capabilities**:
  - Company registration and multi-tenant workspace provisioning.
  - User invitations via secure single-use tokens.
  - Department and team hierarchal assignments.
  - Employee profiles and organizational metadata.
  - Organization settings (branding, regional business hours, working day calendars).

### 2. Project & Task Management (`internal/projects`, `internal/tasks`)
Enables project delivery and cross-functional execution.
- **Key Capabilities**:
  - Project lifecycle tracking (planning, active, paused, completed).
  - Task assignment with hierarchical subtasks and checklists.
  - Multiple view layouts: Kanban boards, list views, and timeline overviews.
  - Priority levels, due dates, estimates, and progress tracking.
  - Task discussions, activity history, and file attachments.
  - Milestone definitions and task dependency relationships.

### 3. Workforce & Time Management (`internal/workforce`, `internal/leave`)
Optimized for regional Sri Lankan workforce operating standards and flexibility.
- **Key Capabilities**:
  - Digital clock-in / clock-out logging with IP / geolocation audit options.
  - Timesheet creation, submission, and managerial approval workflows.
  - Working hour analytics, overtime computation, and attendance exception tracking.
  - Leave entitlements (annual, casual, medical, mercantile holidays).
  - Multi-tier leave approval chains.
  - Team workload distribution and capacity visibility.

### 4. Client & CRM Management (`internal/crm`)
Drives lead generation, customer relationships, and account management.
- **Key Capabilities**:
  - Customer profiles, multi-contact directories, and company hierarchies.
  - Sales pipeline tracking with customizable deal stages.
  - Interaction history (meetings, phone calls, outbound email threads).
  - Direct linking between CRM accounts, ongoing projects, and billing contracts.
  - Restricted client portal access boundaries (planned external stakeholder view).

### 5. Finance & Expense Management (`internal/finance`)
Multi-tenant financial operations configured for Sri Lankan business defaults.
- **Key Capabilities**:
  - Invoice generation, line-item itemization, and tax calculation (VAT/SSCL compliant readiness).
  - Multi-currency support with native LKR primary denomination (PostgreSQL `NUMERIC`).
  - Expense logging with receipts, categorization, and project-level cost tracking.
  - Payment recording, outstanding balance tracking, and partial payment support.
  - Project budget vs. actual expenditure monitoring and real-time profitability margins.

### 6. Business Analytics (`internal/analytics`)
Actionable operational and financial intelligence for business leaders.
- **Key Capabilities**:
  - Executive financial summary (billed revenue, collections, cash flow forecast).
  - Project delivery metrics (velocity, deadline compliance, bottleneck identification).
  - Employee utilization rate and billable hours analytics.
  - Aging accounts receivable and outstanding invoice analytics.
  - Real-time business KPI dashboards.

### 7. Notifications & Automation (`internal/notifications`)
Unified transactional notification hub and background automation scheduler.
- **Key Capabilities**:
  - Multi-channel delivery engine (in-app notifications, transactional email).
  - Automated task reminders and upcoming deadline warnings.
  - Overdue invoice follow-ups and payment receipt acknowledgements.
  - Leave approval request routing and status updates.
  - Recurring task generators and template-driven business workflows.
  - Comprehensive activity and audit trail.

### 8. SaaS Subscription Management (`internal/billing`)
The engine governing WorkSphere customer billing and multi-tenant plan enforcement.
- **Key Capabilities**:
  - Free trial provisioning, duration management, and expiration enforcement.
  - Tiered subscription plans (monthly and annual billing frequencies).
  - Modular payment gateway abstraction layer (PayHere candidate integration).
  - Usage limits enforcement (active user caps, storage limits, project limits).
  - Self-service upgrades, downgrades, and billing statement history.
  - Platform-wide super-administrator administration panel for platform operators.

### 9. Email Marketing & Client Communication (`internal/email`)
Outbound marketing campaigns and segmented email delivery.
- **Key Capabilities**:
  - Contact and subscriber management with CSV import validation.
  - Audience lists and dynamic segment filtering based on CRM properties.
  - Reusable responsive email template builder and previewer.
  - Campaign authoring, scheduling, and rate-limited batch dispatching via River.
  - Delivery telemetry: open tracking, click tracking, bounce, and complaint suppression.
  - One-to-one outbound transactional and sales email logging.
  - Mandatory unsubscribe link handling and consent management.
  - *Note*: Outbound-only. No inbound shared inbox is implemented in this phase.
