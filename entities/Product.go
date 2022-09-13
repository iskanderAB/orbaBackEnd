package entities

import (
	

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Prix        float32
	Option 		Option 
	Additionals []Additional `gorm:"many2many:Product_Additional"`
	Rotes []Rote
	Favorites []Favorite 
	Command_lines []Command_line
	CategoryID uint
}

func NewProduct(title string,description string,prix float64,option Option,categoryid uint)Product{
	prod:=Product{
		Title: title,
		Description: description,
		Prix: float32(prix),
		Additionals: []Additional{},
		Rotes: []Rote{},
		Favorites: []Favorite{},
		Command_lines: []Command_line{},
		CategoryID: categoryid,

	}
	return prod
}
