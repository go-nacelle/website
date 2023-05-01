---
title: "How to run custom code after service injection"
sidebarTitle: "...run custom code after service injection"
category: "howtos"
index: 6
noSidebarLink: true
---

## How to run custom code after service injection

This guide describes a feature of the [`github.com/go-nacelle/service`](https://github.com/go-nacelle/service) package. See [related documentation](/docs/topics/service).

---

If additional behavior is necessary after services are available to a consumer struct (e.g. running injection on the elements of a dynamic slice or map or cross-linking dependencies), the method `PostInject` can be implemented. This method, if it is defined, is invoked immediately after successful injection.

```go
func (c *Consumer) PostInject() error {
	return c.Service.PrepFor("consumer")
}
```
