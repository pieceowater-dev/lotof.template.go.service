package config

import (
	_ "github.com/lib/pq"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func GetPostgresDSN() string {
	envInstance := &gossiper.Env{}
	val, err := envInstance.Get(GossiperConf.Env.Required[0])
	if err != nil {
		return ""
	}
	return val
}
