package app

import (
	"github.com/goofansu/cli/internal/config"
)

type App struct {
	Config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{Config: cfg}
}
