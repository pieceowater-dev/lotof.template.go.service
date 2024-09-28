package cfg

import (
	"application/internal/pkg/items/ent"
	_ "github.com/lib/pq"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

var models = []interface{}{
	&ent.Item{},
	// Add other models here
}

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	dsn := getPostgresDSN()
	log.Println("Connecting to database ...", dsn)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}
}

func getPostgresDSN() string {
	envInstance := &gossiper.Env{}
	val, err := envInstance.Get(GossiperConf.Env.Required[0])
	if err != nil {
		return ""
	}
	return val
}
