+++
title = "Lambda"
category = "base processes"
index = 3
+++

# Base AWS Lambda Server

{{% docmeta "lambdabase" %}}

<!-- Fold -->

### Usage

The supplied server process is an abstract AWS Lambda RPC server whose behavior is determined by a supplied `Handler` interface which wraps the handler defined by [aws-lambda-go](https://github.com/aws/aws-lambda-go/blob/af0b813d5803d9754b920ed666b1cf8c16becfb3/lambda/handler.go#L14). There is an [example](./example) included in this repository.

Supplied are six constructors supplied that create servers that specifically requests from an [event source mapping](https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html). These servers handle the request unmarshalling and add additional log context.

- **NewDynamoDBEventServer** invokes the backing handler with a list of DynamoDBEventRecords.
- **NewDynamoDBRecordServer** invokes the backing handler once for each DynamoDBEventRecord in the batch.
- **NewKinesisEventServer** invokes the backing handler with a list of KinesisEventRecords.
- **NewKinesisRecordServer** invokes the backing handler once for each KinesisEventRecord in the batch.
- **NewSQSEventServer** invokes the backing handler with a list of SQSMessages.
- **NewSQSRecordServer** invokes the backing handler once for each SQSMessage in the batch.

### Configuration

The default process behavior can be configured by the following environment variables.

| Environment Variable | Required | Description |
| -------------------- | -------- | ----------- |
| _LAMBDA_SERVER_PORT  | yes      | The port on which to listen for RPC commands. |
