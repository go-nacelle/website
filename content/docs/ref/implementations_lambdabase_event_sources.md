---
title: "Lambda event source implementations"
sidebarTitle: "Lambda event sources"
category: "implementations"
index: 3
---

## Lambda event source implementations

The [`github.com/go-nacelle/lambdabase`](https://github.com/go-nacelle/lambdabase) package provides first-class support for the following [event sources](https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html).

---

These servers require a more specific handler interface invoked with unmarshalled request data and additional log context.

#### DynamoDB event server

[NewDynamoDBEventServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewDynamoDBEventServer) invokes the backing handler with a list of [DynamoDBEventRecords](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#DynamoDBEventRecord).</dd>

#### DynamoDB record server

[NewDynamoDBRecordServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewDynamoDBRecordServer) invokes the backing handler once for each [DynamoDBEventRecord](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#DynamoDBEventRecord) in the batch.</dd>

#### Kinesis event server

[NewKinesisEventServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewKinesisEventServer) invokes the backing handler with a list of [KinesisEventRecords](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#KinesisEventRecord).</dd>

#### Kinesis record server

[NewKinesisRecordServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewKinesisRecordServer) invokes the backing handler once for each [KinesisEventRecord](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#KinesisEventRecord) in the batch.</dd>

#### SQS event server

[NewSQSEventServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewSQSEventServer) invokes the backing handler with a list of [SQSMessages](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#SQSMessage).</dd>

#### SQS record server

[NewSQSRecordServer](https://pkg.go.dev/github.com/go-nacelle/lambdabase#NewSQSRecordServer) invokes the backing handler once for each [SQSMessage](https://pkg.go.dev/github.com/aws/aws-lambda-go/events#SQSMessage) in the batch.</dd>
