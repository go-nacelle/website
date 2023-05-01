---
title: 'Part 2: Setting up a skeleton "Hello, World!" server'
sidebarTitle: "Part 2: Server"
category: "tutorial"
index: 2
---

## Setting up a skeleton "Hello, World!" server

This section of the tutorial enhances the application to (trivially) respond to HTTP requests.

---

Let's create simple HTTP server process that responds with a 200 OK/Hello World response for each request.

Each process is initialized by calling its `Init` function with a configuration container (more on this in the next section). On initialization success, the `Start` method is invoked in a dedicated go-routine. This method is expected to block for long-running processes such as servers. On application shutdown, the `Stop` method is invoked which should unblock any active work being done in the `Start` method.

Here, the `Init` method creates a server and configures its handler. The `Start` creates a TCP listener and starts serving HTTP traffic. The `Stop` method signals for the server to stop accepting new connections and shutdown.

```go
import (
	"context"
	"fmt"
	"net"
	"net/http"

	// ...
)

type server struct {
	server *http.Server
	port   int
}

func (s *server) Init(ctx context.Context) error {
	s.server = &http.Server{}
	s.server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!\n"))
	})
	return nil
}

func (s *server) Run(ctx context.Context) error {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	defer listener.Close()
	defer s.server.Close()

	// Run server, block until shutdown (do not return ErrServerClosed)
	if err := s.server.Serve(listener); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.server.Shutdown(context.Background())
	return nil
}
```

Now, we modify the setup function to register this server process.

```go
func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	processes.RegisterProcess(&server{port: 5000}, nacelle.WithMetaName("hw-server"))
	return nil
}
```

If you were to run this application, you would see Nacelle initialize and start the *hw-server* process. Curl-ing any path at `http://localhost:5000` will return the same 200-level response.
