# ADR-003: Go and Chi for Core Backend API

## Status
**Decided** (2026-11-17)

## Context
WorkSphere requires a backend capable of high throughput, low memory footprint, strict type safety, fast startup times, and minimal container image sizes. Node.js, Python, and Java were considered alongside Go.

## Decision
We select **Go (Golang) 1.23+ with Chi HTTP router** for `apps/api`:
- **Language**: Go offers compiled native performance, excellent concurrency primitives (goroutines/channels), low memory usage (~25MB container resident set size), and fast build times.
- **HTTP Router**: `go-chi/chi/v5` is lightweight, 100% compatible with Go's `net/http` standard library, provides zero external dependencies, and supports composable middleware (request ID, structured logging, CORS, recovery, rate limiting).
- **Data Access**: `pgx/v5` connection pool with `sqlc` for compile-time verified type-safe SQL queries.

## Consequences
- **Positive**: Extremely fast response times, small Docker images (<30MB based on alpine or distroless), predictable CPU/memory consumption on AWS, and no heavy framework magic.
- **Negative / Risks**: More boilerplate compared to full-stack frameworks (e.g. NestJS or Django); addressed by modular package layout and `sqlc` generation.
