# goboot 🚀

`goboot` is a production-grade CLI tool designed to scaffold Go projects with best practices, clean architecture, and a pluggable microservice factory.

[![Go Version](https://img.shields.io/github/go-mod/go-version/codewithme224/goboot)](https://github.com/codewithme224/goboot)
[![Release](https://img.shields.io/github/v/release/codewithme224/goboot)](https://github.com/codewithme224/goboot/releases)
[![License](https://img.shields.io/github/license/codewithme224/goboot)](LICENSE)

## Features

- **Pluggable Architecture**: Decoupled feature generators for databases, auth, and more.
- **Incremental Adoption**: Add features to existing projects using `goboot add`.
- **Production Ready**: Multi-stage Docker builds, Kubernetes manifests, and OpenTelemetry.
- **Diagnostics**: Built-in `doctor` command to check project health.
- **Interactive Mode**: Guided setup for new projects.

---

## Installation

### From Source (Requires Go 1.22+)

```bash
go install github.com/codewithme224/goboot@latest
```

### From Binaries

Download the latest binary for your platform from the [Releases](https://github.com/codewithme224/goboot/releases) page.

---

## Quick Start

### 1. Create a New Project

The easiest way to start is using the **Interactive Mode**:

```bash
goboot new --interactive
```

Or use flags for a quick setup:

```bash
goboot new --name myapp --module github.com/user/myapp --type rest --docker
```

### 2. Add Features Incrementally

Navigate to your project directory and add what you need:

```bash
cd myapp
goboot add db --type postgres
goboot add observability
goboot add k8s
```

---

## Core Commands

### `goboot new`

Scaffold a new project from scratch.

- **Types**: `rest`, `grpc`, `cli`, `worker`.
- **Flags**: `--docker`, `--db`, `--auth`, `--ci`, `--observability`.

### `goboot add`

Enhance an existing project with new capabilities.

- `add db`: Supports `postgres`, `mysql`, `mongo`.
- `add auth`: Adds JWT middleware (REST).
- `add gateway`: Adds gRPC-Gateway (gRPC).
- `add observability`: Adds OTel + Prometheus + Structured Logging.
- `add k8s`: Generates Deployment, Service, and ConfigMap.
- `add ci`: Generates GitHub Actions workflows.
- `add test`: Adds `testify` scaffolds.
- `add remote`: Pulls templates from a remote Git URL.

### `goboot doctor`

Checks the health of your project, validates `config.yaml`, `go.mod`, and checks for production best practices.

### `goboot upgrade`

Checks if your project templates are out of date and suggests a migration plan.

---

## Project Structure (Generated)

A typical `goboot` project follows a clean architecture:

```text
myapp/
├── cmd/                # Entry points (api, grpc, cli, worker)
├── internal/
│   ├── config/         # Configuration loading
│   ├── server/         # Server setup (HTTP/gRPC)
│   ├── service/        # Business logic
│   ├── db/             # Database connections (if added)
│   └── observability/  # OTel & Metrics (if added)
├── k8s/                # Kubernetes manifests (if added)
├── config.yaml         # Application configuration
├── Dockerfile          # Multi-stage production build
└── go.mod              # Module definition
```

---

## Production Readiness Checklist

- [ ] **Observability**: Enable `observability.enabled: true` in `config.yaml`.
- [ ] **Security**: Review Dockerfile and K8s security contexts.
- [ ] **Resources**: Tune CPU/Memory limits in `k8s/deployment.yaml`.
- [ ] **Secrets**: Move sensitive data from `ConfigMap` to K8s `Secrets`.

---

## License

MIT © [codewithme224](https://github.com/codewithme224)
