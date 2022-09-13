package entities

import (

	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	Name       string
}
func NewType(name string)Type{
   ty := Type{
		Name: name, 
	}
	return ty
}