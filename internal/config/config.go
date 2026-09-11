// Package config 负责加载与解析服务配置。
package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config 是服务的根配置。
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
}

// ServerConfig 是 HTTP 服务相关配置。
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// LogConfig 是日志相关配置。
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Load 从指定 YAML 文件加载配置，环境变量可覆盖同名配置项
// （如 SERVER_PORT 覆盖 server.port）。
func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("server.port", 8080)
	v.SetDefault("server.mode", "debug")
	v.SetDefault("log.level", "info")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	return &cfg, nil
}
