---
title: "Sourcers"
category: "config"
index: 4
---

# Sourcers

#### Sourcers

A sourcer reads values from a particular source based on a configuration struct's tags. Sourcers declare the struct tags that determine their behavior when loading configuration structs. The examples above only work with the environment sourcer. All sources support the `default` and `required` tags (which are mutually exclusive). Tagged fields must be exported. The following six sourcers are supplied. Additional behavior can be added by conforming to the *Sourcer* interface.

<dl>
  <dt>Environment Sourcer</dt>
  <dd>An <a href="https://godoc.org/github.com/go-nacelle/config#NewEnvSourcer">environment sourcer</a> reads the <code>env</code> tag and looks up the corresponding value in the process's environment. An expected prefix may be supplied in order to namespace application configuration from the rest of the system. A sourcer instantiated with <code>NewEnvSourcer("APP")</code> will load the env tag <code>fetch_limit</code> from the environment variable <code>APP_FETCH_LIMIT</code> and falling back to the environment variable <code>FETCH_LIMIT</code>.</dd>

  <dt>Test Environment Sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewTestEnvSourcer">test environment sourcer</a> reads the <code>env</code> tag but looks up the corresponding value from a literal map. This sourcer can be used in unit tests where the full construction of a nacelle <a href="https://nacelle.dev/docs/core/process">process</a> is too burdensome.</dd>

  <dt>Flag Sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFlagSourcer">flag sourcer</a> reads the <code>flag</code> tag and looks up the corresponding value attached to the process's command line arguments.</dd>

  <dt>File Sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFileSourcer">file sourcer</a> reads the <code>file</code> tag and returns the value at the given path. A filename and a file parser musts be supplied on instantiation. Both <a href="https://godoc.org/github.com/go-nacelle/config#ParseYAML">ParseYAML</a> and <a href="https://godoc.org/github.com/go-nacelle/config#ParseTOML">ParseTOML</a> are supplied file parsers -- note that as JSON is a subset of YAML, <code>ParseYAML</code> will also correctly parse JSON files. If a <code>nil</code> file parser is supplied, one is chosen by the filename extension. A file sourcer will load the file tag <code>api.timeout</code> from the given file by parsing it into a map of values and recursively walking the (keys separated by dots). This can return a primitive type or a structured map, as long as the target field has a compatible type. The constructor <a href="https://godoc.org/github.com/go-nacelle/config#NewOptionalFileSourcer">NewOptionalFileSourcer</a> will return a no-op sourcer if the filename does not exist.</dd>

  <dt>Multi sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewMultiSourcer">multi-sourcer</a> is a sourcer wrapping one or more other sourcers. For each configuration struct field, each sourcer is queried in reverse order of registration and the first value to exist is returned. This is useful to allow a chain of configuration files in which some files or directories take precedence over others, or to allow environment variables to take precedence over files.</dd>

  <dt>Directory Sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewDirectorySourcer">directory sourcer</a> creates a multi-sourcer by reading each file in a directory in alphabetical order. The constructor <a href="https://godoc.org/github.com/go-nacelle/config#NewOptionalDirectorySourcer">NewOptionalDirectorySourcer</a> will return a no-op sourcer if the directory does not exist.</dd>

  <dt>Glob Sourcer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewGlobSourcer">glob sourcer</a> creates a multi-sourcer by reading each file that matches a given glob pattern. Each matching file creates a distinct file sourcer and does so in alphabetical order.</dd>
</dl>
