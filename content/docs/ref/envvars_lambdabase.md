---
title: "AWS Lambda worker environment variable configuration"
sidebarTitle: "AWS Lambda worker"
category: "envvars"
index: 5
---

## AWS Lambda worker environment variable configuration

Environment variables available to the Nacelle process change the behavior of the AWS Lambda worker.

---

{{< table table_class="table table-striped table-bordered" >}}
| Environment Variable | Required | Description |
| -------------------- | -------- | ----------- |
| _LAMBDA_SERVER_PORT  | yes      | The port on which to listen for RPC commands. |
{{< /table >}}
