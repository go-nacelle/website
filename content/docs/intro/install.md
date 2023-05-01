---
title: "Quick install guide"
category: "intro"
index: 2
---

## Quick install guide

How to install Nacelle to get up and running.

---

Nacelle requires a minimum version of `Go 1.16`.

```bash
go get -u github.com/go-nacelle/nacelle/v2
```

Installing the `nacelle` package will automatically install core dependencies (`config`, `log`, `service,` and `process`). If your project makes use of non-core Nacelle frameworks, abstract process libraries, or utility libraries, they must be installed separately and explicitly. For example:

```bash
go get -u github.com/go-nacelle/pgutil
```

For the most recent stable tagged version, check the [latest release](https://github.com/go-nacelle/nacelle/releases/latest). For information about the latest bug fixes, updates, and features added to the framework, check out the [changelog](https://github.com/go-nacelle/nacelle/blob/master/CHANGELOG.md).
