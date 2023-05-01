---
title: "At a glance"
category: "intro"
index: 1
---

## At a glance

What is Nacelle?

---

Nacelle is a Golang framework for setting up and monitoring the behavioral components of your application. It is intended to provide a common and convention-guided way to bootstrap an application or set of microsevices. Broadly, boostrapping your application with Nacelle gives you:

- a convention for code organization via processes and initializers
- a convention for declaring, reading, and validating configuration values from the runtime environment
- a convention for registering, declaring, and injecting structure interface dependencies
- a convention for structured logging

Concretely, bootstrapping your application with Nacelle gives you:

- an expanding set of libraries that use the conventions outlined above
- a process initializer that reads and validates declared component configuration
- a process initializer that injects declared component dependencies
- a process runner that invokes each process in a dedicated goroutine
- a process monitor that watches for error and cleanly shutdown your application
