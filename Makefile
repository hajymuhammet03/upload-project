# Proje değişkenleri
PROJECT_NAME := upload-project
BUILD_DIR := bin
SRC_DIR := ./cmd/$(PROJECT_NAME)
PKG := ./...
POSTGRESQL_URL := "postgres://username:password@localhost:5432/dbname?sslmode=disable"

# Komutlar
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

# GO build bayrakları
GOFLAGS := -mod=readonly

# Ana build hedefi
build: $(BUILD_DIR)/$(PROJECT_NAME)

# Binary build
$(BUILD_DIR)/$(PROJECT_NAME):
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) $(GOFLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME) $(SRC_DIR)

# Testleri çalıştır
test:
	@echo "Running tests..."
	@$(GOTEST) $(PKG)

# Linter çalıştır
lint:
	@echo "Linting code..."
	@$(GOLINT) $(PKG)

# Format kontrolü
fmt:
	@echo "Checking code format..."
	@$(GOFMT) -l $(PKG)

# go vet çalıştır
vet:
	@echo "Running go vet..."
	@$(GOVET) $(PKG)

# Build dosyalarını temizle
clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Bağımlılıkları yükle
deps:
	@echo "Installing dependencies..."
	@$(GOINSTALL) $(PKG)

# Go modüllerini başlat
mod-init:
	@echo "Initializing Go modules..."
	@$(GOMOD) init

# Go modüllerini düzenle
mod-tidy:
	@echo "Tidying Go modules..."
	@$(GOMOD) tidy

# Varsayılan hedef
all: fmt vet lint test build

# Göç dosyası oluştur
migrate_create:
	@echo "Creating new migration..."
	@$(MIGRATE) create -ext sql -dir db/migrations -seq init_schema

# Göçleri uygulama
migrate_up:
	@echo "Applying migrations up..."
	@$(MIGRATE) -path db/migrations -database $(POSTGRESQL_URL) up

# Göçleri geri alma
migrate_down:
	@echo "Reverting migrations down..."
	@$(MIGRATE) -path db/migrations -database $(POSTGRESQL_URL) down

# Docker konteynerlarını başlat
compose_up:
	@echo "Starting Docker containers..."
	docker-compose up -d

# Docker konteynerlarını durdur
compose_down:
	@echo "Stopping Docker containers..."
	docker-compose down

.PHONY: all build test lint fmt vet clean deps mod-init mod-tidy migrate_create migrate_up migrate_down compose_up compose_down
