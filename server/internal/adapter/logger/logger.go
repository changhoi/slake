package logger

import (
	"github.com/changhoi/slake/internal/adapter/config"
	"go.uber.org/zap"
)

func NewAdapter(cfg *config.Config) (*zap.Logger, error) {
	if cfg.App.Debug {
		return zap.NewDevelopment()
	}

	if cfg.App.Profile < config.ProfileProd {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
