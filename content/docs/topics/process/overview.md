---
title: "Process"
defines: "process"
category: "topics"
index: 5
---

# Process

Process initializer and supervisor for [nacelle](https://nacelle.dev).

---

Nacelle applications can be coarsely decomposed into several behavioral categories.

<dl>
  <dt>Process</dt>
  <dd>A <a href="https://pkg.go.dev/github.com/go-nacelle/process#Process">process</a> is a long-running component of an application such as a server or an event processor. Processes will usually run for the life of the application (e.g. until a shutdown signal is received or an error occurs). There is a special class of processes that are allowed to exit once running, but these are the exception.</dd>
</dl>

<dl>
  <dt>Initializer</dt>
  <dd>An <a href="https://pkg.go.dev/github.com/go-nacelle/process#Initializer">initializer</a> is a component that is invoked once on application startup. An initializer usually instantiates a service or set up shared state required by other parts of the application</dd>
</dl>

<dl>
  <dt>Service</dt>
  <dd>A <a href="/docs/topics/service">service</a> is an object that encapsulates some data, state, or behavior, but does not have a rigid initialization. A service is generally instantiated by an initializer and inserted into a shared service container.</dd>
</dl>

You can see additional examples of initializer and process definition and registration in the [example repository](https://github.com/go-nacelle/example). Specifically, there is an [initializer](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L80) to create a shared Redis connection and its [registration](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/cmd/worker/main.go#L9) in one of the program entrypoints. This project also provides a set of abstract base processes for common process types: an [AWS Lambda event listener](/docs/topics/base/lambdabase), a [gRPC server](/docs/topics/base/grpcbase), an [HTTP server](/docs/topics/base/httpbase), and a [generic worker process](/docs/topics/base/workerbase), which are a good and up-to-date source for best-practices.

- [health](/docs/topics/process/health)
- [initializer](/docs/topics/process/initializer)
- [lifecycle](/docs/topics/process/lifecycle)
- [parallel](/docs/topics/process/parallel)
- [process](/docs/topics/process/process)
- [registration](/docs/topics/process/registration)
