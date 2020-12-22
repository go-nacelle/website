---
title: "Post injection"
category: "service"
index: 4
---

# Post injection

#### Post Injection Hook

If additional behavior is necessary after services are available to a consumer struct (e.g. running injection on the elements of a dynamic slice or map or cross-linking dependencies), the method `PostInject` can be implemented. This method, if it is defined, is invoked immediately after successful injection.

```go
func (c *Consumer) PostInject() error {
    return c.Service.PrepFor("consumer")
}
```
