# cli

My command-line tool built for agents.

## Features

- Manage bookmarks via [Linkding](https://linkding.link/)
- Manage feeds via [Miniflux](https://miniflux.app/)
- Manage pages via [Wallabag](https://wallabag.org/)

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
cli login miniflux -h
cli login linkding -h
cli login linkding -h
```

## Usage

See [skill/SKILL.md](skill/SKILL.md).
