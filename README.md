# Genst - Golang Server Starter Generator

A command-line tool for generating Go HTTP server projects with essential components and best practices.

## Features

Generated project includes:
- HTTP server using [Gin](https://github.com/gin-gonic/gin)
- API documentation with [gin-swagger](https://github.com/swaggo/gin-swagger)
- PostgreSQL database integration using [GORM](https://gorm.io)
- Configuration management with [Viper](https://github.com/spf13/viper)
- Command-line interface using [Cobra](https://github.com/spf13/cobra)
- Structured logging with [Zap](https://github.com/uber-go/zap)
- Log rotation support
- Example API endpoints
- Makefile for common operations
- Startup script for deployment

## Installation

```bash
go install github.com/vlyl/genst@latest
```

## Usage

1. Create a new project:
```bash
genst new myproject
```

2. The command will generate a complete project structure:
```
myproject/
├── cmd/
│   └── server/       # Application entry point
├── internal/
│   ├── api/         # API handlers and routes
│   ├── config/      # Configuration management
│   ├── middleware/  # HTTP middleware
│   ├── model/       # Database models
│   └── service/     # Business logic
├── pkg/
│   ├── logger/      # Logging utilities
│   └── database/    # Database utilities
├── config/          # Configuration files
├── scripts/         # Utility scripts
└── Makefile        # Build and management commands
```

3. Navigate to your project and start development:
```bash
cd myproject
make deps    # Install dependencies
make run     # Run the server
```

4. Access the API documentation at http://localhost:8080/swagger/index.html

## Development

### Project Structure

```
.
├── README.md           # This file
├── cmd/
│   ├── generate.go     # Generator implementation
│   └── templates/      # Project templates
├── go.mod
└── main.go            # Entry point
```

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/vlyl/genst.git
cd genst
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the binary:
```bash
go build
```

## License

MIT License