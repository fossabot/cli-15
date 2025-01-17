# Calyptia Cloud CLI

[![CI](https://github.com/calyptia/cli/actions/workflows/ci.yml/badge.svg)](https://github.com/calyptia/cloud-cli/actions/workflows/ci.yml)
[![Codecov](https://codecov.io/gh/calyptia/cli/branch/main/graph/badge.svg?token=TY36W7B87A)](https://codecov.io/gh/calyptia/cli)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcalyptia%2Fcli.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcalyptia%2Fcli?ref=badge_shield)

This CLI interacts with the [Calyptia Cloud](https://cloud.calyptia.com) service
using the [API Client](https://github.com/calyptia/api).
Futher documentation on how to use the CLI
can be found on [Calyptia Docs](https://docs.calyptia.com).

## Build instructions

---

```markdown
go mod download
go build
```

---

## Install

You can get the latest release artifacts for the major operating systems
at the [Releases](https://github.com/calyptia/cloud-cli/releases) page.

Alternatively, you can use `Brew`:

```bash
brew tap calyptia/tap
brew install calyptia
```

You can even install latest from `main` branch using `Go`:

---

```bash
go install github.com/calyptia/cli@latest
```

---

## Run

The first command you would want to run is `config set_token` otherwise
you will have to always pass `--token` around.

Get a token (API key) from [cloud.calyptia.com](https://cloud.calyptia.com).

---

```bash
calyptia config set_token TOKEN
```

---

Alternatively, you can set the CALYPTIA_CLOUD_TOKEN environment variable or
pass the token on each command, as an example:

---

```bash
calyptia get members --token TOKEN
```

---

## Environment variables

A list of the supported environment variables that will override the provided flags.

- CALYPTIA_CLOUD_URL: URL of the cloud API (default: https://cloud-api.calyptia.com/)
- CALYPTIA_CLOUD_TOKEN: Cloud project token (default: None)
- CALYPTIA_STORAGE_DIR: Path to store the local configuration (fallback to $HOME/.calyptia)

## Commands

```bash
Calyptia Cloud CLI

Usage:
  calyptia [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Configure Calyptia CLI
  create      Create core instances, pipelines, etc.
  delete      Delete core instances, pipelines, etc.
  get         Display one or many resources
  help        Help about any command
  rollout     Rollout resources to previous versions
  top         Display metrics
  update      Update core instances, pipelines, etc.

Flags:
      --cloud-url string   Calyptia Cloud URL (default "https://cloud-api.calyptia.com")
  -h, --help               help for calyptia
      --token string       Calyptia Cloud Project token
  -v, --version            version for calyptia

Use "calyptia [command] --help" for more information about a command.
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcalyptia%2Fcli.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcalyptia%2Fcli?ref=badge_large)