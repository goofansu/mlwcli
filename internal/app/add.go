package app

import (
	"fmt"
	"strings"

	"github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
)

type AddBookmarkOptions struct {
	URL   string
	Notes string
	Tags  string
}

type AddFeedOptions struct {
	URL        string
	CategoryID int64
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

func (a *App) AddFeed(opts AddFeedOptions) error {
	categoryID := opts.CategoryID
	if categoryID == 0 {
		categoryID = 1 // Default to category 1 if not specified
	}

	feedID, err := miniflux.CreateFeed(a.Config.Miniflux.Endpoint, a.Config.Miniflux.APIKey, miniflux.CreateFeedOptions{
		FeedURL:    opts.URL,
		CategoryID: categoryID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("✓ Feed created successfully (ID: %d)\n", feedID)
	return nil
}
