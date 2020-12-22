---
title: "Retrieval"
category: "service"
index: 1
---

# Retrieval


### Retrieval

A service can be retrieved from the service container by the name with which it is registered. However, this returns a bare interface object and requires the consumer of the service to do a type-check and cast.

Instead, the recommended way to consume dependent services is to **inject** them into a struct decorated with tagged fields. This does the proper type conversion for you.

```go
type Consumer struct {
    Service *SomeExample `service:"example"`
}

consumer := &Consumer{}
if err := services.Inject(consumer); err != nil {
    // handle error
}
```

The `Inject` method fails when a consumer asks for an unregistered service or for a service with the wrong target type. Services can be tagged as optional (e.g. `service:"example" optional:"true"`) which will silence the later class of errors. Tagged fields must be exported.
