# Default command
default:
    @just --list --unsorted

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./cli/...

# Sync Go modules
tidy:
    go mod tidy
    cd cli && go mod tidy
    go work sync
    echo "Go workspace and modules synced successfully!"
