---
title: "Loading"
category: "config"
index: 2
---

# Loading

### Configuration Loading

At initialization time of an application component, the particular subset of configuration variables should be populated, validated, and stored on the service that will later require them.

```go
func (p *Process) Init(config nacelle.Config) error {
    appConfig := &Config{}
    if err := config.Load(appConfig); err != nil {
        return err
    }

    // Use populated appConfig
    return nil
}
```

The `Load` method fails if a value from the source cannot be converted into the correct type, a value from the source cannot be unmarshalled, or is required and not supplied. After each successful load of a configuration struct, the loaded configuration values will are logged. This, however, may be a concern for application secrets. In order to hide sensitive configuration values, add the `mask:"true"` struct tag to the field. This will omit that value from the log message. Additionally, configuration loader object can be initialized with a blacklist of values that should be masked (values printed as `*****` rather than their real value) instead of omitted. These values can be configured in the [bootstrapper](/docs/topics/booting).
