.PHONY: help build up down run fmt lint clean

# Default help target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build the Go project
build: ## Build the Go application
	go build -o ./tmp/main ./cmd/app

# Run the app directly
run: build ## Build and run the app
	./tmp/main

# Docker: bring services up
up: ## Start docker-compose in detached mode
	docker-compose up -d

# Docker: stop services
down: ## Stop docker-compose
	docker-compose down

# Clean build artifacts
clean: ## Remove built binary
	rm -rf ./tmp/main

test: ## Test all test files
	go test ./... 
