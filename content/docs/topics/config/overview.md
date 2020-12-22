---
title: "Config"
defines: "config"
category: "topics"
index: 2
---

# Config

Configuration loading and validation for [nacelle](https://nacelle.dev).

---

Often, [initializers and processes](/docs/topics/process) will need external configuration during their startup process. These values can be pulled from a **configuration loader** backed by a particular [source](#sourcers) (e.g. environment or disk) and assigned to tagged fields of a configuration struct.

You can see an additional example of loading configuration in the [example repository](https://github.com/go-nacelle/example): [definition](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L13) and [loading](https://github.com/go-nacelle/example/blob/843979aaa86786784a1ca3646e8d0d1f69e29c65/internal/redis_initializer.go#L36).

- [json](/docs/topics/config/json)
- [loading](/docs/topics/config/loading)
- [nesting](/docs/topics/config/nesting)
- [sourcers](/docs/topics/config/sourcers)
- [struct](/docs/topics/config/struct)
- [tags](/docs/topics/config/tags)
- [validation](/docs/topics/config/validation)
