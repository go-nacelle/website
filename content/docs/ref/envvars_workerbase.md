---
title: "Abstract background worker environment variable configuration"
sidebarTitle: "Worker"
category: "envvars"
index: 4
---

## Abstract background worker environment variable configuration

Environment variables available to the Nacelle process change the behavior of the worker.

---

{{< table table_class="table table-striped table-bordered" >}}
| Environment Variable | Default | Description |
| -------------------- | ------- | ----------- |
| WORKER_STRICT_CLOCK  | false   | Subtract the duration of the previous tick from the time between calls to the spec's tick function. |
| WORKER_TICK_INTERVAL | 0       | The time (in seconds) between calls to the spec's tick function. |
{{< /table >}}
