package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/goofansu/cli/internal/config"
	"github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
)

func Login(service, endpoint, apiKey string) error {
	service = strings.ToLower(strings.TrimSpace(service))
	endpoint = strings.TrimSpace(endpoint)
	apiKey = strings.TrimSpace(apiKey)

	switch service {
	case config.ServiceMiniflux:
		if err := miniflux.Validate(endpoint, apiKey); err != nil {
			return fmt.Errorf("failed to verify miniflux connection: %w", err)
		}
	case config.ServiceLinkding:
		if err := linkding.Validate(endpoint, apiKey); err != nil {
			return fmt.Errorf("failed to verify linkding connection: %w", err)
		}
	default:
		return fmt.Errorf("invalid service: %s (must be '%s' or '%s')", service, config.ServiceMiniflux, config.ServiceLinkding)
	}

	if err := saveServiceConfig(service, endpoint, apiKey); err != nil {
		return err
	}
	fmt.Println("✓ Configuration saved successfully")

	return nil
}

func saveServiceConfig(service, endpoint, apiKey string) error {
	cfg := config.ServiceConfig{
		Endpoint: endpoint,
		APIKey:   apiKey,
	}

	appCfg, err := config.Load()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load config: %w", err)
	}
	if appCfg == nil {
		appCfg = &config.Config{}
	}

	switch service {
	case config.ServiceMiniflux:
		appCfg.Miniflux = cfg
	case config.ServiceLinkding:
		appCfg.Linkding = cfg
	}

	if err := config.Save(appCfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func Logout(service string) error {
	service = strings.ToLower(strings.TrimSpace(service))

	switch service {
	case config.ServiceMiniflux, config.ServiceLinkding:
		if err := config.RemoveService(service); err != nil {
			return fmt.Errorf("failed to remove %s config: %w", service, err)
		}
		fmt.Printf("✓ Logged out from %s successfully\n", service)
	default:
		return fmt.Errorf("invalid service: %s (must be '%s' or '%s')", service, config.ServiceMiniflux, config.ServiceLinkding)
	}

	return nil
}
