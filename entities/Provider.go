package entities

import (
	"image"

	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Name        string
	img         image.Image
	Description string
	UserID      uint
	Types []Type `gorm:"many2many:Provider_Type"`
	Categorys []Category `gorm:"many2many:Provider_Category"`
}