---
title: "Abstract HTTP server functional options"
sidebarTitle: "Abstract HTTP server"
category: "options"
index: 4
---

## Abstract HTTP server functional options

The [`github.com/go-nacelle/httpbase`](https://github.com/go-nacelle/httpbase) package provides the following functional options to supply the [server constructor](https://pkg.go.dev/github.com/go-nacelle/httpbase#NewServer).

---

#### WithTagModifiers

[WithTagModifiers](https://pkg.go.dev/github.com/go-nacelle/httpbase#WithTagModifiers) registers the tag modifiers to be used when loading the server's [configuration](https://pkg.go.dev/github.com/go-nacelle/httpbase#Config). This can be used to change default hosts and ports, or prefix all target environment variables in the case where more than one HTTP server is registered per application (e.g. health server and application server, data plane and control plane server).

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	dataServerConfigFunc := // ...
	metaServerConfigFunc := // ...

	// Reads data_http_{host,port,cert_file,key_file,shutdown_timeout}
	dataServer := httpbase.NewServer(dataServerConfigFunc, httpbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("data")))
	processes.RegisterProcess(dataServer)

	// Reads meta_http_{host,port,cert_file,key_file,shutdown_timeout}
	metaServer := httpbase.NewServer(metaServerConfigFunc, httpbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("meta")))
	processes.RegisterProcess(metaServer)
}
```
