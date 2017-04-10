# pagerduty-cli

A command line interface for PagerDuty written in Go.

* This project depends on a forked version of the `PagerDuty/go-pagerduty` library. Once [PR #68](https://github.com/PagerDuty/go-pagerduty/pull/68) is merged, we can go back to using the official library. *

# Usage

`pagerduty-cli` only supports creating and listing maintenance windows for the time being. Subcommands also have `--help`.

```bash
$ pagerduty-cli --help
A CLI interface for PagerDuty.

Usage:
  pagerduty-cli [command]

Available Commands:
  mwindow     Maintenance Windows

Flags:
      --api-token string   PagerDuty API Token. Also supports environment variable "PGDUTY_TOKEN".

Use "pagerduty-cli [command] --help" for more information about a command.
```

# Development

To develop this package, you first need to install the forked dependency. `go get` is insufficient for this because it will install the official library. To install the one we want to use, navigate to `$GOPATH/src/github.com` and clone the forked version into the "official" destination.

```
$ mkdir -p PagerDuty/go-pagerduty
$ git clone git@github.com:ethanfrogers/go-pagerduty.git PagerDuty/go-pagerduty
```

From there, you will be using the fork, but it will look like you are using the official library. This is a shortcoming of the Golang dependency system, in my opinion.