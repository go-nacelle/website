---
title: "Struct"
category: "config"
index: 5
---

# Struct

### Configuration Struct Definition

Configuration structs are defined by the application or library developer with the fields needed by the package in which they are defined. Each field is tagged with a *source hint* (e.g. an environment variable name, a key in a YAML file) and, optionally, default values and basic validation. Tagged fields must be exported in order for this package to assign to them.

The following example defines configuration for a hypothetical worker process. For the application to start successfully, the address of an API must be supplied. All other configuration values are optional.

```go
type Config struct {
    APIAddr        string   `env:"api_addr" required:"true"`
    CassandraHosts []string `env:"cassandra_hosts"`
    NumWorkers     int      `env:"num_workers" default:"10"`
    BufferSize     int      `env:"buffer_size" default:"1024"`
}
```
