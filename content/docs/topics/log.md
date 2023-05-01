---
title: "Logging"
defines: "log"
category: "topics"
index: 3
---

## Logging

The [`github.com/go-nacelle/log`](https://github.com/go-nacelle/log) package provides opinionated structured logger.

---

Nacelle loggers emit **structured logs** to **standard error**. It is absolutely essential to be able to correlate and aggregate log messages to form a view of a running system. In order to redirect logs to a secondary target (such as an ELK stack), the application's output should simply be redirected. This keeps the application simple and allows redirection of logs to **any** source without requiring an application update. For an example of redirection when run in a Docker container, see nacelle's [fluentd wrapper](https://github.com/go-nacelle/fluentd).

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

### Replay adapter

The [replay logger](https://pkg.go.dev/github.com/go-nacelle/log/v2#ReplayLogger) is an extension of the logger interface that supports journaling log messages and conditionally re-writing them at a different log level. This is useful in circumstances where all the debug logs for a particular request need to be available without making all debug logs in the process available. Messages which are replayed at a higher level will keep the original message timestamp (if supplied), or use the time the log was first published (if not supplied). Each message will also be sent with an additional field called `replayed-from-level` with a value equal to the original level of the message.

```go
requestLogger := NewReplayAdapter(
	logger,         // base logger
	log.LevelDebug, // track debug messages for replay
	log.LevelInfo,  // also track info messages
)

// handle request

if requestTookTooLong() {
	// Re-log journaled messages at warning level
	requestLogger.Replay(log.LevelWarning)
}
```

### Rollup adapter

The [rollup logger](https://pkg.go.dev/github.com/go-nacelle/log/v2#NewRollupAdapter) is an extension of the logger interface that supports collapsing similar log messages into a multiplicity. This is intended to be used with a chatty subsystem that only logs a handful of messages for which a higher frequency does not provide a benefit (for example, failure to connect to a Redis cache during a network partition). A rollup begins once two messages with the same format string are seen within the rollup window period. During a rollup, all log messages (except for the first in the window) are discarded but counted, and the **first** log message in that window will be sent at the end of the window period with an additional field called `rollup-multiplicity` with a value equal to the number of logs in that window.

```go
logger := NewRollupAdapter(
	logger,      // base logger
	time.Second, // rollup window
)

for i := 0; i < 10000; i++ {
	logger.Debug("Some problem here!")
}
```

### Related resources

- [Logging environment variable configuration](/docs/ref/envvars_log)
- [How to adjust the caller frame in log messages](/docs/howtos/log_adjust_caller)
