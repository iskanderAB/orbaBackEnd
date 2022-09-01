package entities

import (
	"image"
	"time"
)

type Notification struct {
	ID          string
	Description string
	Title       string
	Icon        image.Image
	Img   		image.Image
	Created_at  time.Time
}