package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server
	DB
}

type Server struct {
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"read-timeout"`
	WriteTimeout int    `yaml:"write-timeout"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port" env:"DB_PORT" env_default:"5432"`
	User     string `yaml:"user" env:"DB_USER,POSTGRES_USER" env-required:"true"`
	Password string `yaml:"password" env:"DB_PASSWORD,POSTGRES_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"DB_NAME,POSTGRES_DB" env_default:"segment_db"`
}

func Init(configFile string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load variables from .env")
	}

	config := &Config{}
	if err := cleanenv.ReadConfig(configFile, config); err != nil {
		return nil, fmt.Errorf("failed to read from config: %w", err)
	}

	return config, nil
}
