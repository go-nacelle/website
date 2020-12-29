---
title: "Writing your first Nacelle app, part 1"
sidebarTitle: "Tutorial (part 1)"
category: "intro"
index: 3
---

# Writing your first Nacelle app, part 1

### Fifteen-Minute Walkthrough

The core ideas of Nacelle revolve around *processes* and *initializers*. A process is a behavioral component of your application which does some work over the process lifetime. An initializer is a component of your application which does some work at application startup.

An application can be composed of one or more processes, which are commonly long-running such as a server or a worker that accepts work from an external system. An application may also have zero or more initializers, which generally create or initialize a resource (such as a database connection) used by an application process.

#### Setup

Applications using Nacelle to bootstrap will have the following minimal `main` function. This hands control off to the bootrapper, which will invoke the registered `setup` function in order to populate the process and service containers. The bootrapper will then initialize each process and monitor it for the lifetime of the application.

```go
package main

import "github.com/go-nacelle/nacelle"

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    // Register processes and initializers here
    return nil
}

func main() {
    nacelle.NewBootstrapper("app-name", setup).BootAndExit()
}
```

If you were to run this application, you would see Nacelle trying to initialize each registered initializer (of which there are none), and initialize and start each registered process (of which there are none).
