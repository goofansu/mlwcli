package miniflux

import (
	miniflux "miniflux.app/v2/client"
)

func NewClient(endpoint, apiKey string) (*miniflux.Client, error) {
	return miniflux.NewClient(endpoint, apiKey), nil
}
