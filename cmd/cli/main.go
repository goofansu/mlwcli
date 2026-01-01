package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/goofansu/cli/internal/auth"
	"github.com/goofansu/cli/internal/config"
	"github.com/goofansu/cli/internal/linkding"
	"github.com/goofansu/cli/internal/miniflux"
)

type Options struct {
	Login  LoginCommand  `command:"login" description:"Authenticate with a service"`
	Logout LogoutCommand `command:"logout" description:"Remove credentials for a service"`
	Links  LinksCommand  `command:"links" description:"Manage bookmarks (linkding)"`
	Feeds  FeedsCommand  `command:"feeds" description:"Manage feeds (miniflux)"`
}

type LoginCommand struct {
	Args struct {
		Service string `positional-arg-name:"service" description:"Service name (miniflux or linkding)" required:"yes"`
	} `positional-args:"yes"`
	Endpoint string `long:"endpoint" description:"Service endpoint URL" required:"yes"`
	APIKey   string `long:"api-key" description:"API key" required:"yes"`
}

type LogoutCommand struct {
	Args struct {
		Service string `positional-arg-name:"service" description:"Service name (miniflux or linkding)" required:"yes"`
	} `positional-args:"yes"`
}

type LinksCommand struct {
	Add AddLinksCommand `command:"add" description:"Add a new bookmark"`
}

type AddLinksCommand struct {
	Args struct {
		URL string `positional-arg-name:"url" description:"URL of the bookmark to add" required:"yes"`
	} `positional-args:"yes"`
	Notes string `long:"notes" description:"Optional notes for the bookmark"`
	Tags  string `long:"tags" description:"Optional tags separated by spaces"`
}

type FeedsCommand struct {
	Add  AddFeedsCommand  `command:"add" description:"Add a new feed"`
	List ListFeedsCommand `command:"list" description:"List entries"`
}

type AddFeedsCommand struct {
	Args struct {
		URL string `positional-arg-name:"url" description:"URL of the feed to add" required:"yes"`
	} `positional-args:"yes"`
}

type ListFeedsCommand struct {
	Limit   int    `long:"limit" description:"Maximum number of results" default:"30"`
	Search  string `long:"search" description:"Search query text"`
	Starred bool   `long:"starred" description:"Filter by starred entries"`
	All     bool   `long:"all" description:"List all entries (default is unread only)"`
	JSON    bool   `long:"json" description:"Output in JSON format"`
}

func (c *LoginCommand) Execute(_ []string) error {
	return auth.Login(c.Args.Service, c.Endpoint, c.APIKey)
}

func (c *LogoutCommand) Execute(_ []string) error {
	return auth.Logout(c.Args.Service)
}

func (c *AddLinksCommand) Execute(_ []string) error {
	url := c.Args.URL
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config (run 'cli login linkding' first): %w", err)
	}
	if cfg.Linkding.Endpoint == "" {
		return fmt.Errorf("not logged in to linkding (run 'cli login linkding' first)")
	}
	client, err := linkding.NewClient(cfg.Linkding.Endpoint, cfg.Linkding.APIKey)
	if err != nil {
		return fmt.Errorf("failed to create linkding client: %w", err)
	}
	return linkding.AddBookmark(client, url, c.Notes, c.Tags)
}

func (c *AddFeedsCommand) Execute(_ []string) error {
	feedURL := c.Args.URL
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config (run 'cli login miniflux' first): %w", err)
	}
	if cfg.Miniflux.Endpoint == "" {
		return fmt.Errorf("not logged in to miniflux (run 'cli login miniflux' first)")
	}
	client, err := miniflux.NewClient(cfg.Miniflux.Endpoint, cfg.Miniflux.APIKey)
	if err != nil {
		return fmt.Errorf("failed to create miniflux client: %w", err)
	}
	return miniflux.AddFeed(client, feedURL)
}

func (c *ListFeedsCommand) Execute(_ []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config (run 'cli login miniflux' first): %w", err)
	}
	if cfg.Miniflux.Endpoint == "" {
		return fmt.Errorf("not logged in to miniflux (run 'cli login miniflux' first)")
	}
	client, err := miniflux.NewClient(cfg.Miniflux.Endpoint, cfg.Miniflux.APIKey)
	if err != nil {
		return fmt.Errorf("failed to create miniflux client: %w", err)
	}
	starred := ""
	if c.Starred {
		starred = "1"
	}
	return miniflux.ListEntries(client, c.Search, starred, c.Limit, c.All, c.JSON)
}

func (c *LoginCommand) Usage() string {
	return "<service> [OPTIONS]"
}

func (c *LogoutCommand) Usage() string {
	return "<service>"
}

func (c *AddLinksCommand) Usage() string {
	return "<url>"
}

func (c *AddFeedsCommand) Usage() string {
	return "<url>"
}

func (c *ListFeedsCommand) Usage() string {
	return "[OPTIONS]"
}

func main() {
	opts := Options{}
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	parser.ShortDescription = "Unified bookmarks and feeds command-line interface"
	parser.LongDescription = "cli provides a unified CLI for managing bookmarks (linkding) and feeds (miniflux).\n\nAuthenticate with login, then use links or feeds subcommands.\n\nExamples:\n  cli login miniflux --endpoint https://miniflux.example.com --api-key YOUR_API_KEY\n  cli login linkding --endpoint https://linkding.example.com --api-key YOUR_API_KEY\n  cli links add https://example.com\n  cli links add https://example.com --notes \"Interesting article\" --tags \"golang api\"\n  cli feeds add https://example.com/feed.xml\n  cli feeds list\n  cli feeds list --search \"golang\"\n  cli logout miniflux\n  cli logout linkding"

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		return
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				fmt.Fprint(os.Stdout, flagsErr.Message)
				return
			}
			fmt.Fprintf(os.Stderr, "error: %s\n\n", flagsErr.Message)
			parser.WriteHelp(os.Stderr)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
