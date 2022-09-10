package entities

import (
	

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Prix        float32
	Option 		Options 
	Additionals []Additional `gorm:"many2many:Product_Additional"`
	Rotes []Rote
	Favorites []Favorite 
	Command_lines []Command_line
	CategoryID uint
}