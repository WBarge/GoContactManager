package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	First  string
	Last   string
	Emails []Email
	Phones []Phone `gorm:"many2many:person_phones;"`
}
