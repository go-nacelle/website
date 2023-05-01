---
title: "Abstract gRPC server functional options"
sidebarTitle: "Abstract gRPC server"
category: "options"
index: 5
---

## Abstract gRPC server functional options

The [`github.com/go-nacelle/grpcbase`](https://github.com/go-nacelle/grpcbase) package provides the following functional options to supply the [server constructor](https://pkg.go.dev/github.com/go-nacelle/grpcbase#NewServer).

---

#### WithTagModifiers

[WithTagModifiers](https://pkg.go.dev/github.com/go-nacelle/grpcbase#WithTagModifiers) registers the tag modifiers to be used when loading the server's [configuration](https://pkg.go.dev/github.com/go-nacelle/grpcbase#Config). This can be used to change default hosts and ports, or prefix all target environment variables in the case where more than one gRPC server is registered per application (e.g. health server and application server, data plane and control plane server).

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	dataServerConfigFunc := // ...
	metaServerConfigFunc := // ...

	// Reads data_grpc_{host,port}
	dataServer := grpcbase.NewServer(dataServerConfigFunc, grpcbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("data")))
	processes.RegisterProcess(dataServer)

	// Reads meta_grpc_{host,port}
	metaServer := grpcbase.NewServer(metaServerConfigFunc, grpcbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("meta")))
	processes.RegisterProcess(metaServer)
}
```

#### WithServerOptions

[WithServerOptions](https://pkg.go.dev/github.com/go-nacelle/grpcbase#WithServerOptions) registers options to be supplied directly to the gRPC server constructor.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	grpcServer := grpcbase.NewServer(dataServerConfigFunc, grpcbase.WithServerOptions(grpc.ServerOptions{
		RequireClientCert: true,
	}))
    
	processes.RegisterProcess(grpcServer)
}
```
