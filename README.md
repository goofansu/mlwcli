# mlwcli

Manage Miniflux, Linkding, and Wallabag from terminal.

## Features

- Manage feeds via [Miniflux](https://miniflux.app/)
- Manage links via [Linkding](https://linkding.link/)
- Manage pages via [Wallabag](https://wallabag.org/)

## Installation

```bash
go install github.com/goofansu/mlwcli/cmd/mlwcli@latest
```

Or build from source:

```bash
git clone https://github.com/goofansu/mlwcli.git
cd mlwcli
go build -o mlwcli cmd/mlwcli/main.go
```

## Quick Start

### Authentication

Login and logout use an interactive TUI:

```bash
mlwcli auth login   # Interactive menu to select service and enter credentials
mlwcli auth logout  # Interactive menu showing signed-in services
```

### Managing Feeds (Miniflux)

```bash
mlwcli feed add https://miniflux.example.com/feed.xml
mlwcli feed list --jq='.items[] | {id, title}'
mlwcli entry list --status=unread --limit=50
mlwcli entry save 42  # Save to third-party integration
```

### Managing Links (Linkding)

```bash
mlwcli link add https://linkding.example.com --tags="cool useful"
mlwcli link list --search="example" --limit=20
```

### Managing Pages (Wallabag)

```bash
mlwcli page add https://wallabag.example.com/article --archive
mlwcli page list --starred --per-page=20
```

## Output Filtering

All list commands support JSON output with filtering:

```bash
# Get specific fields
mlwcli entry list --json=id,title,url

# Use jq expressions for complex filtering
mlwcli entry list --jq='.items[] | select(.feed.title == "Tech News")'

# Combine with external jq
mlwcli entry list --json=id,title,changed_at | jq '.items[] | select(.changed_at >= "2025-01-01")'
```

## Configuration

Configuration is stored at `~/.config/mlwcli/auth.toml` and includes endpoints and API keys for each service.

## Documentation

For detailed usage instructions and examples, see [skill/SKILL.md](skill/SKILL.md).
