// Package config loads and parses service configuration.
package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config is the root configuration of the service.
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// LogConfig holds logging settings.
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Load reads configuration from the given YAML file; environment variables
// override same-named keys (e.g. SERVER_PORT overrides server.port).
func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("server.port", 8080)
	v.SetDefault("server.mode", "debug")
	v.SetDefault("log.level", "info")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
