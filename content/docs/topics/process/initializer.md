---
title: "Initializer"
category: "process"
index: 4
---

# Initializer

### Initializer Definition

An **initializer** is a struct with an `Init` method that takes a [config](/docs/topics/config) object and may return an error value. Initializers should **not** perform long-running computation unless it is necessary for the startup of an application as they will block additional application startup. The following example creates a connection to a Postgres database and stores it in a service container for subsequent initializers and processes to use.

```go
import (
    "github.com/go-nacelle/nacelle"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

type DatabaseInitializer struct {
    Services nacelle.ServiceContainer `service:"services"`
}

type Config struct {
    ConnectionString string `env:"connection_string" required:"true"`
}

func (i *DatabaseInitializer) Init(config nacelle.Config) error {
    dbConfig := &Config{}
    if err := config.Load(dbConfig); err != nil {
        return err
    }

    db, err := sqlx.Open("postgres", dbConfig.ConnectionString)
    if err != nil {
        return err
    }

    return i.Services.Set("db", db)
}
```

An initializer can also be a **finalizer** if it defines a `Finalize` method. This can be useful for initializers that need to do some cleanup action before the application shuts down such as closing a log or profile file, closing remote connections, or ensuring that certain buffers get flushed before the application ends.
