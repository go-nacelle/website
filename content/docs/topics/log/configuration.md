---
title: "Configuration"
category: "log"
index: 1
---

# Configuration

### Configuration

The following environment variables change the behavior of the loggers when using the nacelle [bootstrapper](https://nacelle.dev/docs/core).

| Environment Variable         | Default | Description |
| ---------------------------- | ------- | ----------- |
| LOG_COLORIZE                 | true    | Colorize log messages by level when true. Works with `console` encoding only. |
| LOG_JSON_FIELD_NAMES         |         | A JSON-encoded map to rename the fields for `message`, `timestamp`, and/or `level`. |
| LOG_DISPLAY_FIELDS           | true    | Omit log fields from output when false. Works with `console` encoding only. |
| LOG_DISPLAY_MULTILINE_FIELDS | false   | Print fields on one line when true, one field per line when false. Works with `console` encoding only. |
| LOG_ENCODING                 | console | `console` for human-readable output and `json` for JSON-formatted output. |
| LOG_FIELD_BLACKLIST          |         | A JSON-encoded list of fields to omit from logs. Works with `console` encoding only. |
| LOG_FIELDS                   |         | A JSON-encoded map of fields to include in every log. |
| LOG_LEVEL                    | info    | The highest level that will be emitted. |
| LOG_SHORT_TIME               | false   | Omit date from timestamp when true. Works with `console` encoding only. |

Example log output is shown below. The outputs are configure in order with: the default configuration, `LOG_ENCODING=json`, `LOG_DISPLAY_FIELDS=false`, `LOG_SHORT_TIME=true`, and `LOG_DISPLAY_MULTILINE_FIELDS=true`.

```
[I] [2019/07/24 09:15:30.806] Accepted request from 68.6.165.7 caller=derision/main.go:20 requestId=12341234-1234-1234-1234-123412341234 sequenceNumber=2

{"caller":"derision/main.go:20","level":"info","message":"Accepted request from 68.6.165.7","requestId":"12341234-1234-1234-1234-123412341234","sequenceNumber":2,"timestamp":"2019-07-24T09:16:55.673-0700"}

[I] [2019/07/24 09:15:49.517] Accepted request from 68.6.165.7

[I] [09:17:56] Accepted request from 68.6.165.7 caller=derision/main.go:20 requestId=12341234-1234-1234-1234-123412341234 sequenceNumber=2

[I] [2019/07/24 09:16:38.117] Accepted request from 68.6.165.7
    caller = derision/main.go:20
    requestId = 12341234-1234-1234-1234-123412341234
    sequenceNumber = 2
```
