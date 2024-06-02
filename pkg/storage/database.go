package storage

import (
	"log"

	connection "github.com/marz32one/go-build-template/pkg/storage/connection"
	"gorm.io/gorm"
)

// DB is the global database connection.
var DB *gorm.DB

type Database interface {
	Connect(connectionString string) (*gorm.DB, error)
}

var DatabaseFactory = map[string]Database{
	"sqlite": &connection.SQLiteDB{},
	"mysql":  &connection.SQLDB{},
}

func GetDatabase(databaseType string, connectionDSN string) *gorm.DB {
	if db, ok := DatabaseFactory[databaseType]; ok {
		var err error
		DB, err = db.Connect(connectionDSN)
		if err != nil {
			log.Fatal("failed to connect database", err)
		}
	}
	return DB
}
