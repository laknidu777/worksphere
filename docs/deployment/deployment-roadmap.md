# WorkSphere Deployment & Cloud Roadmap

## 1. Deployment Philosophy

WorkSphere follows a staged deployment roadmap designed for high availability, reproducibility, zero-downtime rollouts, and cost control:

- **Local & CI**: Docker Compose, local tools, mock providers.
- **Stage 1 (Single Node / Staging Preview)**: AWS ECS or EKS with Terraform-managed state, automated PR previews.
- **Stage 2 (Production Kubernetes)**: Managed AWS EKS, RDS PostgreSQL Multi-AZ, ElastiCache Redis, Amazon ECR, Argo CD GitOps.

---

## 2. Infrastructure as Code (Terraform)

Located under `infrastructure/terraform/`:
- `modules/vpc`: Dedicated VPC with public, private, and database isolated subnets across 3 Availability Zones.
- `modules/rds`: AWS Aurora PostgreSQL or RDS PostgreSQL Multi-AZ with automated snapshots, encryption at rest via AWS KMS, and connection pooling.
- `modules/eks`: Managed AWS EKS cluster with managed node groups, IAM Roles for Service Accounts (IRSA), and AWS Load Balancer Controller.
- `environments/{dev,staging,prod}`: Environment root modules parameterized for clean state separation.

---

## 3. Configuration Management (Ansible)

Located under `infrastructure/ansible/`:
- Playbooks for auxiliary host provisioning, security baseline hardening, and log forwarder orchestration where non-containerized resources are required.

---

## 4. GitOps Delivery (Argo CD & Helm)

Located under `deployment/`:
- `deployment/helm/worksphere`: Unified Helm chart parameterizing frontend and backend deployments, services, ingress annotations, resource limits, and Horizontal Pod Autoscalers (HPA).
- `deployment/kubernetes/`: Raw declarative manifests with Kustomize overlays for dev, staging, and prod.
- `deployment/argocd/`: Argo CD Application manifests synchronizing Kubernetes state from Git repository revisions.

---

## 5. Production Observability Stack

Located under `monitoring/`:
- **Prometheus**: Scrapes `/metrics` endpoints from the Go API and node metrics.
- **Grafana**: Pre-configured dashboards for HTTP throughput, p95/p99 latency, error rates, database connection pool statistics, and River queue depths.
- **Loki**: Centralized log aggregation with Promtail / Fluentbit log shipping.
- **Alertmanager**: Pager and webhook notification routing for SLO breaches.
