---
title: "JSON"
category: "config"
index: 1
---

# JSON

#### JSON Unmarshalling

The values that are loaded into non-string field of a configuration struct are interpreted as JSON. Supplying the environment `CASSANDRA_HOSTS='["host1", "host2", "host3"]'` to the configuration struct above will populate the `CassandraHosts` field with the values `host1`, `host2`, and `host3`. The defaults for such fields can also be supplied as a JSON-encoded string, but must be escaped to preserve the format of the struct tag.

```go
type Config struct {
    CassandraHosts []string `env:"cassandra_hosts" default:"[\"host1\", \"host2\", \"host3\"]"`
}
```
