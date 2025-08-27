# Makefile for DJINN backend

# Database connection
DATABASE_URL ?= postgres://djinn_user:djinn_password@localhost:5432/djinn?sslmode=disable

# Atlas commands
.PHONY: migrate-up
migrate-up: ## Apply all pending migrations
	@echo "Applying migrations..."
	@cat migrations/*.sql | docker exec -i djinn-postgres psql -U djinn_user -d djinn

.PHONY: migrate-down
migrate-down: ## Reset database
	@echo "Resetting database..."
	@docker exec djinn-postgres psql -U djinn_user -d djinn -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public; GRANT ALL ON SCHEMA public TO djinn_user;"

.PHONY: migrate-status
migrate-status: ## Check migration status
	@echo "Database tables:"
	@docker exec djinn-postgres psql -U djinn_user -d djinn -c "\\dt"

.PHONY: db-console
db-console: ## Open PostgreSQL console
	@docker exec -it djinn-postgres psql -U djinn_user -d djinn

.PHONY: test-db
test-db: ## Test database operations
	@go run cmd/test/main.go

.PHONY: sqlc-generate
sqlc-generate: ## Generate sqlc code
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

.PHONY: mockery-generate
mockery-generate: ## Generate mocks with mockery
	@go run github.com/vektra/mockery/v2@latest

.PHONY: dataloader-generate
dataloader-generate: ## Generate dataloader code
	@echo "Generating dataloaders..."
	@cd internal/dataloader && go generate ./...

.PHONY: generate
generate: sqlc-generate mockery-generate dataloader-generate ## Generate all code (sqlc, mockery, dataloaders)

.PHONY: test
test: generate ## Run all tests
	@go test -v ./...

.PHONY: test-coverage
test-coverage: generate ## Run tests with coverage
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## Run linting
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest ./...

.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help