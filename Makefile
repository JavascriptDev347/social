-include .env
export

MIGRATION_DIR=./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATION_DIR) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATION_DIR) -database "$(DB_MIGRATOR_ADDR)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATION_DIR) -database "$(DB_MIGRATOR_ADDR)" down $(filter-out $@,$(MAKECMDGOALS))


#migrate -path ./cmd/migrate/migrations -database $env:DB_MIGRATOR_ADDR drop -f
.PHONY: migrate-drop
migrate-drop:
	@migrate -path $(MIGRATION_DIR) -database "$(DB_MIGRATOR_ADDR)" drop -f
