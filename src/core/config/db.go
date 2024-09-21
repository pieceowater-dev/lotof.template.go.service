package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	log "template/src/utils/logs"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func CheckDBConnection(dbURL string) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Error(fmt.Errorf("failed to open PostgreSQL connection: %w", err), nil)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(fmt.Errorf("failed to close PostgreSQL connection: %w", err), nil)
		}
	}(db)

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Error(fmt.Errorf("failed to ping PostgreSQL: %w", err), nil)
		return
	}

	log.Info("Connected to PostgreSQL", map[string]interface{}{
		"status": "connected",
		"url":    dbURL,
	})
}
