---
title: "Writing your first app"
defines: "tutorial"
category: "intro"
index: 3
---

## Writing your first Nacelle app

Learn Nacelle by example.

---

This tutorial builds a toy application that exercises the core features of all Nacelle applications including process management, configuration loading and validation, and runtime dependency management. The resulting application will be HTTP server that returns the number of requests made to the application. The server itself is stateless as the request count is stored in a shareable Redis instance.

To follow this tutorial, you will need to:

- Create an empty Go project with the following module dependencies:
  - `github.com/go-nacelle/nacelle/v2` (see [installation docs](/docs/intro/install))
  - `github.com/go-redis/redis/v7`
- [Install Redis](https://redis.io/docs/getting-started/installation/), or make a Redis instance otherwise available to your local machine.

Now, without further ado...

- [Part 1: Setting up the application](/docs/intro/tutorial-01)
- [Part 2: Setting up a skeleton "Hello, World!" server](/docs/intro/tutorial-02)
- [Part 3: Adding runtime configuration](/docs/intro/tutorial-03)
- [Part 4: Injecting shared dependencies](/docs/intro/tutorial-04)
