package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
	JWTSecret  string
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}
	return config
}
