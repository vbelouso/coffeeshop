# Variables
APP_NAME := coffeeshop
APP_PORT := 8080
REGISTRY := <PLACEHOLDER>
VERSION := 0.0.1

.PHONY: help build run push

help:
	@echo "Makefile for building and pushing coffeshop app"
	@echo ""
	@echo "Usage:"
	@echo "  make build		build the app"
	@echo "  make run		run the app"
	@echo "  make sqlc		generate database models, packages and queries"
	@echo "  make push		push the container image to registry"

docker-compose-up:
	@echo "Starting required services with docker-compose..."
	@docker-compose up -d

docker-compose-down:
	@echo "Stopping required services with docker-compose..."
	@docker-compose down

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME)

run: docker-compose-up build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)

# doesn't work, change docker-compose.yaml to include the app
#stop: docker-compose-down
#	@echo "Stopping $(APP_NAME)..."
#	@./$(APP_NAME)

sqlc:
	@echo "Generate database models, packages and queries"
	@sqlc generate --experimental

api_test:
	@echo "Testing $(APP_NAME) with curl commands..."
	@echo ""
	@echo "GET /customers/1"
	@curl http://localhost:$(APP_PORT)/customers/1
	@echo ""
	@echo "GET /customers"
	@curl http://localhost:$(APP_PORT)/customers
	@echo ""
	@echo "GET /orders/1"
	@curl http://localhost:$(APP_PORT)/orders/1
	@echo ""
	@echo "GET /orders"
	@curl http://localhost:$(APP_PORT)/orders
