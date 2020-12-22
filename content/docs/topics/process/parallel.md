---
title: "Parallel"
category: "process"
index: 2
---

# Parallel





#### Parallel Initializers

Initializers run sequentially and non-concurrently so that a previously registered initializer can provide a service required by a subsequently registered initializer. However, this default sequencing behavior is not always necessary and can increase startup time.

When initializers can be run independently (for example, creating multiple SDK client instances for a list of AWS services), it is unnecessary to run them in sequence. Groups of such services can be registered to a `ParallelInitializer`. This is a bag of initializers that run the initializers registered to it in parallel. The initializer will block its siblings until all of its children complete successfully.

```go
awsGroup := nacelle.NewParallelInitializer()
awsGroup.RegisterInitializer(NewDynamoDBClientInitializer())
awsGroup.RegisterInitializer(NewKinesisClientInitializer())
awsGroup.RegisterInitializer(NewLambdaClientInitializer())
awsGroup.RegisterInitializer(NewSQSClientInitializer())

// Register parallel group
processes.RegisterInitializer(awsGroup)
```
