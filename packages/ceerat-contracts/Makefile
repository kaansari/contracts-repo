SHELL := /bin/bash

.PHONY: help tidy test build proto

help:
	@echo "ceerat-contracts commands:"
	@echo "  make tidy   Run go mod tidy"
	@echo "  make test   Run all tests"
	@echo "  make build  Compile packages"
	@echo "  make proto  Regenerate protobuf Go files"

tidy:
	go mod tidy

test:
	go test ./...

build:
	go build ./...

proto:
	protoc --go_out=. --go-grpc_out=. proto/auth/auth.proto proto/patient/patient.proto
