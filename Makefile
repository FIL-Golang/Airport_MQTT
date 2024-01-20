.PHONY: all clean

all: build

clean:
	@echo "Cleaning cache..."
	go clean -modcache
	@echo "Cleaning completed."

build: clean
	@echo "Building executables to ./bin..."
	go build -o bin/ ./cmd/...
	@echo "Build completed."

build-images:
	@echo "Building docker image..."
	scripts/build-images.sh

	@echo "Build docker image completed."
