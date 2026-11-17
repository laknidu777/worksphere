# ADR-006: Kubernetes (EKS), Helm, and Argo CD GitOps Deployment Architecture

## Status
**Decided (Target Architecture)** (2026-11-17)

## Context
WorkSphere requires a reproducible, multi-environment deployment strategy across local, staging, and production tiers. We must balance development velocity with production reliability, declarative infrastructure, zero-downtime rolling upgrades, and infrastructure-as-code best practices.

## Decision
We establish **Kubernetes on AWS EKS with Helm packaging and Argo CD GitOps**:
- **Packaging**: Applications are packaged into Docker OCI containers and defined declaratively via Helm charts located at `deployment/helm/worksphere`.
- **GitOps Continuous Delivery**: Argo CD synchronizes application state directly from the Git repository into the target EKS clusters.
- **Environment Progression**:
  - Local: Docker Compose (`docker-compose.yml`) for rapid developer inner-loop without Kubernetes overhead.
  - Development / Staging: Kustomize overlays in `deployment/kubernetes/overlays/dev` deployed to EKS cluster.
  - Production: Helm-templated releases managed via Argo CD with automated rollback and canary/blue-green support.
- **Resource Limits & Auto-scaling**: Horizontal Pod Autoscaler (HPA) configured based on CPU utilization and request concurrency.

## Consequences
- **Positive**: Strict declarative state, automated configuration drift detection, reproducible environments.
- **Negative / Risks**: Kubernetes operational complexity. Mitigated by keeping local development strictly Docker Compose-based and postponing full production EKS provisioning until staging validation.
