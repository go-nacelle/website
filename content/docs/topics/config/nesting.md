---
title: "Nesting"
category: "config"
index: 3
---

# Nesting

#### Anonymous Structs

Loading configuration values also works with structs containing composite fields. The following example shows the definition of multiple configuration structs with a set of shared fields.

```go
type StreamConfig struct {
    StreamName string `env:"stream_name" required:"true"`
}

type StreamProducerConfig struct {
    StreamConfig
    PublishAttempts int `env:"publish_attempts" default:"3"`
    PublishDelay    int `env:"publish_delay" default:"1"`
}

type StreamConsumerConfig struct {
    StreamConfig
    FetchLimit int `env:"fetch_limit" default:"100"`
}
```
