PROJECT_NAME := "jocasta"

# Default command
_default:
    @just --list --unsorted

# CLI run wrapper
cli *args:
    @go run main.go {{ args }}

# Execute unit tests
test:
    @go clean -testcache
    @echo "Running unit tests!"
    go test -cover ./cli/...

# Sync Go modules
tidy:
    go mod tidy
    cd cli && go mod tidy
    go work sync

# Build CLI binary
build:
    #!/usr/bin/env bash
    echo "Building {{ PROJECT_NAME }} binary..."
    go mod download all
    VERSION=$(jq -r .version specs.json)
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=${VERSION}" -o ./jocasta main.go
    echo "Built binary for {{ PROJECT_NAME }} ${VERSION} successfully!"
