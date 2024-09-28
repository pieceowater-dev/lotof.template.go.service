package ent

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
