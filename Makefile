.PHONY: dev build templ migrate-up migrate-down seed clean lint

# ── Development ──────────────────────────────────────────────────────────────

## dev: run with hot-reload via air
dev:
	@air -c .air.toml

## run: run without hot-reload (useful in CI)
run:
	@go run cmd/main.go

## build: compile a production binary
build:
	@go build -o bin/app cmd/main.go

# ── Templates ─────────────────────────────────────────────────────────────────

## templ: generate Go code from .templ files
templ:
	@templ generate

# ── Database ──────────────────────────────────────────────────────────────────

## migrate-up: apply all pending SQL migrations via goose CLI (optional, app also auto-migrates on startup)
migrate-up:
	@goose -dir db/migrations sqlite3 $${DB_NAME:-app.db} up

## migrate-down: roll back the most recent migration via goose CLI
migrate-down:
	@goose -dir db/migrations sqlite3 $${DB_NAME:-app.db} down

## migrate-status: show migration status via goose CLI
migrate-status:
	@goose -dir db/migrations sqlite3 $${DB_NAME:-app.db} status

## seed: run the database seeder
seed:
	@go run cmd/seed/main.go

# ── Code Quality ──────────────────────────────────────────────────────────────

## lint: run golangci-lint
lint:
	@golangci-lint run ./...

## test: run all tests
test:
	@go test ./... -v

## clean: remove build artifacts and database file
clean:
	@rm -rf bin/ app.db

# ── Help ──────────────────────────────────────────────────────────────────────

## help: display this message
help:
	@echo "Available targets:"
	@grep -E '^## ' Makefile | sed 's/## /  make /'
