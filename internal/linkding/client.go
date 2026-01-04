package linkding

import (
	api "github.com/piero-vic/go-linkding"
)

type CreateLinkOptions struct {
	URL      string
	Notes    string
	TagNames []string
}

type ListLinksOptions struct {
	Query  string
	Limit  int
	Offset int
}

func CreateLink(endpoint, apiKey string, opts CreateLinkOptions) (*api.Bookmark, error) {
	client := api.NewClient(endpoint, apiKey)

	req := api.CreateBookmarkRequest{
		URL:      opts.URL,
		Notes:    opts.Notes,
		TagNames: opts.TagNames,
	}

	return client.CreateBookmark(req)
}

func ListLinks(endpoint, apiKey string, opts ListLinksOptions) (*api.ListBookmarksResponse, error) {
	client := api.NewClient(endpoint, apiKey)
	return client.ListBookmarks(api.ListBookmarksParams{
		Query:  opts.Query,
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
}

func Validate(endpoint, apiKey string) error {
	client := api.NewClient(endpoint, apiKey)
	_, err := client.GetUserPreferences()
	return err
}
