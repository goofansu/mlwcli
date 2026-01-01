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

## Usage

### Authenticate with Services

Before using any commands, you need to authenticate with each service:

**Miniflux:**
```bash
cli login miniflux --endpoint https://miniflux.example.com --api-key YOUR_API_KEY
```

**Linkding:**
```bash
cli login linkding --endpoint https://linkding.example.com --api-key YOUR_API_KEY
```

Configuration is stored in `~/.config/cli/config.toml`.

### Manage Bookmarks (Linkding)

Add a new bookmark:

```bash
cli links add https://example.com
```

Add a bookmark with notes:

```bash
cli links add https://example.com --notes "Interesting article"
```

Add a bookmark with tags:

```bash
cli links add https://example.com --tags "golang api"
```

Add a bookmark with both notes and tags:

```bash
cli links add https://example.com --notes "Great resource" --tags "dev tools"
```

### Manage Feeds (Miniflux)

Add a new feed:

```bash
cli feeds add https://example.com/feed.xml
```

List entries:

```bash
cli feeds list
```

List with options:

```bash
# List all entries (default is unread only)
cli feeds list --all

# Search entries
cli feeds list --search "query text"

# List starred entries
cli feeds list --starred

# Limit results
cli feeds list --limit 10

# Output in JSON format
cli feeds list --json

# Combine options
cli feeds list --all --starred --limit 20
cli feeds list --search "golang" --limit 50
```

### Logout

Remove credentials for a specific service:

```bash
cli logout miniflux
cli logout linkding
```

## Configuration

The CLI stores configuration in `~/.config/cli/config.toml` with the following format:

```toml
[miniflux]
endpoint = "https://miniflux.example.com"
api_key = "your_api_key"

[linkding]
endpoint = "https://linkding.example.com"
api_key = "your_api_key"
```

## Requirements

- Go 1.25.4 or later
- A running Miniflux instance (for feed management)
- A running Linkding instance (for bookmark management)
