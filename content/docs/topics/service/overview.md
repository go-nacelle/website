---
title: "Service"
defines: "service"
category: "topics"
index: 4
---

# Service

Service container and dependency injection for [nacelle](https://nacelle.dev).

---

A **service container** is a collection of objects which are constructed separately from their consumers. This pattern allows for a greater separation of concerns, where consumers care only about a particular concrete or interface type, but do not care about their configuration, construction, or initialization. This separation also allows multiple consumers for the same service which does not need to be initialized multiple times (e.g. a database connection or an in-memory cache layer). Service injection is performed on [initializers and processes](https://nacelle.dev/docs/core/process) during application startup automatically.

You can see an additional example of service injection in the [example repository](https://github.com/go-nacelle/example), specifically the [worker spec](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/worker_spec.go#L15) definition. In this project, the `Conn` and `PubSubConn` services are created by application-defined initializers [here](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L28) and [here](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/pubsub_initializer.go#L32).
