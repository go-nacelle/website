# Nacelle Docs Website

This website is built via [hugo](https://gohugo.io/) and continuously deployed via [CircleCI](https://circleci.com/).

### Developing Locally

Running `hugo serve` will start a local live-reload server on port 1313. Editing the markdown files in `content`, the theme in `layout`, or the style and image files in `static` will trigger a rebuild.

### Automatic Documentation

Many of the markdown files contain only a manually-written header followed by the token `<!-- Fold -->`. These files will be auto-updated based on the README content for that project.

Run `./sync-repos.sh` to update the content *below the fold* for each of these markdown files with the README content presently on GitHub. The following repositories in the `go-nacelle` organization will be synchronized with the associated markdown file.

| Repository | DocumentationPath |
| ---------- | ----------------- |
| awsutil    | ./content/docs/libraries/awsutil.md |
| chevron    | ./content/docs/frameworks/chevron.md |
| config     | ./content/docs/core/config.md |
| grpcbase   | ./content/docs/base processes/grpcbase.md |
| httpbase   | ./content/docs/base processes/httpbase.md |
| lambdabase | ./content/docs/base processes/lambdabase.md |
| log        | ./content/docs/core/log.md |
| nacelle    | ./content/docs/core/overview.md |
| pgutil     | ./content/docs/libraries/pgutil.md |
| process    | ./content/docs/core/process.md |
| scarf      | ./content/docs/frameworks/scarf.md |
| service    | ./content/docs/core/service.md |
| workerbase | ./content/docs/base processes/workerbase.md |

### Enabling Automatic Documentation

The repositories listed above use the orb published from `.circleci/docs.orb.yml` to automatically update the documentation for that repository when the master branch of the repository is built. This orb will query the README content presently on GitHub, update the associated markdown file, and (if there were content changes) create a pull request to this project.

The following is a minimal CircleCI configuration to run the orb for a given repository.

```
version: 2.1

orbs:
  docs: nacelle/docs@0.1.2

jobs:
  update_docs:
    executor: go
    steps:
      - docs/update_docs:
          repo: <my repository name>
```
