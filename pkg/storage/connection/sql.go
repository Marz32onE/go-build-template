package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLDB represents a connection to the SQL database.
type SQLDB struct{}

// Connect connects to the database using the provided connection string.
func (m *SQLDB) Connect(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
