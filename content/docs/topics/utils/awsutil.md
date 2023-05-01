---
title: "AWS"
category: "utils"
index: 2
---

## AWS utilities

The [`github.com/go-nacelle/awsutil`](https://github.com/go-nacelle/awsutil) package provides AWS service utilities.

---

This library contains 178 generated nacelle [initializers](/docs/topics/process) for AWS services. Each initializer creates an instance of an AWS service `NewDynamoDBServiceInitializer` inside the nacelle [service container](/docs/topics/service) with its own configuration (see below).

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

### Related resources

- [AWS environment variable configuration](/docs/ref/envvars_awsutil)
