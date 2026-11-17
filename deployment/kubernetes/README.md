# Kubernetes Declarative Manifests

Contains baseline Kubernetes manifests and Kustomize overlays:
- `base`: Core deployments, services, horizontal pod autoscalers, and configmaps.
- `overlays/dev`: Development namespace overrides.
- `overlays/prod`: Multi-AZ production topology with strict pod disruption budgets.
