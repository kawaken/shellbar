# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Communication

- Responses must be in Japanese in CI/Issues
- Use `Co-Authored-By` when committing

## Build and Development Commands

```bash
# Build the sbar binary
go build -o sbar ./cmd/sbar

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

# Run go vet
go vet ./...

# Run staticcheck (install: go install honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck ./...

# Run security scan with gosec
gosec ./...

# Check for vulnerabilities
govulncheck ./...
```

## Project Architecture

Shellbar is a CLI tool that displays a persistent status bar at the bottom of the terminal during long-running processes. The command name is `sbar`.

### Core Components

- **Main Entry Point**: `cmd/sbar/main.go` - CLI entry point
- **Core Logic**: `shellbar.go` - Main Shellbar struct and Run method
- **Command Parsing**: `command.go` - Command structure and argument parsing with built-in commands (version, help)
- **Configuration**: `config/` - TOML-based configuration management for status bar format and command definitions

### Key Design Decisions

1. **PTY Control**: The tool will wrap processes using a pseudo-terminal to preserve the status bar at the bottom line
2. **Command Execution Model**: External commands are defined in TOML config and executed periodically in separate goroutines
3. **Format String Expansion**: Status bar uses `{variable}` placeholders that are replaced with command outputs
4. **Configuration Priority**: Command-line args > local config > global config > defaults

### Configuration

Configuration files are TOML format (`shellbar.toml`) and support:
- Format string for status bar layout
- Global refresh rate
- Default command timeout/interval settings
- Individual command definitions with custom intervals

Example location: `examples/shellbar.toml`

## Testing Strategy

- Unit tests use Japanese descriptions for test cases
- Focus on configuration parsing and command execution
- No test files exist yet for main logic (shellbar.go, command.go)

## CI/CD

GitHub Actions workflow (`.github/workflows/ci.yml`) runs:
- Multiple Go versions (1.23.x, 1.24.x)
- Tests with race detection and coverage
- Static analysis with staticcheck
- Security scanning with gosec and govulncheck
- Coverage threshold monitoring (80% target, not enforced)

## Development Best Practices

- When specifying a Go version, check the go.mod version