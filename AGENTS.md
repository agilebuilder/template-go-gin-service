# AGENTS.md

Navigation guide for AI coding agents working on <%= projectName %>.

## Project Overview

<%= description %>. An HTTP service built on Go 1.22 + Gin; module path is `<%= modulePath %>`.

## Directory Structure and Responsibilities

| Directory | Responsibility |
| --- | --- |
| `cmd/server` | Program entrypoint: configuration loading and server startup only — no business logic |
| `internal/config` | Configuration structs and loading logic (Viper); add fields to the `Config` struct for new settings |
| `internal/router` | Route registration center; mount new endpoints in `New()` |
| `internal/handler` | HTTP handlers: parse request parameters, call services, return responses; business logic is forbidden here |
| `internal/service` | Business logic layer; must not depend on gin.Context |
| `internal/model` | Request / response structs and domain models |
| `internal/middleware` | Generic middleware (current: logging, panic recovery) |
| `pkg/strutil` | Example of a reusable public utility package; keep `pkg/` free of business-specific code |

## Conventions

- Layering: `handler -> service -> model`; no reverse dependencies; code under `internal/` must not be imported outside the module.
- Every package has a package comment (`// Package xxx ...`); exported identifiers carry doc comments.
- Error handling: return `error` from the bottom layers wrapped with `fmt.Errorf("...: %w", err)`; handlers convert errors into JSON error responses uniformly.
- New configuration: extend the structs in `internal/config/config.go` + the `config.yaml` example; add `v.SetDefault` when necessary.
- New endpoint: add code in the four layers in order — model (structs), service (logic), handler (handler), router (registration) — and add an httptest test for the handler.
- Test naming: `xxx_test.go` in the same directory as the code under test; endpoint tests use `net/http/httptest` + `router.New()` — see `internal/handler/hello_test.go`.

## Common Commands

```bash
make tidy    # download dependencies on first run
make run     # run locally (default port <%= port %>)
make test    # all tests
make vet     # static analysis
make build   # compile into bin/
```

Before committing, make sure `go vet ./...` and `go test ./...` both pass, and format with `gofmt`.
