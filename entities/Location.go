package entities

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	AMD float64
	LAN float64
	UserID uint
}