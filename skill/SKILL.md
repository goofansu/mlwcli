---
name: cli
description: Unified command-line interface for managing bookmarks (linkding) and feeds (miniflux). Use for authentication, managing bookmarks, and managing feeds.
---

# CLI

CLI is a unified command-line interface for managing bookmarks (via Linkding) and feeds (via Miniflux). Use this to authenticate with services, add bookmarks with notes and tags, and manage RSS/Atom feeds.

## Getting Started

Before using any commands (except `login` and `logout`), you must authenticate with each service.

### Login

Authenticate with a service:

```bash
cli login <service> --endpoint <URL> --api-key <KEY>
```

- `<service>`: Service name - either `miniflux` or `linkding`
- `--endpoint`: Your service instance URL (e.g., `https://miniflux.example.com`)
- `--api-key`: API key/token from service Settings

Examples:

**Miniflux:**
```bash
cli login miniflux --endpoint https://miniflux.example.com --api-key YOUR_API_KEY
```

**Linkding:**
```bash
cli login linkding --endpoint https://linkding.example.com --api-key YOUR_API_TOKEN
```

The configuration is saved to `~/.config/cli/config.toml` and verified automatically.

### Logout

Remove stored credentials for a specific service:

```bash
cli logout <service>
```

Examples:
```bash
cli logout miniflux
cli logout linkding
```

## Commands

### Manage Bookmarks (Linkding)

#### Add a New Bookmark

```bash
cli links add <url> [OPTIONS]
```

Add a new bookmark to your Linkding instance.

**Note:** If a bookmark with the same URL already exists, it will be edited/updated with the provided notes and tags instead of creating a duplicate.

**Options:**
- `--notes <text>`: Optional notes for the bookmark (simple string)
- `--tags <tags>`: Optional tags separated by spaces (same convention as Linkding web UI)

Examples:
```bash
# Simple bookmark
cli links add https://example.com

# Bookmark with notes
cli links add https://example.com --notes "Interesting article"

# Bookmark with tags
cli links add https://example.com --tags "golang api"

# Bookmark with notes and tags
cli links add https://example.com --notes "Great resource" --tags "dev tools"
```

### Manage Feeds (Miniflux)

#### Add a New Feed

```bash
cli feeds add <url>
```

Add a new RSS/Atom feed to your Miniflux instance.

**Important:** Before adding a feed, verify that the URL returns valid RSS/Atom XML. The Miniflux instance may be rate limited by the target server if it repeatedly attempts to fetch invalid or non-existent feed URLs. Use a tool like `curl` or a browser to validate the feed first.

Example:
```bash
cli feeds add https://example.com/feed.xml
```

#### List Entries

```bash
cli feeds list [OPTIONS]
```

List feed entries (ordered by publication date, newest first). Default shows only unread entries.

**Options:**
- `--limit <n>`: Maximum number of results (default: 30)
- `--search <query>`: Search through entries with query text
- `--starred`: Filter by starred entries only
- `--all`: List all entries (default is unread only)
- `--json`: Output in JSON format

Examples:
```bash
# List latest 30 unread entries
cli feeds list

# List all entries
cli feeds list --all

# Search entries
cli feeds list --search "golang"

# List starred entries with limit
cli feeds list --starred --limit 50

# Combine multiple filters
cli feeds list --all --search "golang" --limit 20
cli feeds list --search "golang" --json
```

## Configuration

Authentication credentials are stored in `~/.config/cli/config.toml` with the following format:

```toml
[miniflux]
endpoint = "https://miniflux.example.com"
api_key = "your-miniflux-api-key"

[linkding]
endpoint = "https://linkding.example.com"
api_key = "your-linkding-api-token"
```

## Error Handling

If you see "not logged in" errors, run `cli login <service>` to set up credentials for that service.

## Help

Display help for any command:

```bash
cli --help
cli login --help
cli logout --help
cli links add --help
cli feeds add --help
cli feeds list --help
```
