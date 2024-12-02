package config

import (
	"golang-chapter-37/implem-gin-gonic/database"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database.InitDB()
}
