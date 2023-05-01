---
title: "HTTP server"
category: "base"
index: 1
---

## Abstract HTTP server

The [`github.com/go-nacelle/httpbase`](https://github.com/go-nacelle/httpbase) package provides an abstract HTTP server process.

---

This library supplies an abstract HTTP server [process](/docs/topics/process) whose behavior can be configured by implementing a `ServerInitializer` interface. For a more full-featured HTTP server framework built on nacelle, see [chevron](/docs/topics/frameworks/chevron).

An HTTP process is created by supplying an initializer, described below, that controls its behavior.

```go
server := httpbase.NewServer(NewServerInitializer(), options...)
```

A **server initializer** is a struct with an `Init` method that takes a context and an [http.Server](https://golang.org/pkg/net/http/#Server) as parameters.  This method may return an error value, which signals a fatal error to the process that runs it. This method provides an extension point to register handlers to the server instance before the process accepts connections.

The following example registers an HTTP handler function to the server that will handle all incoming requests. Each request atomically increments a request counter on the containing initializer struct and returns its new value. In more complex applications, an HTTP router, such as [gorilla/mux](https://github.com/gorilla/mux) should likely be used.

```go
type Initializer struct {
    requests uint
}

func (i *Initializer) Init(ctx context.Context, server *http.Server) error {
    server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        value := atomic.AddUint32(&i.requests)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(fmt.Sprintf("Hello, #%d!\n", value)))
    })

    return nil
}
```

A simple server initializer stuct that does not need additional methods, state, or dependency instances injected via a service container can use the server initializer function wrapper instead.

```go
_ = ServerInitializerFunc(func(ctx context.Context, server *http.Server) error {
    server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, World!\n"))
    })

    return nil
})
```

You can see an additional example of an HTTP process in the [example repository](https://github.com/go-nacelle/example), specifically the [server initializer](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/http-api/server_initializer.go#L23).

### Related resources

- [Abstract HTTP server environment variable configuration](/docs/ref/envvars_httpbase)
- [Abstract HTTP server functional options](/docs/ref/options_httpbase)
