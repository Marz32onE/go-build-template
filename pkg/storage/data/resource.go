package data

import (
	"gorm.io/gorm"
)

// Resource represents an item in the storage.
type Resource struct {
	gorm.Model
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	CPU             int     `json:"cpu"`
	Memory          int     `json:"memory"`
	CPUThreshold    float64 `json:"cpu_threshold"`
	MemoryThreshold float64 `json:"memory_threshold"`
}
