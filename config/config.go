package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	configViper *viper.Viper
}

func NewConfig() *Config {

	configViper := viper.New()
	configViper.SetConfigFile("./config/env.json")
	newConfig := Config{configViper}

	if err := configViper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &newConfig
}

func (config *Config) GetValue(key string) string {

	strValue := config.configViper.GetString(key)
	return strValue
}
