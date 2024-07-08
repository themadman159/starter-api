package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type databaseConfig struct {
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     int    `mapstructure:"DB_PORT"`
}

func LoadConfig() (*databaseConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config databaseConfig
	if err := viper.UnmarshalKey("database", &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
