package config

import (
	"log"

	"github.com/joho/godotenv"
)

var PostgresDSN string

func init() {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error Loading .env")
	}
	PostgresDSN = envMap["POSTGRES_DSN"]
}
