package storage

import (
	"gorm.io/gorm"
)

type DBHandler interface {
	Create(value interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}
