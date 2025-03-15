# Urusai

Urusai is a Go implementation of [noisy](https://github.com/1tayH/noisy), a simple random HTTP/DNS internet traffic noise generator. It generates random HTTP/DNS traffic noise in the background while you go about your regular web browsing, to make your web traffic data less valuable for selling and for extra obscurity.

## Features

- Generates random HTTP/DNS traffic by crawling websites
- Configurable via JSON configuration file
- Customizable user agents, root URLs, and blacklisted URLs
- Adjustable crawling depth and sleep intervals
- Optional timeout setting

## Installation

```bash
# Clone the repository
git clone https://github.com/calpaliu/urusai.git

# Navigate to the project directory
cd urusai

# Build the project
go build -o urusai
```

## Usage

```bash
# Run with built-in default configuration
./urusai

# Run with custom configuration file
./urusai --config config.json

# Show help
./urusai --help
```

### Command Line Arguments

- `--config`: Path to the configuration file (optional, uses built-in default configuration if not specified)
- `--log`: Logging level (default: "info")
- `--timeout`: For how long the crawler should be running, in seconds (optional, 0 means no timeout)

## Configuration

Urusai comes with a built-in default configuration, but you can also provide your own custom configuration file. The configuration is in JSON format with the following structure:

```json
{
    "max_depth": 25,
    "min_sleep": 3,
    "max_sleep": 6,
    "timeout": 0,
    "root_urls": [
        "https://www.wikipedia.org",
        "https://www.github.com"
    ],
    "blacklisted_urls": [
        ".css",
        ".ico",
        ".xml"
    ],
    "user_agents": [
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
    ]
}
```

## For Developers

### Development

Urusai is developed using standard Go practices. Here are some commands that will help you during development:

```bash
# Run the project directly without building
go run main.go --config config.json

# Run with a specific log level
go run main.go --config config.json --log debug

# Run with a timeout (in seconds)
go run main.go --config config.json --timeout 300
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate a coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Building

```bash
# Build for the current platform
go build -o urusai

# Build for a specific platform (e.g., Linux)
GOOS=linux GOARCH=amd64 go build -o urusai-linux-amd64

# Build for multiple platforms
GOOS=darwin GOARCH=amd64 go build -o urusai-macos-amd64
GOOS=windows GOARCH=amd64 go build -o urusai-windows-amd64.exe
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run linter (requires golint)
go install golang.org/x/lint/golint@latest
golint ./...

# Run static analysis (requires staticcheck)
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
