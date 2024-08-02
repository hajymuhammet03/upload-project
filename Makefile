
PROJECT_NAME := upload-project
BUILD_DIR := bin
SRC_DIR := ./cmd/$(PROJECT_NAME)
PKG := ./...
POSTGRESQL_URL := "postgres://username:password@localhost:5432/dbname?sslmode=disable"

GO := go
GOTEST := $(GO) test
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOLINT := golangci-lint run
GOINSTALL := $(GO) install
GOFMT := gofmt
GOVET := $(GO) vet
GOMOD := $(GO) mod
MIGRATE := migrate

GOFLAGS := -mod=readonly

build: $(BUILD_DIR)/$(PROJECT_NAME)

$(BUILD_DIR)/$(PROJECT_NAME):
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) $(GOFLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME) $(SRC_DIR)

test:
	@echo "Running tests..."
	@$(GOTEST) $(PKG)

lint:
	@echo "Linting code..."
	@$(GOLINT) $(PKG)

fmt:
	@echo "Checking code format..."
	@$(GOFMT) -l $(PKG)

vet:
	@echo "Running go vet..."
	@$(GOVET) $(PKG)

clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

deps:
	@echo "Installing dependencies..."
	@$(GOINSTALL) $(PKG)

mod-init:
	@echo "Initializing Go modules..."
	@$(GOMOD) init

mod-tidy:
	@echo "Tidying Go modules..."
	@$(GOMOD) tidy

all: fmt vet lint test build

migrate_create:
	@echo "Creating new migration..."
	@$(MIGRATE) create -ext sql -dir db/migrations -seq init_schema

migrate_up:
	@echo "Applying migrations up..."
	@$(MIGRATE) -path db/migrations -database $(POSTGRESQL_URL) up

migrate_down:
	@echo "Reverting migrations down..."
	@$(MIGRATE) -path db/migrations -database $(POSTGRESQL_URL) down

compose_up:
	@echo "Starting Docker containers..."
	docker-compose up -d

compose_down:
	@echo "Stopping Docker containers..."
	docker-compose down

.PHONY: all build test lint fmt vet clean deps mod-init mod-tidy migrate_create migrate_up migrate_down compose_up compose_down
