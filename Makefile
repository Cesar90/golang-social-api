include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

#migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

#migrate -verbose -path ./cmd/migrate/migrations -database "postgres://postgres:postgres@localhost:5432/social?sslmode=disable" up
.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)" down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt