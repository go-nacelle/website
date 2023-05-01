---
title: "How to adjust the caller frame in log messages"
sidebarTitle: "...adjust the caller frame in log messages"
category: "howtos"
index: 2
noSidebarLink: true
---

## How to adjust the caller frame in log messages

This guide describes a feature of the [`github.com/go-nacelle/log`](https://github.com/go-nacelle/log) package. See [related documentation](/docs/topics/log).

---

Sometimes it is useful to define helper functions that logs messages for you in a common way (applying common context, performing additional behavior on errors, etc). Unfortunately, this can interfere with the way the caller file and line number are discovered. The logger has a `WithIndirectCaller` method that will increase the depth used when scanning the stack for callers. This should be used at each log location that aggregates log calls (e.g. any place where knowing this source location would not be helpful).

```go
func logForMe(message string) {
	parentLogger.WithIndirectCaller(1).Log(message)
}

logForMe("foobar")
```

Using a logger instance directly does not require this additional hint.
