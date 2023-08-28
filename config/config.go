package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server
	DB
}

type Server struct {
	Port         string `yaml:"port" env:"SERVER_PORT" env_default:"8000"`
	ReadTimeout  int    `yaml:"read-timeout" env:"SERVER_READ_TIMEOUT" env_default:"10"`
	WriteTimeout int    `yaml:"write-timeout" env:"SERVER_WRITE_TIMEOUT" env_default:"10"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env_default:"0.0.0.0"`
	Port     string `yaml:"port" env:"DB_PORT" env_default:"5432"`
	User     string `yaml:"user" env:"DB_USER" env-required:"true"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"DB_NAME" env_default:"segment_db"`
}

func Init(configFile string) (*Config, error) {
	config := &Config{}
	if err := cleanenv.ReadConfig(configFile, config); err != nil {
		return nil, fmt.Errorf("failed to read from config: %w", err)
	}

	log.Infof(config.Server.Port)
	return config, nil
}
