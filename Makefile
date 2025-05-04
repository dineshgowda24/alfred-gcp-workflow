.PHONY: build release

build:
	@echo "Building the project..."
	go build -o run
	@echo "Build complete!"

release:
	@echo "Building the project for release..."
	./tools/release.sh
	@echo "Release build complete."

