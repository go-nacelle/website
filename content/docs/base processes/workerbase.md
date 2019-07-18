+++
title = "Worker"
category = "base processes"
index = 4
+++

# Base Worker Process

{{% docmeta "workerbase" %}}

<!-- Fold -->

### Usage

The supplied process is an abstract busy-loop whose behavior is determined by a supplied `WorkerSpec` interface. This interface has an `Init` method that receives application config and a `Tick` method where each phase of work should be done. The tick method is passed a context that will be canceled on shutdown so that any long-running work can be cleanly abandoned. There is an [example](./example) included in this repository.

- **WithTagModifiers** registers the tag modifiers to be used when loading process configuration (see [below](#Configuration)). This can be used to change the default tick interval, or prefix all target environment variables in the case where more than one worker process is registered per application.

### Configuration

The default process behavior can be configured by the following environment variables.

| Environment Variable | Default | Description |
| -------------------- | ------- | ----------- |
| WORKER_STRICT_CLOCK  | false   | Subtract the duration of the previous tick from the time between calls to the spec's tick function. |
| WORKER_TICK_INTERVAL | 0       | The time (in seconds) between calls to the spec's tick function. |
