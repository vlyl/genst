.PHONY: all build clean run test swagger

# Build variables
BINARY_NAME=server
BUILD_DIR=bin

all: clean build

build:
	@echo "Building..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

run: build
	@echo "Running..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

test:
	@echo "Testing..."
	@go test -v ./...

swagger:
	@echo "Checking swag installation..."
	@which swag > /dev/null || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	@echo "Generating Swagger docs..."
	@swag init -g internal/api/router.go -o docs

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

.PHONY: lint
lint:
	@echo "Linting..."
	@golangci-lint run 