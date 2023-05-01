---
title: "Part 4: Injecting shared dependencies"
sidebarTitle: "Part 4: Dependencies"
category: "tutorial"
index: 4
---

## Injecting shared dependencies

This section of the tutorial enhances the application to persist state in a shared Redis instance.

---

Let's modify the server to return a distinct response for each request. Instead of a canned message, we will print their request count: *Hello #1* for the first request, *Hello #2!* for the second, and so on. We'll store this data in Redis, and atomically increment a request counter each time the handler is invoked.

This creates a dependency for a Redis client in the server process.

```go
import (
	// ...
	"github.com/go-redis/redis/v7"
)

type server struct {
	client *redis.Client
	// ...
}
```

We'll change the HTTP handler implementation as follows.

```go
s.server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	count, err := s.client.Incr("sample").Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, #%d!\n", count)))
})
```

Finally, we need to supply a concrete Redis client to the process on startup. We'll do this *for now* by initializing the client in the setup function and passing it to the process on creation.

```go
func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	processes.RegisterProcess(&server{client: client}, nacelle.WithMetaName("hw-server"))
	return nil
}
```

The application should now produce HTTP responses with an increasing count in the body. If Redis is not running or accessible on your host, then the server should respond with an internal server error response.

Unfortunately, the last change above creates a few issues, namely:

1. We do not check if the client can reach a remote server.
2. We do not pull the Redis address from the environment. We hard-code the address, which has the same issues as hard-coding the port above. Additionally, the bootstrapper hasn't created the configuration object yet, so we *couldn't* read from the environment at this point in the application lifecycle anyway.
3. We are supplying dependencies manually. Right now, the server process has the dependency on the client, but in a larger application this may be a dependency-of-a-dependency, which requires threading dependencies transitively through your application graph.

We can handle all of these issues by writing an initializer. An initializer is like a process, but only has an `Init` method, called in the same fashion. The following `Init` method reads the Redis address from the environment, constructs a client, pings the remote server, and adds the client to the service container with an application-distinct name.

```go
type redisInitializer struct {
	Config   *nacelle.Config           `service:"config"`
	Services *nacelle.ServiceContainer `service:"services"`
}

type redisConfig struct {
	Addr string `env:"redis_addr" default:"localhost:6379"`
}

func (i *redisInitializer) Init(ctx context.Context) error {
	redisConfig := &redisConfig{}
	if err := i.Config.Load(redisConfig); err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{Addr: redisConfig.Addr})
	if _, err := client.Ping().Result(); err != nil {
		return err
	}

	return i.Services.Set("redis", client)
}
```

The redis initializer has field with a `service` tag. This informs the bootstrapper to set the value of that field with the registered service with the same name. The *services* and *logger* services are available to all applications at startup. Similarly, we change the client field of the server process as follows. Note that any uses of this field must also similarly change the casing of the field name.

```go
type server struct {
	Config *nacelle.Config `service:"config"`
	Client *redis.Client `service:"redis"`
	server *http.Server
	port   int
}
```

Note that each injected field must be exported for the bootstrapper to access it. This changed the casing of the field, and will need to be changed within the HTTP handler as well.

Now, we can replace the ad-hoc client creation with the registration of the initializer that replaces it.

```go
func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	processes.RegisterInitializer(&redisInitializer{}, nacelle.WithMetaName("redis"))
	processes.RegisterProcess(&server{}, nacelle.WithMetaName("hw-server"))
	return nil
}
```
