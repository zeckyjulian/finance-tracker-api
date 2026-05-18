.PHONY: help up down dev logs build run migrate seed test lint

# Tampilkan semua perintah
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# ── Docker ────────────────────────────────────────────────────
up: ## Jalankan semua service (API + Postgres + Redis)
	docker compose up -d

dev: ## Jalankan dengan pgAdmin (untuk development)
	docker compose --profile dev up -d

down: ## Hentikan semua service
	docker compose down

down-v: ## Hentikan semua service DAN hapus volume (reset data)
	docker compose down -v

logs: ## Lihat logs API secara live
	docker compose logs -f api

build: ## Build ulang image API
	docker compose build api

restart: ## Restart API saja (setelah perubahan kode)
	docker compose restart api

# ── Development (tanpa Docker, langsung di host) ───────────────
run: ## Jalankan API langsung (butuh Postgres & Redis jalan)
	go run ./cmd/api/...

# ── Database ──────────────────────────────────────────────────
migrate-up: ## Jalankan semua migration
	@which migrate > /dev/null 2>&1 || (echo "Install: go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest" && exit 1)
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down: ## Rollback 1 migration
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" down 1

migrate-status: ## Lihat status migration
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

# ── Code Quality ──────────────────────────────────────────────
test: ## Jalankan semua test
	go test -v -race ./...

lint: ## Jalankan linter (butuh golangci-lint)
	golangci-lint run ./...

tidy: ## Bersihkan go.mod dan go.sum
	go mod tidy

# ── Generate JWT Secrets ──────────────────────────────────────
secrets: ## Generate JWT secret yang aman
	@echo "JWT_ACCESS_SECRET=$(shell openssl rand -base64 48)"
	@echo "JWT_REFRESH_SECRET=$(shell openssl rand -base64 48)"
