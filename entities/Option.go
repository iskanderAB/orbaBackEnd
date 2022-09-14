package entities

import (

	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Title       string
	Description string
	Prix        float64
	ProductID   uint 
	
}
func NewOption(title string,description string,prix float64,productid uint)Option{
	op:=Option{ 
		Title: title,
		Description: description,
		Prix: prix,
		ProductID: productid,

	}
	return op
}