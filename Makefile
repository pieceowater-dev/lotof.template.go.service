# Define variables
APP_NAME = lotof.template.go.service
BUILD_DIR = bin
MAIN_FILE = main.go
GOFMT = gofmt
GO_TEST = go test
GO_COVERAGE = $(GO_TEST) -cover

# Create a build directory if it doesn't exist
.PHONY: all clean build run test update

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