# goboot

`goboot` is a production-grade CLI tool for scaffolding Go projects.

## Phase 1: CLI Foundation

This is Phase 1 of the project, focusing on the CLI structure, flag handling, and validation.

### Features

- CLI foundation using Cobra and Viper.
- Configuration-heavy flag system.
- Clean internal architecture (config, validator, generator, filesystem).
- Input validation for project name, module path, and types.
- Stub generator for configuration summary.

### Installation

```bash
go build -o goboot main.go
```

### Usage

#### Root Command

```bash
./goboot --help
```

#### Version

```bash
./goboot version
```

#### Create a New Project

```bash
./goboot new --name myapp --module github.com/user/myapp --type rest
```

**Flags:**

- `--name`, `-n`: Project name (required)
- `--module`, `-m`: Go module path (required)
- `--type`, `-t`: Project type (`rest` | `grpc` | `cli` | `worker`, default: `rest`)
- `--go-version`: Go version (default: `1.22`)
- `--docker`: Include Dockerfile (default: `false`)
- `--ci`: Include CI/CD workflow (default: `false`)
- `--db`: Database type (`postgres` | `mysql` | `mongo` | `none`, default: `none`)
- `--auth`: Authentication type (`jwt` | `apikey` | `none`, default: `none`)
- `--observability`: Include observability (default: `false`)
- `--dry-run`: Run without creating any files (default: `false`)
- `--output`, `-o`: Output directory (default: `.`)

### Architecture

- `cmd/`: CLI command definitions (Cobra).
- `internal/config/`: Configuration structs and constants.
- `internal/validator/`: Input validation logic.
- `internal/generator/`: Project scaffolding logic (stubbed in Phase 1).
- `internal/filesystem/`: File system abstraction (stubbed in Phase 1).

## Phase 3: gRPC, Config, and Docker

Phase 3 adds support for gRPC microservices, YAML configuration, and Dockerfile generation.

### Features

- gRPC project generation with proto definitions.
- YAML-based configuration for generated services.
- Optional Dockerfile generation via `--docker` flag.
- Refactored generator architecture with `BaseGenerator`.

### Usage

#### Generate a gRPC Service with Docker

```bash
./goboot new --name mygrpc --module github.com/user/mygrpc --type grpc --docker
```

```bash
./goboot new --name myrest --module github.com/user/myrest --type rest
```

## Phase 6: Automation & Ecosystem

Phase 6 expands `goboot` with automation tools, new project types, and remote plugin support.

### Features

- **CI/CD**: Added `goboot add ci` for GitHub Actions workflows.
- **Testing**: Added `goboot add test` for `testify` scaffolds.
- **New Project Types**: Added `cli` and `worker` project types.
- **Remote Plugins**: Added `goboot add remote --url [git-url]` for third-party templates.
- **Interactive Mode**: Added `goboot new --interactive` for a guided setup.

### Usage

#### Create a CLI project interactively

```bash
./goboot new --interactive
```

#### Add CI and Testing to an existing project

```bash
cd myapp
../goboot add ci
../goboot add test
```

#### Add a remote plugin

```bash
../goboot add remote --url https://github.com/user/my-goboot-plugin
```

## Phase 5: Production Readiness & Diagnostics

Phase 5 turns `goboot` into a production-ready platform generator with observability, Kubernetes support, and diagnostic tools.

### Features

- **Observability**: OpenTelemetry, Prometheus, and structured logging.
- **Kubernetes**: Deployment, Service, and ConfigMap manifests.
- **Production Docker**: Multi-stage builds with non-root users.
- **`goboot doctor`**: Project health check and diagnostics.
- **`goboot upgrade`**: Template version tracking and upgrade suggestions.

### Usage

#### Create a new project

```bash
./goboot new --name myapp --module github.com/user/myapp --type grpc --docker
```

#### Add production features

```bash
cd myapp
../goboot add observability
../goboot add k8s
```

#### Run diagnostics

```bash
../goboot doctor
../goboot upgrade
```

## Production Readiness Checklist

- [ ] **Config**: Ensure `config.yaml` is tuned for production.
- [ ] **Observability**: Enable `observability.enabled` in `config.yaml`.
- [ ] **Security**: Review Dockerfile and K8s manifests for security contexts.
- [ ] **Resources**: Adjust K8s resource limits in `k8s/deployment.yaml`.
- [ ] **Secrets**: Move sensitive config to K8s Secrets or a Secret Manager.

## License

MIT

```

```
