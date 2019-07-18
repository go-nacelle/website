+++
title = "AWS Utilities"
category = "libraries"
index = 1
+++

# AWS Utilities

{{% docmeta "awsutil" %}}

<!-- Fold -->

### Usage

This library contains 178 generated nacelle [initializers](https://nacelle.dev/docs/core/process) for AWS services. Each initializer creates an instance of an AWS service `NewDynamoDBServiceInitializer` inside the nacelle [service container](https://nacelle.dev/docs/core/service) with its own configuration (see below).

The following example creates clients for DynamoDB, Kinesis, and S3.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
    processes.RegisterInitializer(awsutil.NewDynamoDBInitializer())
    processes.RegisterInitializer(awsutil.NewKinesisInitializer())
    processes.RegisterInitializer(awsutil.NewS3Initializer())

    // additional setup
    return nil
}
```

### Configuration

The default service behaviors can be configured by the following options. These environment variables must be either prefixed with `AWS` or with the name of a service (e.g. `AWS_REGION` or `S3_REGION`). If any service-specific environment variable is set, then no non-prefix environment variables are read. This can be used for local development with a local AWS service mock such as [localstack](https://github.com/localstack/localstack).

```bash
export KINESIS_ENDPOINT=http://localstack:4568
export DYNAMODB_ENDPOINT=http://localstack:4569
```

| Environment Variable                  | Default | Description |
| ------------------------------------- | ------- | ----------- |
| CREDENTIALS_CHAIN_VERBOSE_ERRORS      | false   | Enables verbose error printing of all credential chain errors. Should be used when wanting to see all errors while attempting to retrieve credentials. |
| DISABLE_COMPUTE_CHECKSUMS             | false   | Disables the computation of request and response checksums, e.g., CRC32 checksums in Amazon DynamoDB. |
| DISABLE_ENDPOINT_HOST_PREFIX          | false   | DisableEndpointHostPrefix will disable the SDK's behavior of prefixing request endpoint hosts with modeled information. |
| DISABLE_PARAM_VALIDATION              | false   | Disables semantic parameter validation, which validates input for missing required fields and/or other semantic request input errors. |
| DISABLE_REST_PROTOCOL_URI_CLEANING    | false   | DisableRestProtocolURICleaning will not clean the URL path when making rest protocol requests. Will default to false. This would only be used for empty directory names in s3 requests. |
| DISABLE_SSL                           | false   | Set this to `true` to disable SSL when sending requests. |
| EC2_METADATA_DISABLE_TIMEOUT_OVERRIDE | false   | Set this to `true` to disable the EC2Metadata client from overriding the default http.Client's Timeout. This is helpful if you do not want the EC2Metadata client to create a new http.Client. This options is only meaningful if you're not already using a custom HTTP client with the SDK. Enabled by default. |
| ENABLE_ENDPOINT_DISCOVERY             | false   | EnableEndpointDiscovery will allow for endpoint discovery on operations that have the definition in its model. By default, endpoint discovery is off. |
| ENDPOINT                              | ""      | An optional endpoint URL (hostname only or fully qualified URI) that overrides the default generated endpoint for a client. |
| ENFORCE_SHOULD_RETRY_CHECK            | false   | EnforceShouldRetryCheck is used in the AfterRetryHandler to always call ShouldRetry regardless of whether or not if request.Retryable is set. This will utilize ShouldRetry method of custom retryers. If EnforceShouldRetryCheck is not set, then ShouldRetry will only be called if request.Retryable is nil. Proper handling of the request.Retryable field is important when setting this field. |
| MAX_RETRIES                           | -1      | The maximum number of times that a request will be retried for failures. Defaults to -1, which defers the max retry setting to the service specific configuration. |
| LOG_LEVEL                             | "off"   | The level at which to log requests. See the note below. |
| REGION                                |         | The region to send requests to. This parameter must be configured globally or on a per-client basis unless otherwise noted. |
| S3_DISABLE_100_CONTINUE               | false   | Set this to `true` to disable the SDK adding the `Expect: 100-Continue` header to PUT requests over 2MB of content. 100-Continue instructs the HTTP client not to send the body until the service responds with a `continue` status. This is useful to prevent sending the request body until after the request is authenticated, and validated. |
| S3_DISABLE_CONTENT_MD5_VALIDATION     | false   | Set this to `true` to disable the S3 service client from automatically adding the ContentMD5 to S3 Object Put and Upload API calls. This option will also disable the SDK from performing object ContentMD5 validation on GetObject API calls. |
| S3_FORCE_PATH_STYLE                   | false   | Set this to `true` to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). |
| S3_USEACCELERATE                      | false   | Set this to `true` to enable S3 Accelerate feature. For all operations compatible with S3 Accelerate will use the accelerate endpoint for requests. Requests not compatible will fall back to normal S3 requests. |
| USE_DUAL_STACK                        | false   | Instructs the endpoint to be generated for a service client to be the dual stack endpoint. The dual stack endpoint will support both IPv4 and IPv6 addressing. |

The available log levels are `off`, `debug`, `debug_with_signing`, `debug_with_http_body`, `debug_with_request_retries`, `debug_with_request_errors`, and `debug_with_event_stream_body`.

For additional documentation on these options, see the [Go AWS SDK](https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config).
