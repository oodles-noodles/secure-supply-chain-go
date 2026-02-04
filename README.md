# Secure Supply Chain Go Demo

A demonstration repository showcasing secure software supply chain practices in Go, including build attestations and Software Bill of Materials (SBOM) generation using GitHub Actions.

## Features

- **RESTful API Server**: A realistic web application with multiple endpoints
- **Structured Logging**: JSON-formatted logging with logrus
- **Prometheus Metrics**: Built-in metrics collection and exposure
- **Configuration Management**: Environment-based configuration with Viper
- **Database Integration**: PostgreSQL driver support
- **Middleware**: Request logging and metrics collection
- **Automated Builds**: GitHub Actions workflow for building the application
- **SBOM Generation**: Automatic creation of Software Bill of Materials using Anchore's SBOM Action
- **Build Attestations**: Cryptographic attestations for both SBOM and build provenance using GitHub's attestation features

## Application Structure

```
.
├── main.go                      # Application entry point
├── internal/
│   ├── config/                  # Configuration management
│   ├── database/                # Database connection handling
│   ├── handlers/                # HTTP request handlers
│   └── middleware/              # HTTP middleware
└── pkg/
    └── logger/                  # Logging utilities
```

## Dependencies

The application uses several production-grade Go libraries:

- **gorilla/mux**: HTTP router and request multiplexer
- **logrus**: Structured logging
- **viper**: Configuration management
- **prometheus/client_golang**: Metrics collection and exposition
- **lib/pq**: PostgreSQL driver
- **google/uuid**: UUID generation
- **yaml.v3**: YAML processing

## API Endpoints

The application exposes the following RESTful endpoints:

- `GET /` - Welcome message with API information
- `GET /health` - Health check endpoint (returns service status)
- `GET /info` - Application information and features
- `GET /api/users` - List users (mock data)
- `GET /metrics` - Prometheus metrics endpoint

## Running the Application

### Build the application:

```bash
go build -o secure-supply-chain-demo main.go
```

### Run the server:

```bash
./secure-supply-chain-demo
```

The server will start on `http://localhost:8080` by default.

### Configuration

Configure the application using environment variables:

- `SERVER_PORT`: Server port (default: 8080)
- `SERVER_HOST`: Server host (default: 0.0.0.0)
- `LOGGING_LEVEL`: Log level (default: info)
- `DATABASE_HOST`: PostgreSQL host (default: localhost)
- `DATABASE_PORT`: PostgreSQL port (default: 5432)
- `DATABASE_USER`: Database user (default: postgres)
- `DATABASE_PASSWORD`: Database password
- `DATABASE_DBNAME`: Database name (default: secureapp)

### Testing the API

```bash
# Health check
curl http://localhost:8080/health

# Application info
curl http://localhost:8080/info

# List users
curl http://localhost:8080/api/users

# View metrics
curl http://localhost:8080/metrics
```

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## GitHub Actions Workflow

The workflow (`.github/workflows/build.yml`) performs the following steps:

1. **Checkout**: Retrieves the source code
2. **Setup Go**: Installs Go 1.21
3. **Build**: Compiles the binary
4. **Generate SBOM**: Creates a Software Bill of Materials in SPDX JSON format
5. **Attest SBOM**: Creates a cryptographic attestation for the SBOM
6. **Attest Build Provenance**: Creates a cryptographic attestation for the build artifact
7. **Upload Artifacts**: Stores the binary and SBOM as workflow artifacts

## Attestations

This demo uses GitHub's attestation features to provide:

- **SBOM Attestation**: Cryptographically signed record of all dependencies
- **Build Provenance Attestation**: Cryptographically signed record of how the artifact was built

These attestations can be verified using the GitHub CLI or API to ensure the integrity and provenance of the built artifacts.

## Requirements

- Go 1.21 or higher
- GitHub Actions (for automated builds and attestations)

## License

This is a demonstration project for educational purposes.