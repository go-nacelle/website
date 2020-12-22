---
title: "Registration"
category: "process"
index: 1
---

# Registration

### Registration

The following example registers a cache and a database initializer, then registers a health server, an HTTP server, and a gRPC server. The database initializer must complete within five seconds or the application fails (and, if orchestrated properly, will restart). The health server is registered at a lower priority than the other two. This ensures that the health server can take requests as early as the HTTP and gRPC servers, and also ensures that it will be shut down only after the higher-priority processes have been shut down.

```go
func setupProcesses(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    processes.RegisterInitializer(
        NewCacheInitializer(),
        nacelle.WithInitializerName("cache"),
    )

    processes.RegisterInitializer(
        NewDatabaseInitializer(),
        nacelle.WithInitializerName("db"),
        nacelle.WithInitializerTimeout(time.Second * 5),
    )

    processes.RegisterProcess(
        NewHealthServer(),
        nacelle.WithProcessName("health-server"),
    )

    processes.RegisterProcess(
        NewHTTPServer(),
        nacelle.WithProcessName("http-server"),
        nacelle.WithPriority(1),
    )

    processes.RegisterProcess(
        NewGRPCServer(),
        nacelle.WithProcessName("grpc-server"),
        nacelle.WithPriority(1),
    )

    return nil
}
```

#### Process Registration Options

The following options can be set on a process during registration.

<dl>
    <dt>WithPriority</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithPriority">WithPriority</a> sets the priority group the process belongs to. Processes of the same - priority are initialized and started in parallel.</dd>

  <dt>WithProcessInitTimeout</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithProcessInitTimeout">WithProcessInitTimeout</a> sets the maximum time that the process can spend in its Init method.</dd>

  <dt>WithProcessName</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithProcessName">WithProcessName</a> sets the name of the process in log messages.</dd>

  <dt>WithProcessShutdownTimeout</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithProcessShutdownTimeout">WithProcessShutdownTimeout</a> sets the maximum time that the process can spend waiting for its Start method to unblock after its Stop method is called.</dd>

  <dt>WithProcessStartTimeout</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithProcessStartTimeout">WithProcessStartTimeout</a> sets the maximum time that the process can spend <strong>unhealthy</strong> after its Start method is called. See <a href="#health">health</a> below.</dd>

  <dt>WithSilentExit</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithSilentExit">WithSilentExit</a> sets a flag that allows a nil error value to be returned without signaling an application shutdown. This can be useful for things like leader election on startup which should not stop hot standby processes from taking client requests.</dd>
</dl>

#### Initializer Registration Options

The following options can be set on an initializer during registration.

<dl>
  <dt>WithFinalizerTimeout</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithFinalizerTimeout">WithFinalizerTimeout</a> sets the maximum time that the initializer can spend on finalization. As the application is already shutting down, this will simply log and error and unblock the finalizer.</dd>

  <dt>WithInitializerName</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithInitializerName">WithInitializerName</a> sets the name of the initializer in log messages.</dd>

  <dt>WithInitializerTimeout</dt>
  <dd><a href="https://godoc.org/github.com/go-nacelle/process#WithInitializerTimeout">WithInitializerTimeout</a> sets the maximum time that the initializer can spend in its Init method. An error will be returned if this time is exceeded.</dd>
</dl>
