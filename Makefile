.PHONY: all clean shared install test linux macos windows cli

# Build directory
BUILD_DIR = build

# Default target
all: shared

# Create build directory if it doesn't exist
$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)

# Build shared library
shared: $(BUILD_DIR)
	@echo "Building shared library..."
	@go build -tags cshared -buildmode=c-shared -o $(BUILD_DIR)/libnbt.so .
	@echo "Shared library built: $(BUILD_DIR)/libnbt.so and $(BUILD_DIR)/libnbt.h"

# Build for different platforms
linux: $(BUILD_DIR)
	@echo "Building shared library for Linux..."
	@GOOS=linux GOARCH=amd64 go build -tags cshared -buildmode=c-shared -o $(BUILD_DIR)/libnbt.so .
	@echo "Linux shared library built: $(BUILD_DIR)/libnbt.so"

macos: $(BUILD_DIR)
	@echo "Building shared library for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -tags cshared -buildmode=c-shared -o $(BUILD_DIR)/libnbt.dylib .
	@echo "macOS shared library built: $(BUILD_DIR)/libnbt.dylib"

macos-arm: $(BUILD_DIR)
	@echo "Building shared library for macOS ARM64..."
	@GOOS=darwin GOARCH=arm64 go build -tags cshared -buildmode=c-shared -o $(BUILD_DIR)/libnbt.dylib .
	@echo "macOS ARM64 shared library built: $(BUILD_DIR)/libnbt.dylib"

windows: $(BUILD_DIR)
	@echo "Building shared library for Windows..."
	@GOOS=windows GOARCH=amd64 go build -tags cshared -buildmode=c-shared -o $(BUILD_DIR)/libnbt.dll .
	@echo "Windows shared library built: $(BUILD_DIR)/libnbt.dll"

# Build the original CLI tool
cli: $(BUILD_DIR)
	@echo "Building CLI tool..."
	@go build -o $(BUILD_DIR)/nbt main.go
	@echo "CLI tool built: $(BUILD_DIR)/nbt"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Install shared library (Linux/macOS)
install: shared
	@echo "Installing shared library..."
	@sudo cp $(BUILD_DIR)/libnbt.so /usr/local/lib/ 2>/dev/null || sudo cp $(BUILD_DIR)/libnbt.dylib /usr/local/lib/ 2>/dev/null || true
	@sudo cp $(BUILD_DIR)/libnbt.h /usr/local/include/ 2>/dev/null || true
	@sudo ldconfig 2>/dev/null || true
	@echo "Installation complete"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"
