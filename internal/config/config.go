package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/marz32one/go-build-template/pkg/storage"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dbType := os.Getenv("DB_TYPE")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	database := storage.GetDatabase(dbType, connectionString)
	if database == nil {
		log.Fatalf("failed to connect database %s:%s", dbType, connectionString)
	}
}
