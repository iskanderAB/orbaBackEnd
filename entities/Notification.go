package entities

import (

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Description string
	Title       string
	Icon        string
	Img   		string
}
func NewNotification(description string , title string,icon string,img string)Notification{
	not :=Notification{
		Description: description,
		Title: title,
		Icon: icon,
		Img: img,
	}
	return not
}