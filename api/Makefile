.PHONY: os gqlgen build lint test updb downdb dropdb

OSFLAG :=
GOARCH :=
VERSION?="1.0.0"
POSTGRES_URL?="$(POSTGRES_CONNECTION_STRING)"

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OSFLAG = "linux"
	GOARCH = "amd64"
endif
ifeq ($(UNAME_S),Darwin)
	OSFLAG = "darwin"
	GOARCH = "arm64"
endif

ifeq ($(POSTGRES_URL),"")
	POSTGRES_URL="postgres://backend_user:backend_password@db:5432/backend_db?sslmode=disable"
endif

os:
	@echo ${OSFLAG}

gen:
	@go run github.com/99designs/gqlgen generate
	@go generate ./ent

db:
	docker compose up -d db

run:
	docker compose up -d

stop:
	docker compose stop

teardown:
	docker compose down -v --remove-orphans
	docker compose rm --force --stop -v

hashdb:
	docker compose run --rm atlas migrate hash

migratedbdryrun: hashdb
	docker compose run --rm atlas migrate apply --url $(POSTGRES_URL) --dry-run

migratedb: hashdb migratedbdryrun
	docker compose run --rm atlas migrate apply --url $(POSTGRES_URL)

api: build
	./server api

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(OSFLAG) GOARCH=$(GOARCH) go build -ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT=$(COMMIT) -X main.DATE=$(DATE) -w -s" -v -o server cmd/main.go