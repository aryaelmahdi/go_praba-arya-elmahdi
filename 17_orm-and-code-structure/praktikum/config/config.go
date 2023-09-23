package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort int
	DBPort     int
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
}

func InitConfig() *Config {
	res := new(Config)
	res = loadConfig()

	if res == nil {
		logrus.Error("failed to load configuration")
		return nil
	}
	return res
}

func loadConfig() *Config {
	res := new(Config)
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : cannot load config file,", err.Error())
		return nil
	}

	if val, found := os.LookupEnv("ServerPort"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error()
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPort"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error()
		}
		res.DBPort = port
	}

	if val, found := os.LookupEnv("DBUsername"); found {
		res.DBUsername = val
	}

	if val, found := os.LookupEnv("DBPassword"); found {
		res.DBPassword = val
	}

	if val, found := os.LookupEnv("DBHost"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DBName"); found {
		res.DBName = val
	}

	return res
}
