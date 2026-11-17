# WorkSphere Local Development Guide

## 1. Prerequisites

To run WorkSphere locally for development, ensure you have:
- **Node.js**: v20+ or v24+ (v24.x LTS recommended)
- **Go**: 1.22+ or 1.23+
- **Docker & Docker Compose**: v2.20+ / Docker Compose v5+
- **Git**

No AWS account, AWS credentials, or live external SaaS subscriptions are required to run WorkSphere locally.

---

## 2. Quickstart

### Step 1: Environment Configuration
Copy the root `.env.example` to `.env`:

```bash
cp .env.example .env
```

### Step 2: Start Local Infrastructure with Docker Compose
Run the local database (PostgreSQL 16) and Redis container stack:

```bash
docker compose up -d postgres redis
```

Verify containers are healthy:
```bash
docker compose ps
```

### Step 3: Run Database Migrations
Run Goose migrations against the local PostgreSQL container:

```bash
# Using make or direct goose
make db-migrate
# Or from apps/api:
# goose -dir db/migrations postgres "postgres://worksphere:worksphere_local_secret@localhost:5432/worksphere?sslmode=disable" up
```

### Step 4: Start the Backend API (Go)
```bash
cd apps/api
go run cmd/server/main.go
```
The API server starts on `http://localhost:8080`.
- Health check: `http://localhost:8080/health`
- Readiness check: `http://localhost:8080/ready`

### Step 5: Start the Frontend Application (Next.js)
In a separate terminal:
```bash
cd apps/web
npm install
npm run dev
```
The web dashboard opens on `http://localhost:3000`.

---

## 3. Running the Entire Stack via Docker Compose

To run the complete system (PostgreSQL, Redis, Go API, Next.js frontend) in containerized mode:

```bash
docker compose up --build
```

Access services:
- Frontend: `http://localhost:3000`
- Backend API: `http://localhost:8080`
- PostgreSQL: `localhost:5432`
- Redis: `localhost:6379`
