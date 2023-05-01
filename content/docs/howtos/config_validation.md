---
title: "How to validate user-supplied configuration values"
sidebarTitle: "...validate user-supplied configuration values"
category: "howtos"
index: 1
noSidebarLink: true
---

## How to validate user-supplied configuration values

This guide describes a feature of the [`github.com/go-nacelle/config`](https://github.com/go-nacelle/config) package. See [related documentation](/docs/topics/config).

---

After successful loading of a configuration struct, the method named `PostLoad` will be called if it is defined. This allows a place for additional validation (such as mutually exclusive settings, regex value matching, etc) and deserialization of more complex types (such enums from strings, durations from integers, etc). The following example parses and stores a `text/template` from a user-supplied string.

```go
import "text/template"

type Config struct {
	RawTemplate    string `env:"template" default:"Hello, {{.Name}}!"`
	ParsedTemplate *template.Template
}

func (c *Config) PostLoad() (err error) {
	c.ParsedTemplate, err = template.New("ConfigTemplate").Parse(c.RawTemplate)
	return
}
```

An error returned by `PostLoad` will be returned via the `Load` method.
