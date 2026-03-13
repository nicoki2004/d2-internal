// Package config, admin all the app config
package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/subosito/gotenv"

	"github.com/nicoki2004/d2-internal/internal/logger"
)

type Config struct {
	APIKey      string
	ClientID    string
	Secret      string
	RedirectURL string
}

var (
	instance *Config
	once     sync.Once
)

// Get returns the unique instance of the configuration
func Get() *Config {
	once.Do(func() {
		err := gotenv.Load()
		if err != nil {
			logger.GetLogger().Debug("No env files")
		}

		instance = &Config{
			APIKey:      os.Getenv("BUNGIE_API_KEY"),
			ClientID:    os.Getenv("BUNGIE_OAUTH_CLIENT_ID"),
			Secret:      os.Getenv("BUNGIE_OAUTH_CLIENT_SECRET"),
			RedirectURL: os.Getenv("BUNGIE_OAUTH_REDIRECT_URI"),
		}

		// Validation of required fields
		if err := instance.Validate(); err != nil {
			logger.GetLogger().Fatal("Error en configuración: %v", err)
		}
	})

	return instance
}

// Validate verifies that all required fields are present
func (c *Config) Validate() error {
	missing := []string{}

	if c.APIKey == "" {
		missing = append(missing, "BUNGIE_API_KEY")
	}
	if c.ClientID == "" {
		missing = append(missing, "BUNGIE_OAUTH_CLIENT_ID")
	}
	if c.Secret == "" {
		missing = append(missing, "BUNGIE_OAUTH_CLIENT_SECRET")
	}
	if c.RedirectURL == "" {
		missing = append(missing, "BUNGIE_OAUTH_REDIRECT_URI")
	}

	if len(missing) > 0 {
		return fmt.Errorf("configuración incompleta. Variables faltantes: %v", missing)
	}

	return nil
}
