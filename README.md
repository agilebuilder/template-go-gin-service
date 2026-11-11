# <%= projectName %>

<%= description %><% if (author) { %>

Author: <%= author %><% } %>

## Tech Stack

- Go 1.22
- [Gin](https://github.com/gin-gonic/gin): HTTP web framework
- [Viper](https://github.com/spf13/viper): configuration loading (YAML + environment variable overrides)
- Makefile: common build / run / test commands

## Project Structure

```
.
├── cmd/
│   └── server/          # Service entrypoint (main.go)
├── internal/            # Private application code (cannot be imported by external modules)
│   ├── config/          # Configuration loading (Viper)
│   ├── handler/         # HTTP handlers (controller layer)
│   ├── service/         # Business logic (service layer)
│   ├── model/           # Data structures and domain models
│   ├── middleware/      # Middleware (logging, panic recovery)
│   └── router/          # Route registration
├── pkg/
│   └── strutil/         # Example of a reusable public utility package
├── config.yaml          # Example configuration file
├── Makefile
└── go.mod               # module: <%= modulePath %>
```

## Quick Start

### Generate the Project with AgileBuilder

```bash
ag create --git-url <template-repo-url> --target <directory>
```

Fill in the variables as prompted to generate a complete project; `git init` and the first commit run automatically afterwards.

### Run Locally

```bash
go mod tidy        # download dependencies
make run           # or: go run ./cmd/server -config config.yaml
```

The service listens on port `<%= port %>` by default. Verify with:

```bash
curl http://localhost:<%= port %>/healthz
curl "http://localhost:<%= port %>/api/v1/hello?name=gin"
```

## Template Variables

| Variable | Description | Default |
| --- | --- | --- |
| `projectName` | Project name (used for the build artifact name, README title, etc.) | `demo-service` |
| `modulePath` | Go module path (the `module` line in go.mod; also the import prefix for internal packages) | `github.com/example/demo-service` |
| `description` | Project description | A Go HTTP service |
| `author` | Author name (optional) | (empty) |
| `port` | Server listen port (written to config.yaml) | `8080` |

## Common Commands

| Command | Description |
| --- | --- |
| `make build` | Compile the binary into `bin/` |
| `make run` | Run the service directly |
| `make test` | Run all tests (`go test ./...`) |
| `make vet` | Static analysis (`go vet ./...`) |
| `make tidy` | Tidy dependencies (`go mod tidy`) |
| `make clean` | Remove build artifacts |

## Configuration

Example `config.yaml`:

```yaml
server:
  port: <%= port %>
  mode: debug # debug | release | test
log:
  level: info
```

Every setting can be overridden by an environment variable: replace `.` with `_` and uppercase, e.g. `SERVER_PORT=9090` overrides `server.port`.

## Testing

The project ships with an `net/http/httptest`-based endpoint test (`internal/handler/hello_test.go`) and a utility package unit test (`pkg/strutil/strutil_test.go`):

```bash
go test ./...
```

## License

MIT
