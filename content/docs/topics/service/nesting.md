---
title: "Nesting"
category: "service"
index: 5
---

# Nesting

#### Anonymous Structs

Injection also works on structs containing composite fields. The following example successfully assigns the registered value to the field `Child.Base.Service`.

```go
type Base struct {
    Service *SomeExample `service:"example"`
}

type Child struct {
    *Base
}

child := &Child{}
if err := container.Inject(child); err != nil {
    // handle error
}
```
