package models

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	Address      string
	DefaultEmail bool
	PersonID     uint
}
