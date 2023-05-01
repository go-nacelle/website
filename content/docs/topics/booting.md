---
title: "Application setup"
defines: "booting"
category: "topics"
index: 1
---

## Booting

The [`github.com/go-nacelle/nacelle`](https://github.com/go-nacelle/nacelle) package provides a common [bootstrapper](https://pkg.go.dev/github.com/go-nacelle/nacelle/v2#Bootstrapper) object that initializes and supervises the core framework behaviors.

---

Applications written with nacelle should have a common entrypoint, as follows. The application-specific functionality is passed to a boostrapper on construction as a reference to a function that populates a [process container](/docs/topics/process). The `BootAndExit` function initializes and supervises the application, blocks until the application shut down, then calls `os.Exit` with the appropriate status code. A symmetric function called `Boot` will perform the same behavior, but will return the integer status code instead of calling `os.Exit`.

```go
func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	// register initializer and process instances
}

func main() {
	nacelle.NewBootstrapper("app-name", setup).BootAndExit()
}
```

You can see additional examples of the bootstrapper in the [example repository](https://github.com/go-nacelle/example). Specifically, the main function of the [HTTP API](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/http-api/main.go#L17), the [gRPC API](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/grpc-api/main.go#L16), and the [worker](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/main.go#L17) processes.

### Related resources

- [Bootstrapper functional options](/docs/ref/options_booting)
