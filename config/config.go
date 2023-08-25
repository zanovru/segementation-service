package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Port         string `env-required:"true" yaml:"port" env:"PORT" env_default:"8000"`
		ReadTimeout  int    `yaml:"read-timeout" env:"SERVER_READ_TIMEOUT" env_default:"10"`
		WriteTimeout int    `yaml:"write-timeout" env:"SERVER_WRITE_TIMEOUT" env_default:"10"`
	} `yaml:"server"`
}

func Init(configFile string) (*Config, error) {
	config := &Config{}
	if err := cleanenv.ReadConfig(configFile, config); err != nil {
		return nil, fmt.Errorf("failed to read from config: %w", err)
	}

	log.Infof(config.Server.Port)
	return config, nil
}
