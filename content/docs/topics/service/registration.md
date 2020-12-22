---
title: "Registration"
category: "service"
index: 2
---

# Registration

### Registration

A concrete service can be registered to the service container with a unique name by which it can later be retrieved. The `Set` method fails when a service name is reused. There is also an analogous `MustSet` method that panics on error.

```go
func Init(services nacelle.ServiceContainer) error {
    example := &Example{}
    if err := services.Set("example", example); err != nil {
        return err
    }

    // ...
}
```

The [logger](/docs/topics/log) (under the name `logger`), the [health tracker](/docs/topics/process/health) (under the name `health`), and the service container itself (under the name `services`) are available in all applications using the nacelle [bootstrapper](/docs/topics/booting).
