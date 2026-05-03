DB_DRIVER = postgres
MIGRATIONS_PATH = ./migrations

.PHONY: migration
migration:
	@goose -s -dir $(MIGRATIONS_PATH) $(DB_DRIVER) $(DB_URL) create $(filter-out $@,$(MAKECMDGOALS)) sql

.PHONY: migrate-up
migrate-up:
	@goose -s -dir $(MIGRATIONS_PATH) $(DB_DRIVER) $(DB_URL) up

.PHONY: migrate-down
migrate-down:
	@goose -s -dir $(MIGRATIONS_PATH) $(DB_DRIVER) $(DB_URL) down
