# Include .env if exists
-include .env
export

CURRENT_DIR := $(shell pwd)
POSTGRES_USER := postgres
POSTGRES_PASSWORD := root
POSTGRES_HOST := localhost
POSTGRES_PORT := 5432
POSTGRES_DATABASE := yalp_db

DB_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable

# Run service
.PHONY: run
run:
	go run cmd/main.go

# Generate protobuf
.PHONY: proto-gen
proto-gen:
	./scripts/gen_proto.sh

# Database migrations
.PHONY: migrate migrate-up migrate-down migrate-force migrate-file
migrate:
	migrate -source file://migrations -database ${DB_URL} up

migrate-up:
	migrate -path internal/db/migrations -database ${DB_URL} -verbose up

migrate-down:
	migrate -path internal/db/migrations -database ${DB_URL} -verbose down

migrate-force:
	migrate -path internal/db/migrations -database ${DB_URL} -verbose force 1

migrate-file:
	migrate create -ext sql -dir internal/db/migrations -seq yalp_create_psql

# Swagger initialization
.PHONY: swag-init
swag-gen:
	~/go/bin/swag init -g ./internal/app/app.go -o api/docs force 1