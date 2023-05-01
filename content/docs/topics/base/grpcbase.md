---
title: "gRPC server"
category: "base"
index: 2
---

## Abstract gRPC server

The [`github.com/go-nacelle/grpcbase`](https://github.com/go-nacelle/grpcbase) package provides an abstract gRPC server process.

---

This library supplies an abstract gRPC server [process](/docs/topics/process) whose behavior can be configured by implementing a `ServerInitializer` interface. For a more full-featured gRPC server framework built on nacelle, see [scarf](/docs/topics/frameworks/scarf).

A gRPC process is created by supplying an initializer, described below, that controls its behavior.

```go
server := grpcbase.NewServer(NewServerInitializer(), options...)
```

A **server initializer** is a struct with an `Init` method that takes a context and an [grpc.Server](https://pkg.go.dev/google.golang.org/grpc#Server) as parameters.  This method may return an error value, which signals a fatal error to the process that runs it. This method provides an extension point to register services to the server instance before the process accepts connections.

The following example registers a gRPC service to the server that will atomically increment a request counter and return it in a payload defined in the `proto` package that also contains the service definition.

```go
type Initializer struct {}

func (i *Initializer) Init(ctx context.Context, server *http.Server) error {
    proto.RegisterRequestCounterServiceServer(server, &RequestCounterService{})
    return nil
}

type RequestCounterService {
    requests uint
}

func (kvs *RequestCounterService) Get(ctx context.Context, r *proto.Request) (*proto.Response, error) {
    value := atomic.AddUint32(&i.requests)
    return &proto.Response{count: value}, nil
}
```

A simple server initializer stuct that does not need additional methods, state, or dependency instances injected via a service container can use the server initializer function wrapper instead.

```go
_ = ServerInitializerFunc(func(ctx context.Context, server *grpc.Server) error {
    proto.RegisterRequestCounterServiceServer(server, &RequestCounterService{})
    return nil
})
```

You can see an additional example of a gRPC process in the [example repository](https://github.com/go-nacelle/example), specifically the [server initializer](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/grpc-api/server_initializer.go#L21).

### Related resources

- [Abstract gRPC server environment variable configuration](/docs/ref/envvars_grpcbase)
- [Abstract gRPC server functional options](/docs/ref/options_grpcbase)
