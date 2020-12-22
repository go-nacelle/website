---
title: "Writing your first Nacelle app, part 3"
category: "intro"
index: 5
---

#### Process Configuration

The application above creates a process with hard-coded port of 5000. This is problematic in the case you need to change the port when running on a different environment, or run two servers on the same host.

We can instead accept this value from the environment (environment variable, file, configmap, etc) at runtime so that no code change is required to configure this value.

We declare the configuration values accepted by the server process with a configuration struct. Here, we tag the port field with `env` which indicates the environment variable that should be read to populate this field.

In the `Init` method of the process, we populate an instance of this struct with values and pull the required values inot the process struct for later use.

```go
type serverConfig struct {
    Port int `env:"port" default:"5000"`
}

func (s *server) Init(config nacelle.Config) error {
    serverConfig := &serverConfig{}
    if err := config.Load(serverConfig); err != nil {
        return err
    }

    s.port = serverConfig.Port
    // ...
}

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    processes.RegisterProcess(&server{}, nacelle.WithProcessName("hw-server"))
    return nil
}
```

Running the application the same way will show the same behavior. Running the application with `PORT=3000` will cause the application to listen to the non-default port 3000. Running the application with a non-integer port value will cause the application to fail on startup with an error message.

