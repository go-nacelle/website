+++
title = "Core"
+++

# Core Functionality

{{% docmeta "nacelle" %}}

<!-- Fold -->

For more details, see [the website](https://nacelle.dev), the [getting started gide](https://nacelle.dev/getting-started), and [the motivating blog post](https://eric-fritz.com/articles/nacelle/).

## Goals

Core goals:

- Provide a common convention for application organization so that developers can quickly dive into the meaningful logic of an application.
- Support a common convention for declaring, reading, and validating configuration values from the runtime environment.
- Support a common convention for registering, declaring, and injecting struct and interface dependencies.
- Support a common convention for structured logging.

Additional goals:

- Provide additional non-core functionality via separate opt-in libraries. Keep the dependencies for the core-functionality minimal.
- Operate within existing infrastructures and do not require tools or technologies outside of what this project provides.

## Non-goals

- Impose opinions on service discovery.
- Impose opinions on inter-process or inter-service communication.
- Impose opinions on runtime environment, deployment, orchestration.services.
