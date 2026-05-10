package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

var lock = &sync.Mutex{}

var Conf *Config

func InitializeConfig() {
	if Conf == nil {
		lock.Lock()
		defer lock.Unlock()
		if Conf == nil {
			Conf = GetConfig()
		}
	}
}

func getIntEnvVariable(name string) int {
	envVar := os.Getenv(name)
	if envVar == "" {
		logger.Fatal("Env variable %s required", name)
	}

	intValue, err := strconv.Atoi(envVar)
	if err != nil {
		logger.Fatal("Env variable %s must be integer", name)
	}

	return intValue
}

func getStringEnvVariable(name string) string {
	envVar := os.Getenv(name)
	if envVar == "" {
		logger.Fatal("Env variable %s required", name)
	}

	return envVar
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	return &Config{
		Token: getStringEnvVariable("TOKEN"),
	}

}
