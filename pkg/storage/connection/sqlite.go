package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDB represents a SQLite database connection.
type SQLiteDB struct{}

// Connect establishes a connection to the SQLite database using the provided connection string.
func (s *SQLiteDB) Connect(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
