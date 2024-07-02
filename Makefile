SHELL=/bin/bash -o pipefail

.PHONY: lint
lint:
	@echo "Starting formatting..."
	go run golang.org/x/tools/cmd/goimports@latest -w -l -local "github.com/starquake/montyhall" .
	go run github.com/daixiang0/gci@latest write -s standard -s default -s "prefix(github.com/starquake/montyhall)" .
	go run mvdan.cc/gofumpt@latest -w -l .
	golangci-lint run --fix
	@echo "Finished formatting..."

.PHONY: build
build:
	mkdir -p bin
	env CGO_ENABLED=0 go build -v -o bin/montyhall ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: database
database:
	GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=${TRIV_DSN} goose -dir migrations up