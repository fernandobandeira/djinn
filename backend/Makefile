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
	@sqlc generate

.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help