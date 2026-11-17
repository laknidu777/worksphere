# Observability & Monitoring

WorkSphere telemetry architecture:
- `prometheus`: Scrape definitions and alerting rules.
- `grafana`: Pre-built operational dashboards for Chi HTTP metrics, pgx pool metrics, and River job queues.
- `loki`: Distributed log indexing configurations.
- `alerts`: Alertmanager routing policies for P1/P2 operational incident escalation.
