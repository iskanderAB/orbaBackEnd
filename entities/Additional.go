package entities

import (
	"image"

	"gorm.io/gorm"
)

type Additional struct {
    gorm.Model
	Title string
	Img   image.Image
	Attribute string
	Prix float32 
	Description string
}