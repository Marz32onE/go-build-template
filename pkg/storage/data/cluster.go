package storage

import (
	"gorm.io/gorm"
)

// Item represents an item in the storage.
type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
