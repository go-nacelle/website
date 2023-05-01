---
title: "Config tag modifier implementations"
sidebarTitle: "Config tag modifiers"
category: "implementations"
index: 2
---

## Config tag modifier implementations

The [`github.com/go-nacelle/config`](https://github.com/go-nacelle/config) package provides the following [tag modifier](https://pkg.go.dev/github.com/go-nacelle/config#TagModifier) implementations.

---

#### Default tag setter

A [default tag setter](https://pkg.go.dev/github.com/go-nacelle/config#NewDefaultTagSetter) sets the `default` tag for a particular field. This is useful when the default values supplied by a library are inappropriate for a particular application. This would otherwise require a source change in the library.

#### Display tag setter

A [display tag setter](https://pkg.go.dev/github.com/go-nacelle/config#NewDisplayTagSetter) sets the `display` tag to the value of the `env` tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.

#### Flag tag setter

A [flag tag setter](https://pkg.go.dev/github.com/go-nacelle/config#NewFlagTagSetter) sets the `flag` tag to the value of the `env` tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.

#### File tag setter

A [file tag setter](https://pkg.go.dev/github.com/go-nacelle/config#NewFileTagSetter) sets the `file` tag to the value of the `env` tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.

#### Env tag prefixer

A [environment tag prefixer](https://pkg.go.dev/github.com/go-nacelle/config#NewEnvTagPrefixer) inserts a prefix on each `env` tags. This is useful when two distinct instances of the same configuration are required, and each one should be configured independently from the other (for example, using the same abstraction to consume from two different event busses with the same consumer code).

#### Flag tag prefixer

A [flag tag prefixer](https://pkg.go.dev/github.com/go-nacelle/config#NewFlagTagPrefixer) inserts a prefix on each `flag` tag. This effectively looks in a distinct top-level namespace in the parsed configuration. This is similar to the env tag prefixer.

#### File tag prefixer

A [file tag prefixer](https://pkg.go.dev/github.com/go-nacelle/config#NewFileTagPrefixer) inserts a prefix on each `file` tag. This effectively looks in a distinct top-level namespace in the parsed configuration. This is similar to the env tag prefixer.

**Example**

Tag modifiers are supplied at the time that a configuration struct is loaded. In the following example, each env tag is prefixed with `ACME_`, and the CassandraHosts field is given a default. Notice that you supply the *field* name to the tag modifier (not a tag value) when targeting a particular field value.

```go
if err := config.Load(
    appConfig,
    NewEnvTagPrefixer("ACME"),
    NewDefaultTagSetter("CassandraHosts", "[127.0.0.1:9042]"),
); err != nil {
    // handle error
}
```
