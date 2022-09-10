package entities

import (
  
	"gorm.io/gorm"
	
)

type User struct {
	gorm.Model
	Fullname string
	Role     string
	Email    string
	Gender string
	Telephone int32
	Password int16
	location Location 
	Providers[] Provider 
	Favorites[] Favorite 
	Reclamations[] Reclamation
	Commands [] Command
	Rotes [] Rote
	Notifications []Notification `gorm:"many2many:User_Notifications;"`


 
}
