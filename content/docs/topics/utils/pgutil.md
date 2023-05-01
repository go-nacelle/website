---
title: "Postgres"
category: "utils"
index: 1
---

## Postgres utilities

The [`github.com/go-nacelle/pgutil`](https://github.com/go-nacelle/pgutil) package provides Postgres utilities.

---

This library creates a [sqlx](https://github.com/jmoiron/sqlx) connection wrapped in a nacelle [logger](/docs/topics/log). The supplied initializer adds this connection into the nacelle [service container](/docs/topics/service) under the key `db`. The initializer will block until a ping succeeds.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    processes.RegisterInitializer(pgutil.NewInitializer())

    // additional setup
    return nil
}
```

This library uses [golang migrate](https://github.com/golang-migrate/migrate) to optionally run migrations on application startup. To configure migrations, supply a [source driver](https://github.com/golang-migrate/migrate#migration-sources) to the initializer, as follows.

```go
import (
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/golang-migrate/migrate/v4/source"
)

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    migrationSourceDriver, err := source.Open("file:///migrations")
	if err != nil {
		return err
	}

    processes.RegisterInitializer(pgutil.NewInitializer(
        pgutil.WithMigrationSourceDriver(migrationSourceDriver)
    ))

    // ...
}
```

### Related resources

- [Postgres environment variable configuration](/docs/ref/envvars_pgutil)
