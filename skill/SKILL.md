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
cli login linkding --endpoint https://linkding.example.com --api-key YOUR_API_KEY
```

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
cli add bookmark <url> [OPTIONS]
```

Add a new bookmark to your Linkding instance.

Examples:
```bash
# Simple bookmark
cli add bookmark https://example.com

# Bookmark with notes
cli add bookmark https://example.com --notes "Interesting article"

# Bookmark with tags
cli add bookmark https://example.com --tags "golang api"

# Bookmark with notes and tags
cli add bookmark https://example.com --notes "Great resource" --tags "dev tools"
```

#### List Bookmarks

```bash
cli list bookmarks [OPTIONS]
```

List bookmarks from your Linkding instance.

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

#### Add a New Feed

```bash
cli add feed <url> [OPTIONS]
```

Add a new RSS/Atom feed to your Miniflux instance.

**Important:** Before adding a feed, verify that the URL returns valid RSS/Atom XML. The Miniflux instance may be rate limited by the target server if it repeatedly attempts to fetch invalid or non-existent feed URLs. Use a tool like `curl` or a browser to validate the feed first.

Examples:
```bash
# Add feed to default category (category 1)
cli add feed https://example.com/feed.xml

# Add feed to a specific category
cli add feed https://example.com/feed.xml --category-id 5
```

#### List Entries

```bash
cli list entries [OPTIONS]
```

List feed entries (ordered by publication date, newest first). Default shows only unread entries.

Examples:
```bash
# List latest 10 unread entries
cli list entries

# List all entries
cli list entries --all

# Search entries
cli list entries --search "golang"

# List starred entries with limit
cli list entries --starred --limit 50

# Filter by feed ID
cli list entries --feed-id 42

# Combine multiple filters
cli list entries --all --search "golang" --limit 20
```

#### List Feeds

```bash
cli list feeds
```

List all subscribed feeds.

Example:
```bash
cli list feeds
```

## Help

Display help for any command:

```bash
cli --help
cli login --help
cli logout --help
cli add --help
cli add feed --help
cli add bookmark --help
cli list --help
cli list feeds --help
cli list entries --help
cli list bookmarks --help
```
