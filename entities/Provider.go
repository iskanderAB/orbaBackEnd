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
}