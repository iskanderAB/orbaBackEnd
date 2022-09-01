package entities

import "image"

type Provider struct {
	ID          string
	Name        string
	img         image.Image
	Description string
}