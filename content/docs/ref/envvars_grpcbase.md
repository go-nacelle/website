---
title: "Abstract gRPC server environment variable configuration"
sidebarTitle: "gRPC server"
category: "envvars"
index: 3
---

## Abstract gRPC server environment variable configuration

Environment variables available to the Nacelle process change the behavior of the gRPC server.

---

{{< table table_class="table table-striped table-bordered" >}}
| Environment Variable | Default | Description |
| -------------------- | ------- | ----------- |
| GRPC_HOST            | 0.0.0.0 | The host on which to accept connections. |
| GRPC_PORT            | 5000    | The port on which to accept connections. |
{{< /table >}}
