# WorkSphere Terraform Infrastructure

This directory manages cloud infrastructure as code (IaC) for WorkSphere on AWS.

## Structure
- `modules/vpc`: VPC networking, public subnets, private subnets, NAT Gateways.
- `modules/rds`: Amazon RDS PostgreSQL Multi-AZ with automated backups and KMS encryption.
- `modules/eks`: Amazon EKS managed Kubernetes cluster configuration and node groups.
- `environments/dev`: Development environment terraform root.
- `environments/staging`: Staging environment terraform root.
- `environments/prod`: Production environment terraform root.

> [!NOTE]
> Per architectural guidelines, paid AWS resources are not provisioned during initial repository foundation.
