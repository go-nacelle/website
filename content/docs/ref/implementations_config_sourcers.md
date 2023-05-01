---
title: "Config sourcer implementations"
sidebarTitle: "Config sourcers"
category: "implementations"
index: 1
---

## Config sourcer implementations

The [`github.com/go-nacelle/config`](https://github.com/go-nacelle/config) package provides the following [sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#Sourcer) implementations.

---

#### Environment sourcer

An [environment sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewEnvSourcer) reads the `env` tag and looks up the corresponding value in the process's environment. An expected prefix may be supplied in order to namespace application configuration from the rest of the system. A sourcer instantiated with `NewEnvSourcer("APP")` will load the env tag `fetch_limit` from the environment variable `APP_FETCH_LIMIT` and falling back to the environment variable `FETCH_LIMIT`.

#### Test environment sourcer

A [test environment sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewTestEnvSourcer) reads the `env` tag but looks up the corresponding value from a literal map. This sourcer can be used in unit tests where the full construction of a nacelle [process](/docs/topics/process) is too burdensome.

#### Flag sourcer

A [flag sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewFlagSourcer) reads the `flag` tag and looks up the corresponding value attached to the process's command line arguments.

#### File sourcer

A [file sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewFileSourcer) reads the `file` tag and returns the value at the given path. A filename and a file parser musts be supplied on instantiation. Both [ParseYAML](https://pkg.go.dev/github.com/go-nacelle/config/v3#ParseYAML) and [ParseTOML](https://pkg.go.dev/github.com/go-nacelle/config/v3#ParseTOML) are supplied file parsers -- note that as JSON is a subset of YAML, `ParseYAML` will also correctly parse JSON files. If a `nil` file parser is supplied, one is chosen by the filename extension. A file sourcer will load the file tag `api.timeout` from the given file by parsing it into a map of values and recursively walking the (keys separated by dots). This can return a primitive type or a structured map, as long as the target field has a compatible type. The constructor [NewOptionalFileSourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewOptionalFileSourcer) will return a no-op sourcer if the filename does not exist.

#### Multi sourcer

A [multi-sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewMultiSourcer) is a sourcer wrapping one or more other sourcers. For each configuration struct field, each sourcer is queried in reverse order of registration and the first value to exist is returned. This is useful to allow a chain of configuration files in which some files or directories take precedence over others, or to allow environment variables to take precedence over files.

#### Directory sourcer

A [directory sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewDirectorySourcer) creates a multi-sourcer by reading each file in a directory in alphabetical order. The constructor [NewOptionalDirectorySourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewOptionalDirectorySourcer) will return a no-op sourcer if the directory does not exist.

#### Glob sourcer

A [glob sourcer](https://pkg.go.dev/github.com/go-nacelle/config/v3#NewGlobSourcer) creates a multi-sourcer by reading each file that matches a given glob pattern. Each matching file creates a distinct file sourcer and does so in alphabetical order.
