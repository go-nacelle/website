# Nacelle Docs Website

This website is built via [hugo](https://gohugo.io/) and continuously deployed via [CircleCI](https://circleci.com/).

### Developing Locally

Running `hugo serve` will start a local live-reload server on port 1313. Editing the markdown files in `content`, the theme in `layout`, or the style and image files in `static` will trigger a rebuild.

### Automatic Documentation

Many of the markdown files contain only a manually-written header followed by the token `<!-- Fold -->`. These files will be auto-updated based on the README content for that project.

Run `./sync-repos.sh` to update the content *below the fold* in each markdown file with the content of the markdown file on GitHub. The following repositories in the `go-nacelle` organization will be synchronized.

| Repository | Source path             | Documentation path |
| ---------- | ----------------------- | ------------------ |
| awsutil    | README.md               | ./content/docs/libraries/awsutil.md |
| chevron    | README.md               | ./content/docs/frameworks/chevron.md |
| config     | README.md               | ./content/docs/core/config.md |
| grpcbase   | README.md               | ./content/docs/base processes/grpcbase.md |
| httpbase   | README.md               | ./content/docs/base processes/httpbase.md |
| lambdabase | README.md               | ./content/docs/base processes/lambdabase.md |
| log        | README.md               | ./content/docs/core/log.md |
| nacelle    | docs/docs.md            | ./content/docs/core/overview.md |
| nacelle    | docs/getting-started.md | ./content/docs/quickstart/overview.md |
| pgutil     | README.md               | ./content/docs/libraries/pgutil.md |
| process    | README.md               | ./content/docs/core/process.md |
| scarf      | README.md               | ./content/docs/frameworks/scarf.md |
| service    | README.md               | ./content/docs/core/service.md |
| workerbase | README.md               | ./content/docs/base processes/workerbase.md |

### Enabling Automatic Documentation

The repositories listed above use the orb published from `.circleci/docs.orb.yml` to automatically update the documentation for that repository when the master branch of the repository is built. This orb will query the README content presently on GitHub, update the associated markdown file, and (if there were content changes) create a pull request to this project.

The following is a minimal CircleCI configuration to run the orb for a given repository.

```
version: 2.1

orbs:
  docs: nacelle/docs@0.1.3

jobs:
  update_docs:
    executor: go
    steps:
      - docs/update_docs:
          repo: <my repository name>
```
