# Load environment variables from .env file
include .env.local
export $(shell sed 's/=.*//' .env.local)
# Run the application
run:
	go run ./cmd/ 

# Run tests for the application
test:
	go test ./...

# Update swagger document
swag: 
	swag init -g cmd/main.go

admin-run:
	go run ./db/seed

# Use goos for dbmigration	
# postgres://pgsql:@localhost:$5432/$test_db1?sslmode=disable"
db-mig-create:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) create $(filter-out $@,$(MAKECMDGOALS)) sql
db-seed-create:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(SEED_DIR) create $(filter-out $@,$(MAKECMDGOALS)) sql
db-seed:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" \
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(SEED_DIR) $(cmd)
db-mig:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" \
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) $(cmd)
