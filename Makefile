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
	mockgen -source ./internal/user/userrepo/repo_interface.go -destination ./internal/user/userrepo/mock/repo_mock.go
	mockgen -source ./internal/skill/skillrepo/repo_interface.go -destination ./internal/skill/skillrepo/mock/repo_mock.go
	mockgen -source ./internal/match/matchrepo/repo_interface.go -destination ./internal/match/matchrepo/mock/repo_mock.go

