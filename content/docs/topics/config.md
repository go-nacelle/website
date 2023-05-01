---
title: "Configuration"
defines: "config"
category: "topics"
index: 2
---

## Configuration

The [`github.com/go-nacelle/config`](https://github.com/go-nacelle/config) package provides configuration loading and validation.

---

The major behavioral components of an application will often need external configuration during their setup. Such values can be pulled from a **configuration loader** initialized with a particular **sourcer** (e.g. that reads the environment or disk) and assigned to the tagged fields of a user-defined struct.

We'll use the following code snippet as a running example. This example defines and loads configuration for a hypothetical worker [process](/docs/topics/process). For the application to start successfully, the address of an API must be supplied. All other configuration values are optional.

```go
type Config struct {
	APIAddr        string   `env:"api_addr" required:"true"`
	CassandraHosts []string `env:"cassandra_hosts"`
	NumWorkers     int      `env:"num_workers" default:"10"`
	BufferSize     int      `env:"buffer_size" default:"1024"`
}

type Process struct {
	Config *nacelle.Config `service:"config"`
}

func (p *Process) Init(ctx context.Context) error {
	appConfig := &Config{}
	if err := p.Config.Load(appConfig); err != nil {
		return err
	}

	// Use or store populated appConfig
	p.config = appConfig
	return nil
}
```

Configuration structs are defined by the application or library developer with the fields needed by the package in which they are defined. Each field is tagged with a **source hint** (e.g. an environment variable name, a key in a YAML file) that instructs a particular _sourcer_ how to load a value and, optionally, default values and basic validation. 

A [sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#Sourcer) reads values from a particular configuration source based on the tags of a given struct to be populated. Sourcers advertise a set of struct tags that control their behavior. The example above (presumably) uses an environment sourcer, thus the configuration struct's _source hints_ refers to the process environment. All sources support the `default` and `required` tags (which are mutually exclusive). Tagged fields must be exported (so the user supplied values can be assigned to them).

The values that are loaded into non-string field of a configuration struct are interpreted as JSON. Supplying the process the (partial) environment `CASSANDRA_HOSTS='["host1", "host2", "host3"]'` would populate the struct's `CassandraHosts` field with a slice containing the values `host1`, `host2`, and `host3`. The defaults for such fields can also be supplied as a JSON-encoded string, but must be escaped to preserve the format of the struct tag.

```go
type Config struct {
	CassandraHosts []string `env:"cassandra_hosts" default:"[\"host1\", \"host2\", \"host3\"]"`
}
```

Loading configuration values also works with structs containing composite fields, which is a good way to share a "core" of configuration values shared by distinct components of an application.

While struct tags are static, they can be dynamically altered by supplying a [tag modifier](https://pkg.go.dev/github.com/go-nacelle/config/v3#TagModifier) to the configuration loader. Tag modifiers make certain tasks possible, such as re-using a configuration struct with a different prefix, or conditionally disabling required attribute validation based on some runtime property.

In the example above, the required configuration is populated and stored at process initialization time. The `Load` method of the configuration loader fails if a value from the source cannot be converted into the correct type, a value from the source cannot be unmarshalled into the target shape, or is required and not supplied. 

After each successful load of a configuration struct, the loaded configuration values will be logged. This, however, may be a concern for application secrets. In order to hide sensitive configuration values, add the `mask:"true"` struct tag to the field. This will omit that value from log messages related to configuration loading. Additionally, configuration loader object can be initialized with a blacklist of values that should be masked (values printed as `*****` rather than their real value) instead of omitted. These values can be configured in the [bootstrapper](/docs/topics/booting).

You can see an additional example of loading configuration in the [example repository](https://github.com/go-nacelle/example): [definition](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L13) and [loading](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L36).

### Related resources

- [Sourcer implementations](/docs/ref/implementations_config_sourcers)
- [Tag modifier implementations](/docs/ref/implementations_config_tag_modifiers)
- [How to validate user-supplied configuration values](/docs/howtos/config_validation)
