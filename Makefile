.PHONY: help install test test-backend test-frontend lint build docker-up docker-down

help:
	@echo "WorkSphere Development Commands:"
	@echo "  make install        Install frontend and backend dependencies"
	@echo "  make test           Run all tests (frontend and backend)"
	@echo "  make test-backend   Run Go backend tests"
	@echo "  make test-frontend  Run Next.js frontend vitest"
	@echo "  make lint           Run Go format check and Next.js ESLint"
	@echo "  make build          Build both backend binary and frontend Next.js app"
	@echo "  make docker-up      Start PostgreSQL and Redis local services"
	@echo "  make docker-down    Stop all Docker containers"
	@echo "  make dev-backend    Start local Go backend server"
	@echo "  make dev-frontend   Start local Next.js frontend server"

install:
	cd apps/api && go mod download
	cd apps/web && npm install

test: test-backend test-frontend

test-backend:
	cd apps/api && go test -v ./...

test-frontend:
	cd apps/web && npm run test

lint:
	cd apps/api && go fmt ./...
	cd apps/web && npm run lint

build: build-backend build-frontend

build-backend:
	cd apps/api && go build -v -o bin/server cmd/server/main.go

build-frontend:
	cd apps/web && npm run build

docker-up:
	docker compose up -d postgres redis

docker-down:
	docker compose down

docker-all:
	docker compose up --build

dev-backend:
	cd apps/api && go run cmd/server/main.go

dev-frontend:
	cd apps/web && npm run dev
