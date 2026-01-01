package linkding

import (
	"fmt"
	"strings"

	"github.com/piero-vic/go-linkding"
)

func AddBookmark(client *linkding.Client, url, notes, tags string) error {
	tagNames := []string{}
	if tags != "" {
		tagNames = strings.Fields(tags)
	}

	req := linkding.CreateBookmarkRequest{
		URL:      url,
		Notes:    notes,
		TagNames: tagNames,
	}

	bookmark, err := client.CreateBookmark(req)
	if err != nil {
		return fmt.Errorf("failed to create bookmark: %w", err)
	}

	fmt.Printf("✓ Bookmark created successfully (ID: %d)\n", bookmark.ID)
	return nil
}
