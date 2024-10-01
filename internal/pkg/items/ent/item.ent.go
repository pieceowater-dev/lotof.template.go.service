package ent

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
