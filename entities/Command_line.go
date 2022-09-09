package entities

import "gorm.io/gorm"

type Command_line struct {
	gorm.Model
	Quantity int16
}