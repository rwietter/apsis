# Binary name
BINARY_NAME=apsis

# Go command
GO=go

# Build flags
BUILD_FLAGS=-v

.PHONY: all build clean run test fmt vet lint help

all: build

build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME)

clean:
	$(GO) clean
	rm -f $(BINARY_NAME)

run:
	$(GO) run .

test:
	$(GO) test -v ./...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

lint:
	@if [ ! -x "$$(command -v golangci-lint)" ]; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin; \
	fi
	golangci-lint run ./...

help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  clean    - Remove build artifacts"
	@echo "  run      - Run the application"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format code"
	@echo "  vet      - Run go vet"
	@echo "  lint     - Run golangci-lint"
	@echo "  help     - Show this help message"
