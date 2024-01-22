.PHONY: all clean

all: build

clean:
	@echo "Cleaning cache..."
	go clean -modcache
	@echo "Cleaning completed."

build: clean
	@echo "Building project..."
	go build -o . -v ./...
	@echo "Build completed."

test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Tests completed."
