---
title: "Health"
category: "process"
index: 5
---

# Health

#### Tracking Process Health

Processes can dynamically report their own health via a shared **health tracker** object available in the [service container](https://nacelle.dev/docs/core/service). The tracker maintains a list of *reasons* that an application is not fully healthy. When this list is empty, the application should be fully functional.

A process that does not interact with the health instance is assumed to be healthy when it is live. Usage of a global health instance should be used as follows.

```go
type HealthConsciousProcess struct {
    Health nacelle.Health `service:"health"`
}

func (p *HealthConsciousProcess) Init(config nacelle.Config) error {
    // Before the process starts, add an "unhealthy" token unique to this process
    if err := p.Health.AddReason("health-conscious-process"); err != nil {
        return err
    }

    // ...
    return nil
}

func (p *HealthConsciousProcess) Start() error {
    // pre-healthy processing
    // ...

    // Once healthy, remove the reason registered above
    if err := p.Health.RemoveReason("health-conscious-process"); err != nil {
        return err
    }

    // post-healthy processing
    // ...

    return nil
}

func (p *HealthConsciousProcess) Stop() error {
    // ...
    return nil
}
```