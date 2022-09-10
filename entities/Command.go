package entities

import (

	"gorm.io/gorm"
)
type Command struct{
	gorm.Model
	TotalPrix float64
	Status string
	Attribute string
	UserID uint
	Command_lines []Command_line
}