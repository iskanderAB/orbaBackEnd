package entities

import (
	"image"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Description string
	Title       string
	Icon        image.Image
	Img   		image.Image
}