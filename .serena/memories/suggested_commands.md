# Suggested Commands

## Development Commands

### Testing
```bash
go test ./...                    # Run all tests
go test ./day01                  # Run tests for specific day
go test -v ./...                 # Run tests with verbose output
go test -cover ./...             # Run tests with coverage
```

### Code Quality
```bash
gofmt -l .                       # Check if code is formatted
gofmt -w .                       # Format all Go files
go vet ./...                     # Run static analysis
```
