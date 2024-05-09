# Simple Makefile for a Go project

# Build the application
all: build

build_finder:
	@echo "Building..."
	@go build -o tmuxs cmd/finder/main.go

build: build_finder

