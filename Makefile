.PHONY: all build run clean

# Default target: build the application
all: build

# Build the application binary
build:
	@echo "Building application..."
	go build -o bin/expense-tracker cmd/api/main.go

# Run the application (without building a binary)
run:
	@echo "Running application..."
	go run cmd/api/main.go

clean:
	@echo "Cleaning up..."
	rm -rf bin
