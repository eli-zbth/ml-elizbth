package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Errorf("error trying to read local .env file: %s", err)
	}

	properties := []string{
		"MONGO_DB_URI",
		"SHORT_URL_DOMAIN",
		"REDIS_URL",
		"PORT",
	}

	for _,key := range properties {
		viper.Set(key, os.Getenv(key))
	}
}