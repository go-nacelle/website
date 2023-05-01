---
title: "How to inject dependencies into composite fields"
sidebarTitle: "...inject dependencies into composite fields"
category: "howtos"
index: 5
noSidebarLink: true
---

## How to inject dependencies into composite fields

This guide describes a feature of the [`github.com/go-nacelle/service`](https://github.com/go-nacelle/service) package. See [related documentation](/docs/topics/service).

---

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
