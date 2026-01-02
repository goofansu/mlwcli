# cli

A unified command-line interface for managing bookmarks (via Linkding) and feeds (via Miniflux).

## Installation

```bash
go install github.com/goofansu/cli/cmd/cli@latest
```

Or build from source:

```bash
git clone https://github.com/goofansu/cli.git
cd cli
go build -o cli cmd/cli/main.go
```

## Authentication

```bash
cli login miniflux --endpoint <URL> --api-key <KEY>
cli login linkding --endpoint <URL> --api-key <KEY>
```

## Usage

See [skill/SKILL.md](skill/SKILL.md).
