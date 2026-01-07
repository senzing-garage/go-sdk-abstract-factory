# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library implementing the Abstract Factory pattern for the Senzing SDK. It provides a unified interface (`senzing.SzAbstractFactory`) with three concrete implementations:

- **Core** (`CreateCoreAbstractFactory`): Native Senzing SDK calls via C library bindings
- **gRPC** (`CreateGrpcAbstractFactory`): Remote calls to a Senzing gRPC server
- **Mock** (`CreateMockAbstractFactory`): Mock implementation for unit testing

## Prerequisites

The Senzing C library must be installed:

- Shared objects: `/opt/senzing/er/lib`
- SDK headers: `/opt/senzing/er/sdk/c`
- Configuration: `/etc/opt/senzing`

Set `LD_LIBRARY_PATH=/opt/senzing/er/lib` before running tests.

## Common Commands

```bash
# Install development tools (one-time)
make dependencies-for-development

# Update Go dependencies
make dependencies

# Run linting (golangci-lint, govulncheck, cspell)
make lint

# Run tests (requires setup first)
make clean setup test

# Run tests with coverage
make clean setup coverage

# Auto-fix lint issues
make fix

# Run a single test
go test -v -run TestSzfactorycreator_CreateCoreAbstractFactory ./szfactorycreator/...
```

## Testing Notes

- Tests require a running Senzing gRPC server (started automatically by `make setup` via Docker)
- The gRPC server runs on `localhost:8261`
- Tests use SQLite database at `/tmp/sqlite/G2C.db`
- Use `SENZING_LOG_LEVEL=TRACE` for verbose test output

## Code Style

- Uses `golangci-lint` with extensive linter configuration at `.github/linters/.golangci.yaml`
- Maximum line length: 120 characters
- Uses `gofumpt` for formatting
- Test files use `_test` package suffix (e.g., `szfactorycreator_test`)
- Tests should use `test.Parallel()` and `test.Context()`
