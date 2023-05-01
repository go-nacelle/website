---
title: "Abstract background worker functional options"
sidebarTitle: "Abstract background worker"
sidebarTitle: "Abstract worker"
category: "options"
index: 6
---

## Abstract background worker functional options

The [`github.com/go-nacelle/workerbase`](https://github.com/go-nacelle/workerbase) package provides the following functional options to supply the [worker constructor](https://pkg.go.dev/github.com/go-nacelle/workerbase#NewWorker).

---

#### WithTagModifiers
  
[WithTagModifiers](https://pkg.go.dev/github.com/go-nacelle/workerbase#WithTagModifiers) registers the tag modifiers to be used when loading the worker's [configuration](https://pkg.go.dev/github.com/go-nacelle/workerbase#Config). This can be used to change the default tick interval, or prefix all target environment variables in the case where more than one worker process is registered per application.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	kinesisWorkerSpec := // ...
	dynamoDBWorkerSpec := // ...

	// Reads kinesis_worker_{strict_clock,tick_interval}
	kinesisWorker := NewWorker(kinesisWorkerSpec, workerbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("kinesis")))
    processes.RegisterProcess(kinesisWorker)

	// Reads dynamodb_worker_{strict_clock,tick_interval}
	dynamoDBWorker := NewWorker(dynamoDBWorkerSpec, workerbase.WithTagModifiers(nacelle.NewEnvTagPrefixer("dynamodb")))
    processes.RegisterProcess(dynamoDBWorker)
}
```
