.PHONY: build-sever
build:
	go build -o server ./cmd/server/.

.PHONY: run
run: build
	 ./server

.PHONY: create-new-migration
create-new-migration:
	migrate create -ext .sql -seq -dir data/migrations

.PHONY: local-lint
local-lint:
	golangci-lint run

.PHONY: generate-mock
generate-mock:
	mockgen -source ./internal/dal/base.go -destination ./internal/dal/mock/mock.go

