package config

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
