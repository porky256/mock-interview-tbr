.PHONY: build
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
	mockgen -source ./internal/user/repo_interface.go -destination ./internal/user/mock/repo_mock.go

