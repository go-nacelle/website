---
title: "Process management"
defines: "process"
category: "topics"
index: 5
---

## Process management

The [`github.com/go-nacelle/process`](https://github.com/go-nacelle/process) package provides process initializer and supervisor.

---

Nacelle applications can be coarsely decomposed into several behavioral categories.

<dl>
  <dt>Process</dt>
  <dd>A <a href="https://pkg.go.dev/github.com/go-nacelle/process/v2#Process">process</a> is a long-running component of an application such as a server or an event processor. Processes will usually run for the life of the application (e.g. until a shutdown signal is received or an error occurs). There is a special class of processes that are allowed to exit once running, but these are the exception.</dd>
</dl>

<dl>
  <dt>Initializer</dt>
  <dd>An <a href="https://pkg.go.dev/github.com/go-nacelle/process/v2#Initializer">initializer</a> is a special type of process that is not expected to last for the life of the application. An initializer usually instantiates a service or set up shared state required by other parts of the application</dd>
</dl>

<dl>
  <dt>Service</dt>
  <dd>A <a href="/docs/topics/service">service</a> is an object that encapsulates some data, state, or behavior, but does not have a rigid initialization. A service is generally instantiated by an initialized process and inserted into a shared service container.</dd>
</dl>

Concrete instances of processes are registered to a **process container builder** at application startup (as in the example shown below). The nacelle [bootstrapper](/docs/topics/booting) then handles initialization and invocation of the process container, service container, and runner that will control the initialization and supervision of registered components. For each priority to which a process is registered, that priority group is initialized, invoked, and supervised by the following sequence of events:

First, each process is initialized. First, [services](/docs/topics/service) are injected into the process instance via the shared service container. Each process in a group is initialized concurrently, but if a process from a lower priority had already registered a service into the container, it will be available at this time. After injection, process's `Init` method is invoked. If an error occurs in either stage, the remainder of the boot process is abandoned. Any process that had successfully completed is **unwound**: each process that implements a `Finalize` method will have it be invoked. Processes in the reverse order of their initialization.

Next, the `Run` method of each process that defines it is invoked. Each invocation is made concurrently and in a different goroutine. The remainder of the boot process is suspended until all processes within this priority become [healthy](#tracking-process-health). If a process returns from its `Run` method or does not become healthy within the given timeout period, the boot process is abandoned and the process is unwound, as described in the next stage.

Lastly, once all priority groups have been started and have become healthy after initialization, the supervisory stage begins. This stage listens for one of the following events and begins to unwind the process.

  - The user sends the process a signal
  - A process's Start method returns with an error
  - A process's Start method returns without an error, but is not marked for silent exit

The process is unwound by stopping each process for which a Start goroutine was created. The Stop method of each process is called concurrently with all processes within its priority batch, and each priority batch is stopped by (descending) priority order. Finally, processes are unwound as described above.

### Defining initializers

An **initializer** is a struct with an `Init` method that takes a context and may return an error value. Generally, initializers are setting up some shared state in the application, and will likely want to have access to a [Config](/docs/topics/config) object. This example achieves this by adding a struct-tag field populated by the boot process.

Initializers should **not** perform long-running computation unless it is necessary for the startup of an application as they will block additional application startup. The following example creates a connection to a Postgres database and stores it in a service container for subsequent initializers and processes to use.

```go
import (
    "github.com/go-nacelle/nacelle"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

type DatabaseInitializer struct {
    Config   *nacelle.Config           `service:"config"`
    Services *nacelle.ServiceContainer `service:"services"`
}

type Config struct {
    ConnectionString string `env:"connection_string" required:"true"`
}

func (i *DatabaseInitializer) Init(ctx context.Context) error {
    dbConfig := &Config{}
    if err := i.Config.Load(dbConfig); err != nil {
        return err
    }

    db, err := sqlx.Open("postgres", dbConfig.ConnectionString)
    if err != nil {
        return err
    }

    return i.Services.Set("db", db)
}
```

An initializer can also be a **finalizer** if it defines a `Finalize` method. This can be useful for initializers that need to do some cleanup action before the application shuts down such as closing a log or profile file, closing remote connections, or ensuring that certain buffers get flushed before the application ends.

### Defining processes

A **process** is a struct with an `Init`, a `Run`, and a `Stop` method. Each method takes a context object as a parameter and may return an error value. For long-running processes, such as servers, the `Run` method should be blocking. The `Stop` method may signal the process to gracefully shut-down (via a channel or synchronization primitive), but does not need to wait until the application exits. A process is also an initializer, so the above details also apply. The following example uses the database connection created by the initializer defined in the section above, injected by the service container, and pings it on a loop to logs its latency. The stop method closes a channel to inform the start method to unblock.

```go
type PingProcess struct {
	Config       *nacelle.Config `service:"config"`
	Logger       nacelle.Logger  `service:"logger"`
	DB           *sqlx.DB        `service:"db"`
	halt         chan struct{}
	once         sync.Once
	tickInterval time.Duration
}

type Config struct {
	TickInterval int `env:"tick_interval"`
}

func (p *PingProcess) Init(ctx context.Context) error {
	pingConfig := &Config{}
	if err := p.Config.Load(pingConfig); err != nil {
		return err
	}

	p.tickInterval = time.Duration(pingConfig.tickInterval) * time.Second
	return nil
}

func (p *PingProcess) Run(ctx context.Context) error {
	for {
		select {
			case <-p.halt:
				return nil
			case <-time.After(p.tickInterval):
		}

		start := time.Now()
		err := p.DB.Ping()
		duration := time.Now().Sub(start)
		durationMs := float64(duration) / float64(time.Milliseconds)

		if err != nil {
			return err
		}

		p.Logger.Debug("Ping took %.2fms", durationMs)
	}

	return nil
}

func (p *PingProcess) Stop(ctx context.Context) error {
	p.once.Do(func() {
		close(p.halt)
	})

	return nil
}
```

You can see additional examples of initializer and process definition and registration in the [example repository](https://github.com/go-nacelle/example). Specifically, there is an [initializer](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L80) to create a shared Redis connection and its [registration](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/main.go#L9) in one of the program entrypoints. This project also provides a set of abstract base processes for common process types: an [HTTP server](/docs/topics/base/httpbase), a [gRPC server](/docs/topics/base/grpcbase), a [generic worker process](/docs/topics/base/workerbase), and an [AWS Lambda event listener](/docs/topics/base/lambdabase) which are a good and up-to-date source for best-practices.

### Related resources

- [Process registration functional options](/docs/ref/options_process)
- [How to track process health](/docs/howtos/process_health)
