package storage

import (
	"log"

	data "github.com/marz32one/go-build-template/pkg/storage/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(filepath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&data.Item{})
}
