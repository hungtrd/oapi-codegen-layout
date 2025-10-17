.PHONY: help generate build run clean test install-tools docker-build docker-run docker-compose-up docker-compose-down docker-compose-logs docker-compose-build docker-compose-restart docker-clean

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

install-tools: ## Install required tools
	@echo "Installing oapi-codegen..."
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@echo "Tools installed successfully"

generate: ## Generate API code from OpenAPI spec
	@echo "Generating API code from split specs..."
	@mkdir -p pkg/api/models
	@mkdir -p pkg/api/users
	@mkdir -p pkg/api/products
	@mkdir -p pkg/api/health
	@echo "Generating models..."
	@oapi-codegen -config api/configs/models.yaml api/specs/models.yaml
	@echo "Generating users handler..."
	@oapi-codegen -config api/configs/users.yaml api/specs/users.yaml
	@echo "Generating products handler..."
	@oapi-codegen -config api/configs/products.yaml api/specs/products.yaml
	@echo "Generating health handler..."
	@oapi-codegen -config api/configs/health.yaml api/specs/health.yaml
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
	@rm -rf pkg/api/
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

docker-compose-up: ## Start all services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose up -d
	@echo "Services started. API available at http://localhost:8080"
	@echo "Swagger UI available at http://localhost:8080/swagger/index.html"

docker-compose-down: ## Stop all services
	@echo "Stopping services..."
	@docker-compose down
	@echo "Services stopped"

docker-compose-logs: ## Show logs from all services
	@docker-compose logs -f

docker-compose-build: ## Build docker-compose services
	@echo "Building docker-compose services..."
	@docker-compose build
	@echo "Build complete"

docker-compose-restart: ## Restart all services
	@echo "Restarting services..."
	@docker-compose restart
	@echo "Services restarted"

docker-clean: ## Remove all containers, images, and volumes
	@echo "Cleaning Docker resources..."
	@docker-compose down -v
	@docker rmi oapi-codegen-layout:latest || true
	@echo "Docker resources cleaned"
