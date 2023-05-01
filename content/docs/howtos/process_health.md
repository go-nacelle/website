---
title: "How to track process health"
sidebarTitle: "...track process health"
category: "howtos"
index: 3
noSidebarLink: true
---

## How to track process health

This guide describes a feature of the [`github.com/go-nacelle/process`](https://github.com/go-nacelle/process) package. See [related documentation](/docs/topics/process).

---

Processes can dynamically report their own health via a shared **health tracker** object available in the [service container](/docs/topics/service). The tracker maintains a list of *reasons* that an application is not fully healthy. When this list is empty, the application should be fully functional.

A process that does not interact with the health instance is assumed to be healthy when it is live. Usage of a global health instance should be used as follows.

```go
type HealthConsciousProcess struct {
	healthStatus *nacelle.HealthComponentStatus
}

type healthConsciousProcessKey struct{}

var HealthConsciousProcessKey = healthConsciousProcessKey{}

func (p *HealthConsciousProcess) Init(ctx context.Context) error {
	health := process.HealthFromContext(ctx)

	// Register this process with the health object
	healthStatus, err := health.Register(HealthConsciousProcessKey)
	if err != nil {
		return err
	}
	p.healthStatus = healthStatus
	p.healthStatus.Update(false) // Set unhealthy initially

	// ...
	return nil
}

func (p *HealthConsciousProcess) Run(ctx context.Context) error {
	// pre-healthy processing
	// ...

	// Once healthy, update the health component status
	p.healthStatus.Update(true)

	// post-healthy processing
	// ...

	return nil
}
```
