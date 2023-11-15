package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Secret string
}

func Init() *Config {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	res := loadConfig()
	if res == nil {
		logrus.Fatal("Cannot load configuration")
		return nil
	}
	return res
}

func loadConfig() *Config {
	cfg := Config{}
	if val, found := os.LookupEnv("SECRET"); found {
		cfg.Secret = val
	}
	return &cfg
}
