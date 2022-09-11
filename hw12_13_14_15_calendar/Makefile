BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"
MIGRATOR := "./bin/db_migration"
SCHEDULER := "./bin/calendar_scheduler"
SENDER := "./bin/calendar_sender"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(MIGRATOR) -ldflags "$(LDFLAGS)" ./cmd/db_migration
	go build -v -o $(SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/calendar_scheduler
	go build -v -o $(SENDER) -ldflags "$(LDFLAGS)" ./cmd/calendar_sender

run: build
	$(BIN) -config ./configs/config.yaml
	$(BIN) -config ./configs/config_rmq.yaml
	$(BIN) -config ./configs/config_schedule.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

# install-lint-deps:
#	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1


# lint: install-lint-deps
#	golangci-lint run ./...

.PHONY: build run build-img run-img version test lint

.PHONY: generate
generate:
	go generate ./...

createdb:
	PGHOST=localhost PGUSER=postgres PGPORT=5432 createdb сalendar
	PGHOST=localhost PGUSER=postgres PGPORT=5432 createdb сalendar_test


# migrateUp: почему-то так не работает
#	$(MIGRATOR) -config ./configs/config.yaml -dir=./migrations up

migrateUp:
	goose -dir migrations postgres "user=postgres password=password dbname=calendar sslmode=disable" up

migrateDowm:
	goose -dir migrations postgres "user=postgres password=password dbname=calendar sslmode=disable" down

start-evans:
	evans --proto api/*.proto repl