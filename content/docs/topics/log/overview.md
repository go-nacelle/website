---
title: "Log"
defines: "log"
category: "topics"
index: 3
---

# Log

Opinionated structured logger for [nacelle](https://nacelle.dev).

---

Nacelle loggers emit **structured** logs to **standard error**. It is absolutely essential to be able to correlate and aggregate log messages to form a view of a running system. In order to redirect logs to a secondary target (such as an ELK stack), the application's output should simply be redirected. This keeps the application simple and allows redirection of logs to **any** source without requiring an application update. For an example of redirection when run in a Docker container, see nacelle's [fluentd wrapper](https://github.com/go-nacelle/fluentd).

There are five standard log levels: `Debug`, `Info`, `Warning`, `Error`, and `Fatal`. Logging at the fatal level will abort the application after flushing any outstanding log messages. The logger interface has a method for each log level with printf-like arguments (a format string and a variable number of arguments used to construct the message).

```go
logger.Error("Failed to dial database (%s)", err.Error())
```

In addition, the logger interface has a `WithFields` variant, which takes a map of additional log data as a first argument. A `nacelle.LogFields` value is a map from strings to interface types and can be used interchangeably.

```go
logger.DebugWithFields(nacelle.LogFields{
    "requestId": "00001111-2222-3333-4444-555566667777",
}, "Accepted request from %s", remoteAddr)
```

A logger can also be decorated with a set of fields so that multiple calls to the logger share the same set of base fields. This is useful for message correlation in servers where a logger instance can be given a unique request or client identifier. Creating a decorated logger does not modify the base logger, thus it is safe to create multiple concurrent decorated loggers from the same logger instance without worrying about interference.

```go
requestLogger := logger.WithFields(nacelle.LogFields{
    "requestId": "00001111-2222-3333-4444-555566667777",
})

requestLogger.Info("Accepted request from %s", remoteAddr)
```



