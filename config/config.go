package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func LoadConfig() (Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}
