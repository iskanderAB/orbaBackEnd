package entities

import "image"

type Additional struct {
	ID    string
	Title string
	Img   image.Image
	Attribute string
	Prix float32 
	Description string
}