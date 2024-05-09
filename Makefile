# Simple Makefile for a Go project

# Build the application
all: build

build_finder:
	@echo "Building..."
	# validate if the folder files if not exists, if not create the folder
	@if [ ! -d "files" ]; then \
		mkdir files; \
		touch files/directories.json; \
		echo "[]" > files/directories.json; \
	fi \

	# validate if directories.json if not exists, if not create the file as '[]'
	@if [ ! -f "files/directories.json" ]; then \
		touch files/directories.json; \
		echo "[]" > files/directories.json; \
	fi \
	
	@go build -o tmux-go cmd/finder/main.go

