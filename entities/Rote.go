package entities

import "gorm.io/gorm"

type Rote struct {
	gorm.Model
	Attribute  string
	ProductID  string
	Rate int16
}