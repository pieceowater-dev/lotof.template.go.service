# Define variables
APP_NAME = lotof.template.go.service
BUILD_DIR = bin
MAIN_FILE = main.go
GOFMT = gofmt
GO_TEST = go test
GO_COVERAGE = $(GO_TEST) -cover
MIGRATION_DIR = internal/core/db/migrations
DB_DSN = $(shell grep DATABASE_DSN .env | cut -d '"' -f2)

export PATH := /usr/local/bin:$(PATH)

# Create a build directory if it doesn't exist
.PHONY: all clean build run test update migration migrate check-migration setup install-flyway install-atlas install-postgres install-atlas-cli

all: build

# Update dependencies
update:
	go mod tidy

# Build the project
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Run unit tests
test:
	$(GO_TEST)

# Run unit tests with coverage
test-coverage:
	$(GO_COVERAGE)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Install Atlas CLI using Homebrew
install-atlas-cli:
	@brew install ariga/tap/atlas

# Generate a new migration file with the timestamp using Atlas
migration:
	@mkdir -p $(MIGRATION_DIR); \
		PATH=/usr/local/bin:$$PATH atlas migrate diff --env gorm; \
		echo "Migration files generated in $(MIGRATION_DIR)"; \
		git add $(MIGRATION_DIR)/*

# Apply migrations using Atlas
migrate:
	@PATH=/usr/local/bin:$$PATH atlas migrate apply --url "$(DB_DSN)" --dir="file://$(shell pwd)/$(MIGRATION_DIR)"

# Generate a new migration & Apply migrations using Atlas
db-sync:migration migrate

# Setup the environment
setup: install-atlas-cli
	@echo "Setup completed!"; \
		go mod tidy