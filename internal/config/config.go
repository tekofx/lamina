package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/tekofx/lamina/internal/logger"
)

type Config struct {
	Token       string
	MediaFolder string
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
		Token:       getStringEnvVariable("TOKEN"),
		MediaFolder: "media",
	}

}

func SetupFolders() {
	err := os.MkdirAll("media", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
