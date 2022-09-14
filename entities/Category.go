package entities

import (

	"gorm.io/gorm"
)
type Category struct{
	gorm.Model
	Name string
	Products []Product 
}
func NewCategory(name string)Category{
	cat := Category{
		Model:    gorm.Model{},
		Name:     name,
		Products: []Product{},
	}
	return cat
}