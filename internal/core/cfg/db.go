package cfg

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once // Ensures DB is only initialized once
)

// GetDB safely returns the initialized *gorm.DB instance.
// If db has not been initialized, this will return nil.
func GetDB() *gorm.DB {
	return db
}

// SetDB safely sets the *gorm.DB instance.
// Uses sync.Once to prevent reinitialization of the DB.
func SetDB(database *gorm.DB) {
	once.Do(func() {
		db = database
	})
}
