---
title: "Lifecycle"
category: "process"
index: 3
---

# Lifecycle


### Application Lifecycle

Concrete instances of initializers and processes are registered to a **process container** at application startup (as in the example shown below). The nacelle [bootstrapper](https://nacelle.dev/docs/core) then handles initialization and invocation of the process container and runner that will control the initialization and supervision of registered components. The following sequence of events occur on application boot.

- **Initializer boot stage:** Each initializer is booted sequentially in the order that it is registered. First, [services](https://nacelle.dev/docs/core/service) are injected into the initializer instance via the shared service container. If a previous initializer had registered a service into the container, it will be available at this time. Next, the initializer's Init method is invoked. If an error occurs in either stage, the remainder of the boot process is abandoned. Any initializer that had successfully completed is **unwound**: each initializer that implements the finalizer interface will be have their finalization method invoked. Initializers are finalized sequentially in the reverse order of their initialization.

- **Process boot stage:** After all initializers have completed successfully, the processes will begin to boot. First, services are injected into **all** processes, regardless of their priority order. If this fails, then the remainder of the boot process is abandoned and the initializers are unwound. Then, processes continue to boot in **batches** based on their (ascending) priority order. Within each batch, the following sequence of events occur.

  - **Batch initialization:** The Init method of each process is invoked sequentially in the order that it is registered. If this fails, then the remainder of the boot process is abandoned the initializers are unwound.

  - **Batch launch:** The Start method of each process is invoked. Each invocation is made concurrently and in a different goroutine. The remainder of the boot process is suspended until all processes within this priority become [healthy](#tracking-process-health). If a process returns from its Start method or does not become healthy within the given timeout period, the boot process is abandoned and the process is unwound, as described in the next stage.

- **Supervisory stage:** Once all process batches have been started and have become healthy after initialization, the supervisory stage begins. This stage listens for one of the following events and begins to unwind the process.

  - The user sends the process a signal
  - A process's Start method returns with an error
  - A process's Start method returns without an error, but is not marked for silent exit

The process is unwound by stopping each process for which a Start goroutine was created. The Stop method of each process is called concurrently with all processes within its priority batch, and each priority batch is stopped by (descending) priority order. Finally, initializers are unwound as described above.
