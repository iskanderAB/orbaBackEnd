package entities

import "gorm.io/gorm"

type Rote struct {
	gorm.Model
	Attribute  string
	Rate int16
	UserID uint
	ProductID uint
}