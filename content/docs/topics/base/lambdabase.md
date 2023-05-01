---
title: "AWS Lambda worker"
category: "base"
index: 3
---

## Abstract AWS Lambda worker process

The [`github.com/go-nacelle/lambdabase`](https://github.com/go-nacelle/lambdabase) package provides an abstract AWS Lambda worker process.

---

This library supplies an abstract AWS Lambda RPC server [process](/docs/topics/process) whose behavior can be be configured by implementing a `Handler` interface. This interface wraps the handler defined by [aws-lambda-go](https://github.com/aws/aws-lambda-go/blob/af0b813d5803d9754b920ed666b1cf8c16becfb3/lambda/handler.go#L14).

A Lambda server process is created by supplying a handler, described below, that controls its behavior.

```go
server := lambdabase.NewServer(NewHandler(), options...)
```

A **handler** is a struct with an `Init` and a `Handle` method. The initialization method, like the process that runs it, that takes a context as a parameter adn returns an error value. The handle method of the base server takes a context object and the request payload as parameters and returns the response payload and an error value. The handle method of an event-specific server takes a context object, the request payload, and a logger populated with request and event identifiers as parameters and returns an error value. Return an error from either method signals a fatal error to the process that runs it.

The following example implements a handler that transforms a request value to upper-case.

```go
type Handler struct {}

type ReqResp struct {
    Value string `json:"text"`
}

func (h *Handler) Init(ctx context.Context) error {
    return nil
}

func (h *Handler) (ctx context.Context, payload []byte) ([]byte, error) {
    request := &ReqResp{}
    if err := json.Unmarshal(payload, &request); err != nil {
        return nil, err
    }

    return json.Marshal(&ReqResp{Value: strings.ToUpper(request.Value)})
}
```

This library comes with an [example](https://github.com/go-nacelle/lambdabase/tree/main/example) project that logs values received from a Kinesis stream.

### Related resources

- [Lambda event source implementations](/docs/ref/implementations_lambdabase_event_sources)
- [AWS Lambda worker environment variable configuration](/docs/ref/envvars_lambdabase)
