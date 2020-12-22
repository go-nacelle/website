---
title: "Tags"
category: "config"
index: 6
---

# Tags

### Tag Modifiers

A tag modifier dynamically alters the tags of a configuration struct. The following five tag modifiers are supplied. Additional behavior can be added by conforming to the *TagModifier* interface.

<dl>
  <dt>Default Tag Setter</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewDefaultTagSetter">default tag setter</a> sets the <code>default</code> tag for a particular field. This is useful when the default values supplied by a library are inappropriate for a particular application. This would otherwise require a source change in the library.</dd>

  <dt>Display Tag Setter</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewDisplayTagSetter">display tag setter</a> sets the <code>display</code> tag to the value of the <code>env</code> tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.</dd>

  <dt>Flag Tag Setter</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFlagTagSetter">flag tag setter</a> sets the <code>flag</code> tag to the value of the <code>env</code> tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.</dd>

  <dt>File Tag Setter</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFileTagSetter">file tag setter</a> sets the <code>file</code> tag to the value of the <code>env</code> tag. This tag modifier can be used to provide sane defaults to the tag without doubling the length of the struct tag definition.</dd>

  <dt>Env Tag Prefixer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewEnvTagPrefixer">environment tag prefixer</a> inserts a prefix on each <code>env</code> tags. This is useful when two distinct instances of the same configuration are required, and each one should be configured independently from the other (for example, using the same abstraction to consume from two different event busses with the same consumer code).</dd>

  <dt>Flag Tag Prefixer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFlagTagPrefixer">flag tag prefixer</a> inserts a prefix on each <code>flag</code> tag. This effectively looks in a distinct top-level namespace in the parsed configuration. This is similar to the env tag prefixer.</dd>

  <dt>File Tag Prefixer</dt>
  <dd>A <a href="https://godoc.org/github.com/go-nacelle/config#NewFileTagPrefixer">file tag prefixer</a> inserts a prefix on each <code>file</code> tag. This effectively looks in a distinct top-level namespace in the parsed configuration. This is similar to the env tag prefixer.</dd>
</dl>

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
