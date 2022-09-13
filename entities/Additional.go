package entities

import (
	

	"gorm.io/gorm"
)

type Additional struct {
    gorm.Model
	Title string
	Img   string
	Attribute string
	Prix float32 
	Description string
}
func NewAdditional (tittle string, img string,attribute string,prix float32,description string)Additional{
	add := Additional{
		Title: tittle,
		Img: img,
		Attribute: attribute,
		Prix: prix,
		Description: description,

	}
	return add

}