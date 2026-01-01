package auth

import (
	"fmt"
	"strings"

	"github.com/goofansu/cli/internal/config"
	linkdingclient "github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
	linkdinglib "github.com/piero-vic/go-linkding"
)

func Login(service, endpoint, apiKey string) error {
	service = strings.ToLower(strings.TrimSpace(service))
	endpoint = strings.TrimSpace(endpoint)
	apiKey = strings.TrimSpace(apiKey)

	switch service {
	case "miniflux":
		cfg := config.ServiceConfig{
			Endpoint: endpoint,
			APIKey:   apiKey,
		}

		client, err := miniflux.NewClient(endpoint, apiKey)
		if err != nil {
			return fmt.Errorf("failed to create miniflux client: %w", err)
		}

		if _, err := client.Me(); err != nil {
			return fmt.Errorf("failed to verify connection: %w", err)
		}

		appCfg, err := config.Load()
		if err != nil && !strings.Contains(err.Error(), "no such file") {
			return fmt.Errorf("failed to load config: %w", err)
		}
		if appCfg == nil {
			appCfg = &config.Config{}
		}
		appCfg.Miniflux = cfg

		if err := config.Save(appCfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("✓ Configuration saved successfully")
		fmt.Println("✓ Connection verified")

	case "linkding":
		cfg := config.ServiceConfig{
			Endpoint: endpoint,
			APIKey:   apiKey,
		}

		client, err := linkdingclient.NewClient(endpoint, apiKey)
		if err != nil {
			return fmt.Errorf("failed to create linkding client: %w", err)
		}

		if _, err := client.ListBookmarks(linkdinglib.ListBookmarksParams{}); err != nil {
			return fmt.Errorf("failed to verify connection: %w", err)
		}

		appCfg, err := config.Load()
		if err != nil && !strings.Contains(err.Error(), "no such file") {
			return fmt.Errorf("failed to load config: %w", err)
		}
		if appCfg == nil {
			appCfg = &config.Config{}
		}
		appCfg.Linkding = cfg

		if err := config.Save(appCfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("✓ Configuration saved successfully")
		fmt.Println("✓ Connection verified")

	default:
		return fmt.Errorf("invalid service: %s (must be 'miniflux' or 'linkding')", service)
	}

	return nil
}

func Logout(service string) error {
	service = strings.ToLower(strings.TrimSpace(service))

	switch service {
	case "miniflux", "linkding":
		if err := config.RemoveService(service); err != nil {
			return fmt.Errorf("failed to remove %s config: %w", service, err)
		}
		fmt.Printf("✓ Logged out from %s successfully\n", service)
	default:
		return fmt.Errorf("invalid service: %s (must be 'miniflux' or 'linkding')", service)
	}

	return nil
}
