---
title: "Adapters"
category: "log"
index: 3
---

# Adapters

### Adapters

Nacelle ships with a handful of useful logging adapters. These are extensions of the logger interface that add additional behavior or additional structured data. A custom adapter can be created for behavior that is not provided here.

The **replay adapter** supports journaling log messages and conditionally re-writing them at a different log level. This is useful in circumstances where all the debug logs for a particular request need to be available without making all debug logs in the process available. Messages which are replayed at a higher level will keep the original message timestamp (if supplied), or use the time the log was first published (if not supplied). Each message will also be sent with an additional field called `replayed-from-level` with a value equal to the original level of the message.

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

The **rollup adapter** supports collapsing similar log messages into a multiplicity. This is intended to be used with a chatty subsystem that only logs a handful of messages for which a higher frequency does not provide a benefit (for example, failure to connect to a Redis cache during a network partition). A rollup begins once two messages with the same format string are seen within the rollup window period. During a rollup, all log messages (except for the first in the window) are discarded but counted, and the **first** log message in that window will be sent at the end of the window period with an additional field called `rollup-multiplicity` with a value equal to the number of logs in that window.s

```go
logger := NewRollupAdapter(
    logger,      // base logger
    time.Second, // rollup window
)

for i := 0; i < 10000; i++ {
    logger.Debug("Some problem here!")
}
```
