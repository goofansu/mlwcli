package app

import (
	"fmt"
	"strings"

	"github.com/goofansu/cli/internal/format"
	"github.com/goofansu/cli/internal/linkding"
)

type AddBookmarkOptions struct {
	URL   string
	Notes string
	Tags  string
}

type ListBookmarksOptions struct {
	Query  string
	Limit  int
	Offset int
	JSON   string
	JQ     string
}

func (a *App) AddBookmark(opts AddBookmarkOptions) error {
	tagNames := []string{}
	if opts.Tags != "" {
		tagNames = strings.Fields(opts.Tags)
	}

	bookmark, err := linkding.CreateBookmark(a.Config.Linkding.Endpoint, a.Config.Linkding.APIKey, linkding.CreateBookmarkOptions{
		URL:      opts.URL,
		Notes:    opts.Notes,
		TagNames: tagNames,
	})
	if err != nil {
		return fmt.Errorf("failed to create bookmark: %w", err)
	}

	fmt.Printf("✓ Bookmark created successfully (ID: %d)\n", bookmark.ID)
	return nil
}

func (a *App) ListBookmarks(opts ListBookmarksOptions) error {
	result, err := linkding.ListBookmarks(a.Config.Linkding.Endpoint, a.Config.Linkding.APIKey, linkding.ListBookmarksOptions{
		Query:  opts.Query,
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
	if err != nil {
		return fmt.Errorf("failed to list bookmarks: %w", err)
	}

	output := map[string]any{
		"total": result.Count,
		"items": result.Results,
	}
	return format.Output(output, opts.JSON, opts.JQ)
}
