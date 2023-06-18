.PHONY: build-sever
build:
	go build -o server ./cmd/server/.

.PHONY: run
run: build
	 ./server

.PHONY: create-new-migration
create-new-migration:
	migrate create -ext .sql -seq -dir data/migrations


