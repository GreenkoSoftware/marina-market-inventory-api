package sql

import (
	"gorm.io/gorm"
)

func By(field string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" = ?", value)
	}
}
