package entities

import (
  
	"gorm.io/gorm"
	
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	Email    string  `json:"email" `
	Gender string    `json:"gendre"`
	Telephone int32   `json:"telephone"`    
	Password string   `json:"password"`
	location Location  
	Providers[] Provider 
	Favorites[] Favorite 
	Reclamations[] Reclamation
	Commands [] Command
	Rotes [] Rote
	Notifications []Notification `gorm:"many2many:User_Notifications;"`


 
}
func NewUser(name string,role string,email string,gendre string, telephone int32,password string,location Location)User{
	user:=User{
		Model:         gorm.Model{},
		Fullname:      name,
		Role:          role,
		Email:         email,
		Gender:        gendre,
		Telephone:     telephone,
		Password:      password,
		// location:      location,
		// Providers:     []Provider{},
		// Favorites:     []Favorite{},
		// Reclamations:  []Reclamation{},
		// Commands:      []Command{},
		// Rotes:         []Rote{},
		// Notifications: []Notification{},
	}
	return user
}