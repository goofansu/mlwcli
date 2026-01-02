---
name: cli
description: Unified command-line interface for managing bookmarks (linkding) and feeds (miniflux). Use for authentication, managing bookmarks, and managing feeds.
---

# CLI

Unified CLI for managing bookmarks (Linkding) and feeds (Miniflux).

## Authentication

Services: `miniflux` (feeds), `linkding` (bookmarks). Use `cli login --help` for setup.

## Default List Commands

```bash
cli list bookmarks --jq ".[] | { id, url, title, description, notes, tag_names }"
cli list entries --jq ".[] | { id, url, title, published_at, status, feed_id: .feed.id, feed_title: .feed.title }"
```

## Discovering Fields

```bash
cli list entries --jq ".[0] | keys"
cli list bookmarks --jq ".[0] | keys"
```

## Commands

```bash
cli login <service>       # Authenticate with miniflux or linkding
cli logout <service>      # Remove stored credentials
cli add bookmark <url>    # Add bookmark to Linkding
cli add feed <url>        # Add feed to Miniflux
cli list bookmarks        # List bookmarks
cli list entries          # List feed entries
```

Use `--help` on any command for options.

## Notes

- `--jq` works independently (doesn't require `--json`)
- Validate feed URLs before adding to avoid rate limiting
