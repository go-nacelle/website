---
title: "Bootstrapper functional options"
sidebarTitle: "Bootstrapper"
category: "options"
index: 1
---

## Bootstrapper functional options

The [`github.com/go-nacelle/nacelle`](https://github.com/go-nacelle/nacelle) package provides the following functional options to supply the [bootstrapper constructor](https://pkg.go.dev/github.com/go-nacelle/nacelle#NewBootstrapper).

---

The following options can be supplied to the bootstrapper to tune its behavior.

#### WithContextFilter

[WithContextFilter](https://pkg.go.dev/github.com/go-nacelle/nacelle#WithContextFilter) sets a function that accepts a context object and returns a possibly modified context object after the initial health, log, service, and configuration objects are inserted. This allows the user to add additional objects available to all processes prior to initialization.

#### WithConfigSourcer

[WithConfigSourcer](https://pkg.go.dev/github.com/go-nacelle/nacelle#WithConfigSourcer) changes the default source for configuration variables. The default sourcer is the application environment using the name given to the bootstrapper as a prefix.

```go
func main() {
	configSourcer := config.NewFileSourcer("config.toml", config.ParseTOML)

	nacelle.NewBootstrapper(
		"example",
		setup,
		nacelle.WithConfigSourcer(configSourcer),
	).BootAndExit()
}

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error { /* ... */ }
```

#### WithConfigMaskedKeys

[WithConfigMaskedKeys](https://pkg.go.dev/github.com/go-nacelle/nacelle#WithConfigMaskedKeys) sets the keys to mask from log messages when loading configuration data. This is used to hide sensitive configuration values.

```go
func main() {
	nacelle.NewBootstrapper(
		"example",
		setup,
		nacelle.WithConfigMaskedKeys([]string{"GITHUB_ACCESS_TOKEN"}),
	).BootAndExit()
}

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error { /* ... */ }
```

#### WithLoggingInitFunc

[WithLoggingInitFunc](https://pkg.go.dev/github.com/go-nacelle/nacelle#WithLoggingInitFunc) sets the factory used to create the base logger. This can be set to supply a different log backend.

```go
func main() {
	// This matches the default behavior, using go-nacelle/log
	initLogger := func(config *nacelle.Config) (nacelle.Logger, error) {
		c := &log.Config{}
		if err := config.Load(c); err != nil {
			return nil, fmt.Errorf("could not load logging config (%s)", err.Error())
		}

		return log.InitLogger(c)
	}

	nacelle.NewBootstrapper(
		"example",
		setup,
		nacelle.WithLoggingInitFunc(initLogger),
	).BootAndExit()
}

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error { /* ... */ }
```

#### WithLoggingFields

[WithLoggingFields](https://pkg.go.dev/github.com/go-nacelle/nacelle#WithLoggingFields) adds additional fields to every log message. This can be useful to present build information (time, hash, branch), process name, or operating environment.

```go
func main() {
	nacelle.NewBootstrapper(
		"example",
		setup,
		nacelle.WithLoggingFields(nacelle.LogFields{
			"instance_id": uuid.New(),
			"instance_start": time.Now().UTC(),
		}),
	).BootAndExit()
}

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error { /* ... */ }
```
