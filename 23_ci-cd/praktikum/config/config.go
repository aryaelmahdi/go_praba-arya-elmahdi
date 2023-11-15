package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SECRET      string
	ProjectID   string
	SDKPath     string
	DatabaseURL string
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
	res := new(Config)
	if val, found := os.LookupEnv("SECRET"); found {
		res.SECRET = val
	}
	if val, found := os.LookupEnv("ProjetID"); found {
		res.SECRET = val
	}
	if val, found := os.LookupEnv("SDKPath"); found {
		res.SECRET = val
	}
	if val, found := os.LookupEnv("DatabaseURL"); found {
		res.SECRET = val
	}

	return res
}
