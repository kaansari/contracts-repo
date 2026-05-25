# Ceerat Contracts

Shared protobuf contracts, generated Go code, domain DTOs, and mappers for Ceerat services.

## Contains

```text
proto/auth/       Auth service protobuf definitions and generated Go code
proto/customer/   Customer protobuf definitions and generated Go code
proto/order/      Order protobuf definitions and generated Go code
proto/service/    Service protobuf definitions and generated Go code
domain/           Pure shared business objects
mapper/           Conversion helpers between protobuf messages and domain objects
```

## gRPC Services

```text
auth.Auth
customer.CustomerService
order.OrderService
service.ServiceManager
```

## Important Boundary

This module should stay free of service-specific persistence concerns:

- no GORM tags
- no database models
- no repository interfaces
- no service implementation logic

## Regenerate Protobuf Code

Install the protobuf plugins, then run from this directory:

```bash
make proto
```

The `proto` target uses source-relative paths so generated Go files stay under their source directories.

The `go_package` options point to:

```text
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/customer
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/order
github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/service
```

## Test

```bash
go test ./...
```

## Make Commands

```bash
make help
make tidy
make test
make proto
make push GITHUB_USER=kaansari
make tag VERSION=v0.1.0
```
