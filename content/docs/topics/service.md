---
title: "Dependency injection"
defines: "service"
category: "topics"
index: 4
---

## Dependency injection

The [`github.com/go-nacelle/service`](https://github.com/go-nacelle/service) package provides service container and dependency injection.

---

A **service container** is a collection of objects which are constructed separately from their consumers. This pattern allows for a greater separation of concerns, where consumers care only about a particular concrete or interface type, but do not care about their configuration, construction, or initialization. This separation also allows multiple consumers for the same service which does not need to be initialized multiple times (e.g. a database connection or an in-memory cache layer). Service injection is performed on [processes](/docs/topics/process) during application startup automatically.

A concrete service can be registered to the service container with a unique name by which it can later be retrieved. The `Set` method fails when a service name is reused. There is also an analogous `MustSet` method that panics on error. The [logger](/docs/topics/log) (under the name `logger`), the [health tracker](/docs/howtos/process_health) (under the name `health`), and the service container itself (under the name `services`) are available in all applications using the nacelle [bootstrapper](/docs/topics/booting).

```go
type Process struct {
	Services *nacelle.ServiceContainer `service:"services"`
}

func (p *Process) Init(ctx context.Context) error {
	example := &Example{}
	if err := p.Services.Set("example", example); err != nil {
		return err
	}

	// ...
}
```

A service can be retrieved from the service container by the name with which it is registered. An alternative way to consume dependent services is to **inject** them into a struct decorated with tagged fields.

```go
type Consumer struct {
    Service *SomeExample `service:"example"`
}

consumer := &Consumer{}
if err := services.Inject(consumer); err != nil {
    // handle error
}
```

The `Inject` method fails when a consumer asks for an unregistered service or for a service with the wrong target type. Services can be tagged as optional (e.g. `service:"example" optional:"true"`) which will silence the later class of errors. Tagged fields must be exported.

You can see an additional example of service injection in the [example repository](https://github.com/go-nacelle/example), specifically the [worker spec](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/worker_spec.go#L15) definition. In this project, the `Conn` and `PubSubConn` services are created by application-defined initializers [here](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L28) and [here](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/pubsub_initializer.go#L32).

### Related resources

- [How to run custom code after service injection](/docs/howtos/service_post_injection_hook)
- [How to inject services recursively into struct fields](/docs/howtos/service_recursion)
- [How to inject dependencies into composite fields](/docs/howtos/service_composite)
