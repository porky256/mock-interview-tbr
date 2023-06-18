include .env
export

.PHONY: local-db-up
local-db-up:
	docker-compose up db

.PHONY: local-db-down
local-db-down:
	docker-compose stop db


migration-up:
	migrate -path ./../migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@0.0.0.0:5432/$(POSTGRES_DB)?sslmode=disable' up

migration-down:
	yes | migrate -path ./../migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@0.0.0.0:5432/$(POSTGRES_DB)?sslmode=disable' down

migration-force:
	migrate -path ./../migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@0.0.0.0:5432/$(POSTGRES_DB)?sslmode=disable' force 1
