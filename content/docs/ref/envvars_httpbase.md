---
title: "Abstract HTTP server environment variable configuration"
sidebarTitle: "HTTP server"
category: "envvars"
index: 2
---

## Abstract HTTP server environment variable configuration

Environment variables available to the Nacelle process change the behavior of the HTTP server.

---

{{< table table_class="table table-striped table-bordered" >}}
| Environment Variable  | Default | Description |
| --------------------- | ------- | ----------- |
| HTTP_HOST             | 0.0.0.0 | The host on which to accept connections. |
| HTTP_PORT             | 5000    | The port on which to accept connections. |
| HTTP_CERT_FILE        |         | The path to the TLS cert file. |
| HTTP_KEY_FILE         |         | The path to the TLS key file. |
| HTTP_SHUTDOWN_TIMEOUT | 5       | The time (in seconds) the server can spend in a graceful shutdown. |
{{< /table >}}

##### Notes

The environment variables `HTTP_CERT_FILE` and `HTTP_KEY_FILE` are mutual and must both be set (if set). Setting these will start a TLS server.
