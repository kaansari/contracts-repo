# Ceerat Contracts

Shared contracts for Ceerat services.

## Contains

```text
proto/auth/       Auth service protobuf definitions and generated Go code
proto/patient/    Patient protobuf definitions and generated Go code
domain/           Pure shared business objects
mapper/           Conversion helpers between protobuf messages and domain objects
```

## Important boundary

This module should stay free of service-specific persistence concerns:

- no GORM tags
- no database models
- no repository interfaces
- no service implementation logic

## Test

```bash
go test ./...
```

## Regenerate protobuf code

Install the protobuf plugins, then run from this directory:

```bash
make proto
```

The `proto` target uses source-relative paths so generated Go files stay under `proto/auth` and `proto/patient`.

The `go_package` options point to:

```text
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/patient
```
## Make commands

```bash
make help
make tidy
make test
make push GITHUB_USER=kaansari
make tag VERSION=v0.1.0
```
