package entities

import (
	

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Prix        float32
}