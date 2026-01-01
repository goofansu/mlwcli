package miniflux

import (
	"encoding/json"
	"fmt"

	miniflux "miniflux.app/v2/client"

	"github.com/mmcdole/gofeed"
)

func AddFeed(client *miniflux.Client, feedURL string) error {
	fp := gofeed.NewParser()
	_, err := fp.ParseURL(feedURL)
	if err != nil {
		return fmt.Errorf("failed to parse feed URL: %w", err)
	}

	req := &miniflux.FeedCreationRequest{
		FeedURL:    feedURL,
		CategoryID: 1,
	}

	feedID, err := client.CreateFeed(req)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("✓ Feed created successfully (ID: %d)\n", feedID)
	return nil
}

func ListEntries(client *miniflux.Client, query string, starred string, limit int, all bool, jsonOutput bool) error {
	filter := &miniflux.Filter{
		Search:    query,
		Limit:     limit,
		Order:     "published_at",
		Direction: "desc",
	}
	if starred != "" {
		filter.Starred = starred
	}
	if !all {
		filter.Status = "unread"
	}

	result, err := client.Entries(filter)
	if err != nil {
		return fmt.Errorf("failed to list entries: %w", err)
	}

	if jsonOutput {
		output, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		fmt.Println(string(output))
	} else {
		fmt.Printf("Found %d entries:\n\n", result.Total)
		for _, entry := range result.Entries {
			fmt.Printf("• %s\n  %s\n\n", entry.Title, entry.URL)
		}
	}

	return nil
}
