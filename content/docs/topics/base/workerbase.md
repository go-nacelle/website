---
title: "Background worker"
category: "base"
index: 4
---

## Abstract background worker process

The [`github.com/go-nacelle/workerbase`](https://github.com/go-nacelle/workerbase) package provides an abstract background worker process.

---

A **worker** is a process that periodically polls an external resource in order to discover or perform its main unit of work. A **base worker** is an abstract [process](/docs/topics/process) whose behavior can be be configured by implementing the `WorkerSpec` interface.

A worker process is created by supplying a specification, described below, that controls its behavior.

```go
worker := workerbase.NewWorker(NewWorkerSpec(), options...)
```

A **worker specification** is a struct with an `Init` and a `Tick` method. Both methods take a context as a parameter. On process shutdown, this context object is cancelled so that any long-running work in the tick method can be cleanly abandoned. Each method may return an error value, which signals a fatal error to the process that runs it.

The following example uses a database connection injected by the service container, and pings it to logs its latency. The worker process will call the tick method in a loop based on its interval configuration while the process remains active.

```go
type Spec struct {
    DB     *sqlx.DB       `service:"db"`
    Logger nacelle.Logger `service:"logger"`
}

func (s *Spec) Init(ctx context.Context) error {
    return nil
}

func (s *Spec) Tick(ctx context.Context) error {
    start := time.Now()
    err := s.DB.Ping()
    duration := time.Now().Sub(start)
    durationMs := float64(duration) / float64(time.Milliseconds)

    if err != nil {
        return err
    }

    s.Logger.Debug("Ping took %.2fms", durationMs)
}
```

If the worker specification also implements the `Finalize` method, it will be called after the last invocation of the tick method (regardless of its return value).

```go
func (s *Spec) Finalize() error {
    s.Logger.Debug("Shutting down")
    return nil
}
```

This library comes with an [example](https://github.com/go-nacelle/workerbase/tree/main/example) project. You can see an additional example of a worker process in the [example repository](https://github.com/go-nacelle/example), specifically the [worker spec](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/worker_spec.go#L15) definition.

### Related resources

- [Abstract background worker environment variable configuration](/docs/ref/envvars_workerbase)
- [Abstract background worker functional options](/docs/ref/options_workerbase)
