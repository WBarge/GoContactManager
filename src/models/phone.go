package models

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	CountryCode string
	AreaCode    string
	Number      string
	PhoneType   string
}
