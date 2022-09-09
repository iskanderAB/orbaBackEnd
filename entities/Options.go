package entities

import (

	"gorm.io/gorm"
)

type Options struct {
	gorm.Model
	Title       string
	Description string
	Prix        float64
	
}