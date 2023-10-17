package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	OpenAPIKey string
}

func InitConfig() *Config {
	res := new(Config)
	res = LoadConfig()
	if res == nil {
		logrus.Error("Config : configuration error")
		return nil
	}
	return res
}

func LoadConfig() *Config {
	res := new(Config)

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("cannot load config file", err.Error())
		return nil
	}

	if key, found := os.LookupEnv("OpenAPIKey"); found {
		res.OpenAPIKey = key
	}

	return res
}
