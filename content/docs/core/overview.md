+++
title = "Core"
+++

# Core Functionality

{{% docmeta "nacelle" %}}

<!-- Fold -->

Nacelle provides a bootstrapper in order to control the population of [configuration](https://nacelle.dev/docs/core/config) structs, the injection of [dependencies](https://nacelle.dev/docs/core/service), the initialization and supervision of [processes](https://nacelle.dev/docs/core/process), and the initialization of [logging](https://nacelle.dev/docs/core/log).

Applications written in with nacelle should have a common entrypoint, as follows.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    // register initializer and process instances
}

func main() {
    nacelle.NewBootstrapper("app-name", setup).BootAndExit()
}
```

The following options can be supplied to the bootstrapper to tune its behavior.

- **WithConfigSourcer** changes the default source for configuration variables. The default sourcer is the application environment using the name given to the bootstrapper as a prefix.
- **WithConfigMaskedKeys** sets the keys to mask from log messages when loading configuration data. This is used to hide sensitive configuration values.
- **WithLoggingInitFunc** sets the factory used to create the base logger. This can be set to supply a different log backend.
- **WithLoggingFields** adds additional fields to every log message. This can be useful to present build information (time, hash, branch), process name, or operating environment.
- **WithRunnerOptions** sets configuration for the process runner. See the [process](https://nacelle.dev/docs/core/process) library for additional details.
