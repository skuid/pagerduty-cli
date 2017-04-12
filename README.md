# pagerduty-cli

A command line interface for PagerDuty written in Go.

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