package app

import (
	"encoding/json"
	"fmt"

	"github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
)

func outputJSON(data any) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(output))
	return nil
}

type ListEntriesOptions struct {
	FeedID  int64
	Search  string
	Starred string
	Limit   int
	All     bool
}

type ListFeedsOptions struct{}

type FeedOutput struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	FeedURL string `json:"feed_url"`
}

type ListBookmarksOptions struct {
	Query string
	Limit int
}

type EntryOutput struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type BookmarkOutput struct {
	ID          int64    `json:"id"`
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Notes       string   `json:"notes"`
	TagNames    []string `json:"tag_names"`
}

func (a *App) ListEntries(opts ListEntriesOptions) error {
	result, err := miniflux.Entries(a.Config.Miniflux.Endpoint, a.Config.Miniflux.APIKey, miniflux.ListEntriesOptions{
		FeedID:  opts.FeedID,
		Search:  opts.Search,
		Starred: opts.Starred,
		Limit:   opts.Limit,
		All:     opts.All,
	})
	if err != nil {
		return fmt.Errorf("failed to list entries: %w", err)
	}

	entries := make([]EntryOutput, len(result.Entries))
	for i, entry := range result.Entries {
		entries[i] = EntryOutput{
			ID:    entry.ID,
			Title: entry.Title,
			URL:   entry.URL,
		}
	}

	return outputJSON(entries)
}

func (a *App) ListFeeds(opts ListFeedsOptions) error {
	result, err := miniflux.Feeds(a.Config.Miniflux.Endpoint, a.Config.Miniflux.APIKey)
	if err != nil {
		return fmt.Errorf("failed to list feeds: %w", err)
	}

	feeds := make([]FeedOutput, len(result))
	for i, feed := range result {
		feeds[i] = FeedOutput{
			ID:      feed.ID,
			Title:   feed.Title,
			FeedURL: feed.FeedURL,
		}
	}

	return outputJSON(feeds)
}

func (a *App) ListBookmarks(opts ListBookmarksOptions) error {
	result, err := linkding.ListBookmarks(a.Config.Linkding.Endpoint, a.Config.Linkding.APIKey, linkding.ListBookmarksOptions{
		Query: opts.Query,
		Limit: opts.Limit,
	})
	if err != nil {
		return fmt.Errorf("failed to list bookmarks: %w", err)
	}

	bookmarks := make([]BookmarkOutput, len(result.Results))
	for i, bookmark := range result.Results {
		bookmarks[i] = BookmarkOutput{
			ID:          int64(bookmark.ID),
			URL:         bookmark.URL,
			Title:       bookmark.Title,
			Description: bookmark.Description,
			Notes:       bookmark.Notes,
			TagNames:    bookmark.TagNames,
		}
	}

	return outputJSON(bookmarks)
}
