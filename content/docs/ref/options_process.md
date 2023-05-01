---
title: "Process registration functional options"
sidebarTitle: "Process registration"
category: "options"
index: 3
---

## Process registration functional options

The [`github.com/go-nacelle/process/v2`](https://github.com/go-nacelle/process/v2) package provides the following functional options supplied during process registration.

---

#### WithMetaName

[WithMetaName](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaName) sets the name of the process used in log messages.

#### WithMetaPriority

[WithMetaPriority](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaPriority) sets the relative priority of the process. Lower priorities are initialized earlier (and finalized alter).

#### WithMetadata

[WithMetadata](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetadata) sets additional metadata on the process available in log messages.

#### WithEarlyExit

[WithEarlyExit](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithEarlyExit) marks the process as being able to exit before process shutdown is initiated without it being considered a fatal error. This option is what separates *initializers* from *processes*.

#### WithMetaInitTimeout

[WithMetaInitTimeout](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaInitTimeout) sets the timeout of the `Init` method.

#### WithMetaStartupTimeout

[WithMetaStartupTimeout](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaStartupTimeout) sets the timeout between the invocation of the `Run` method and the process reporting healthy.

#### WithMetaStopTimeout

[WithMetaStopTimeout](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaStopTimeout) sets the timeout of the `Stop` method.

#### WithMetaShutdownTimeout

[WithMetaShutdownTimeout](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaShutdownTimeout) sets the timeout between the initiation of the shutdown sequence and the return of the `Run` method.

#### WithMetaFinalizeTimeout

[WithMetaFinalizeTimeout](https://pkg.go.dev/github.com/go-nacelle/process/v2#WithMetaFinalizeTimeout) sets the timeout of the `Finalize` method.
