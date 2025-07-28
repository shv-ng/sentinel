# Makefile for Docker Compose operations and proto generation

DC := docker-compose -f ./deploy/docker-compose.yml --env-file .env

.PHONY: migrate up-d up build down redis-cli postgres-cli gen-proto

## Run migration container (exits after completion)
migrate:
	$(DC) --profile migration up migration --abort-on-container-exit

## Start services in detached mode
up-d:
	$(DC) up -d

## Start services in foreground
up:
	$(DC) up

## Build services without cache
build:
	$(DC) build --no-cache

## Stop and remove containers, networks, volumes, and images
down:
	$(DC) down

## Access Redis CLI
redis-cli:
	$(DC) exec redis redis-cli

## Access PostgreSQL CLI
postgres-cli:
	$(DC) exec postgres psql -U postgres

## Generate Go gRPC code from protobuf files
gen-proto:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ --proto_path=. ./proto/proto/*.proto
