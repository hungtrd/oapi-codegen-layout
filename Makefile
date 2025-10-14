.PHONY: help generate build run clean test install-tools

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

install-tools: ## Install required tools
	@echo "Installing oapi-codegen..."
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@echo "Tools installed successfully"

generate: ## Generate API code from OpenAPI spec
	@echo "Generating API code..."
	@mkdir -p pkg/api
	@oapi-codegen -config api/oapi-codegen.yaml api/openapi.yaml
	@echo "Code generation complete"

build: generate ## Build the application
	@echo "Building application..."
	@go build -o build/server ./cmd/server
	@echo "Build complete: build/server"

run: generate ## Run the application
	@echo "Running application..."
	@go run ./cmd/server/main.go

clean: ## Clean build artifacts and generated files
	@echo "Cleaning..."
	@rm -rf build/
	@rm -rf pkg/api/*.gen.go
	@echo "Clean complete"

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Tests complete. Coverage report: coverage.html"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting complete"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...
	@echo "Linting complete"

dev: generate ## Run in development mode with hot reload (requires air)
	@echo "Starting development server..."
	@air

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t oapi-codegen-layout:latest .
	@echo "Docker image built"

docker-run: docker-build ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 8080:8080 oapi-codegen-layout:latest
