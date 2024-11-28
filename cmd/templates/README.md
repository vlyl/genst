# {{.ProjectName}}

A modern Go HTTP server with essential components and best practices.

## Features

- HTTP server using [Gin](https://github.com/gin-gonic/gin) framework
- API documentation with [gin-swagger](https://github.com/swaggo/gin-swagger)
- PostgreSQL database integration using [GORM](https://gorm.io)
- Configuration management with [Viper](https://github.com/spf13/viper)
- Command-line interface using [Cobra](https://github.com/spf13/cobra)
- Structured logging with [Zap](https://github.com/uber-go/zap)
- Log rotation support
- Example API endpoints
- Makefile for common operations
- Startup script for deployment

## Quick Start

1. Install dependencies:
   ```bash
   make deps
   ```

2. Configure the application:
   - Copy `config/config.yaml.example` to `config/config.yaml`
   - Edit the configuration according to your environment

3. Run the server:
   ```bash
   make run
   ```

4. Test the server:
   ```bash
   curl http://localhost:8080/ping
   ```

5. View API documentation:
   - Open http://localhost:8080/swagger/index.html in your browser

## Project Structure

```
.
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

## Development

### Adding New API Endpoints

1. Create a new handler in `internal/api/`:
   ```go
   func NewHandler(c *gin.Context) {
       // Handler implementation
   }
   ```

2. Register the route in `internal/api/router.go`:
   ```go
   router.GET("/new-endpoint", NewHandler)
   ```

3. Add swagger documentation comments
4. Generate swagger docs:
   ```bash
   make swagger
   ```

### Database Operations

1. Define models in `internal/model/`:
   ```go
   type User struct {
       gorm.Model
       Name  string
       Email string
   }
   ```

2. Use GORM for database operations:
   ```go
   db := database.Get()
   db.AutoMigrate(&User{})
   ```

### Available Make Commands

- `make build`: Build the application
- `make run`: Run the application
- `make test`: Run tests
- `make swagger`: Generate swagger documentation
- `make deps`: Install dependencies
- `make lint`: Run linter
- `make clean`: Clean build artifacts

### Configuration

The application uses a YAML configuration file with the following structure:

```yaml
server:
  port: 8080
  mode: debug  # debug or release

database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  dbname: {{.ProjectName}}
  sslmode: disable

log:
  level: info
  filename: logs/app.log
  maxsize: 100    # megabytes
  maxage: 7       # days
  maxbackups: 3   # number of backups
```

### Environment Variables

All configuration values can be overridden using environment variables:
- `SERVER_PORT`: HTTP server port
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `LOG_LEVEL`: Logging level

## License

MIT License
