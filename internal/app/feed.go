package app

import (
	"fmt"

	"github.com/goofansu/cli/internal/miniflux"
)

type AddFeedOptions struct {
	URL        string
	CategoryID int64
}

func (a *App) AddFeed(opts AddFeedOptions) error {
	categoryID := opts.CategoryID
	if categoryID == 0 {
		categoryID = 1
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
