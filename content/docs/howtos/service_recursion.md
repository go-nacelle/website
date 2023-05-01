---
title: "How to inject services recursively into struct fields"
sidebarTitle: "...inject services recursively into struct fields"
category: "howtos"
index: 7
noSidebarLink: true
---

## How to inject services recursively into struct fields

This guide describes a feature of the [`github.com/go-nacelle/service`](https://github.com/go-nacelle/service) package. See [related documentation](/docs/topics/service).

---

It should be noted that injection does not work **recursively**. The procedure does not look into the values of non-anonymous fields. If this behavior is needed, it can be performed during a post-injection hook. The following example demonstrates this behavior.

```go
type RecursiveInjectionConsumer struct {
	Services *service.ServiceContainer `service:"services"`
	FieldA   *A
	FieldB   *B
	FieldC   *C
}

func (c *RecursiveInjectionConsumer) PostInject() error {
	fields := []interface{}{
		c.FieldA,
		c.FieldB,
		c.FieldC,
	}

	for _, field := range fields {
		if err := c.Services.Inject(field); err != nil {
			return err
		}
	}

	return nil
}
```
