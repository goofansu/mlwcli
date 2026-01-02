package app

import (
	"fmt"

	"github.com/goofansu/cli/internal/format"
	"github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
)

type ListEntriesOptions struct {
	FeedID  int64
	Search  string
	Limit   int
	Status  string
	Starred string
	JSON    string
	JQ      string
}

type ListBookmarksOptions struct {
	Query string
	Limit int
	JSON  string
	JQ    string
}

func (a *App) ListEntries(opts ListEntriesOptions) error {
	result, err := miniflux.Entries(a.Config.Miniflux.Endpoint, a.Config.Miniflux.APIKey, miniflux.ListEntriesOptions{
		FeedID:  opts.FeedID,
		Search:  opts.Search,
		Limit:   opts.Limit,
		Status:  opts.Status,
		Starred: opts.Starred,
	})
	if err != nil {
		return fmt.Errorf("failed to list entries: %w", err)
	}

	return format.Output(result.Entries, opts.JSON, opts.JQ)
}

func (a *App) ListBookmarks(opts ListBookmarksOptions) error {
	result, err := linkding.ListBookmarks(a.Config.Linkding.Endpoint, a.Config.Linkding.APIKey, linkding.ListBookmarksOptions{
		Query: opts.Query,
		Limit: opts.Limit,
	})
	if err != nil {
		return fmt.Errorf("failed to list bookmarks: %w", err)
	}

	return format.Output(result.Results, opts.JSON, opts.JQ)
}
