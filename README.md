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
cli add bookmark https://example.com
```

Add a bookmark with notes:

```bash
cli add bookmark https://example.com --notes "Interesting article"
```

Add a bookmark with tags:

```bash
cli add bookmark https://example.com --tags "golang api"
```

Add a bookmark with both notes and tags:

```bash
cli add bookmark https://example.com --notes "Great resource" --tags "dev tools"
```

List bookmarks:

```bash
cli list bookmarks [OPTIONS]
```

**Options:**
- `--limit <n>`: Maximum number of results (default: 10)
- `--search <query>`: Search through bookmarks with query text

Examples:
```bash
# List all bookmarks
cli list bookmarks

# List bookmarks with limit
cli list bookmarks --limit 10

# Search bookmarks
cli list bookmarks --search "golang"

# Combine options
cli list bookmarks --search "api" --limit 20
```

### Manage Feeds (Miniflux)

Add a new feed:

```bash
cli add feed https://example.com/feed.xml
```

Add a feed to a specific category:

```bash
cli add feed https://example.com/feed.xml --category-id 5
```

Note: Before adding a feed, verify that the URL returns valid RSS/Atom XML. Use a tool like `curl` or a browser to validate the feed first. If `--category-id` is not specified, the feed will be added to category 1 by default.

List feeds:

```bash
cli list feeds
```

List entries:

```bash
cli list entries
```

List entries with options:

```bash
# List all entries (default is unread only)
cli list entries --all

# Search entries
cli list entries --search "query text"

# List starred entries
cli list entries --starred

# Limit results
cli list entries --limit 10

# Filter by feed ID
cli list entries --feed-id 42

# Combine options
cli list entries --all --starred --limit 20
cli list entries --search "golang" --limit 50
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
