package linkding

import (
	"github.com/piero-vic/go-linkding"
)

func NewClient(endpoint, apiKey string) (*linkding.Client, error) {
	return linkding.NewClient(endpoint, apiKey), nil
}
