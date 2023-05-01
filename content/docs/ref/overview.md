---
title: "References"
defines: "ref"
category: docs
---

## References

Available interface implementations, available functional options, and available envvar configuration options, etc.

---

### Implementations

Implementations of interfaces provided by default.

- [Config sourcers](/docs/ref/implementations_config_sourcers)
- [Config tag modifiers](/docs/ref/implementations_config_tag_modifiers)
- [Lambda event sources](/docs/ref/implementations_lambdabase_event_sources)

### Functional options

Functional constructor arguments provided by each package.

- [Boostrapper](/docs/ref/options_booting)
- [Initializer registration](/docs/ref/options_process_initializer)
- [Process registration](/docs/ref/options_process)
- [Abstract HTTP server](/docs/ref/options_httpbase)
- [Abstract gRPC server](/docs/ref/options_grpcbase)
- [Abstract background worker](/docs/ref/options_workerbase)

### Environment variables

Environment variables supplied to a Nacelle process to dynamically change behavior of the framework.

- [Logging](/docs/ref/envvars_log)
- [HTTP server](/docs/ref/envvars_httpbase)
- [gRPC server](/docs/ref/envvars_grpcbase)
- [Background worker](/docs/ref/envvars_workerbase)
- [AWS Lambda worker](/docs/ref/envvars_lambdabase)
- [Postgres](/docs/ref/envvars_pgutil)
- [AWS](/docs/ref/envvars_awsutil)

### API documentation

Documentation on [pkg.go.dev](https://pkg.go.dev/github.com/go-nacelle/nacelle).

- Core functionality
  - [github.com/go-nacelle/config](https://pkg.go.dev/github.com/go-nacelle/config)
  - [github.com/go-nacelle/log](https://pkg.go.dev/github.com/go-nacelle/log)
  - [github.com/go-nacelle/process](https://pkg.go.dev/github.com/go-nacelle/process)
  - [github.com/go-nacelle/service](https://pkg.go.dev/github.com/go-nacelle/service)
- Frameworks
  - [github.com/go-nacelle/chevron](https://pkg.go.dev/github.com/go-nacelle/chevron)
  - [github.com/go-nacelle/scarf](https://pkg.go.dev/github.com/go-nacelle/scarf)
- Base processes
  - [github.com/go-nacelle/httpbase](https://pkg.go.dev/github.com/go-nacelle/httpbase)
  - [github.com/go-nacelle/grpcbase](https://pkg.go.dev/github.com/go-nacelle/grpcbase)
  - [github.com/go-nacelle/workerbase](https://pkg.go.dev/github.com/go-nacelle/workerbase)
  - [github.com/go-nacelle/lambdabase](https://pkg.go.dev/github.com/go-nacelle/lambdabase)
- Utility libraries
  - [github.com/go-nacelle/pgutil](https://pkg.go.dev/github.com/go-nacelle/pgutil)
  - [github.com/go-nacelle/awsutil](https://pkg.go.dev/github.com/go-nacelle/awsutil)
