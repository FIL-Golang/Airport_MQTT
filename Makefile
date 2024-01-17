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

build-image: build
	@echo "Building docker image..."
	script/build-image.sh

	@echo "Build docker image completed."
