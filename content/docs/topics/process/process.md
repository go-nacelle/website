---
title: "Process"
category: "process"
index: 6
---

# Process

### Process Definition

A **process** is a struct with an `Init`, a `Start`, and a `Stop` method. The initialization method that takes a [config](https://nacelle.dev/docs/core/config) object as a parameter. Each method may return an error value. For long-running processes, such as servers, the start method should be blocking. The stop method may signal the process to gracefully shut-down (via a channel or synchronization primitive), but does not need to wait until the application exits. A process is also an initializer, so the above also applies. The following example uses the database connection created by the initializer defined [below](https://nacelle.dev/docs/core/process#initializer-definition), injected by the service container, and pings it on a loop to logs its latency. The stop method closes a channel to inform the start method to unblock.

```go
type PingProcess struct {
    DB           *sqlx.DB       `service:"db"`
    Logger       nacelle.Logger `service:"logger"`
    halt         chan struct{}
    once         sync.Once
    tickInterval time.Duration
}

type Config struct {
    TickInterval int `env:"tick_interval"`
}

func (p *PingProcess) Init(config nacelle.Config) error {
    pingConfig := &Config{}
    if err := config.Load(pingConfig); err != nil {
        return err
    }

    p.tickInterval = time.Duration(pingConfig.tickInterval) * time.Second
    return nil
}

func (p *PingProcess) Start() error {
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

func (p *PingProcess) Stop() error {
    p.once.Do(func() {
        close(p.halt)
    })

    return nil
}
```

