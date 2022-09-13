package entities

import (
  
	"gorm.io/gorm"
	
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	Email    string  `json:"email"`
	Gender string    `json:"gendre"`
	Telephone int32   `json:"telephone"`    
	Password int16   `json:"password"`
	location Location  
	Providers[] Provider 
	Favorites[] Favorite 
	Reclamations[] Reclamation
	Commands [] Command
	Rotes [] Rote
	Notifications []Notification `gorm:"many2many:User_Notifications;"`


 
}
func NewUser(name string,role string,email string,gendre string, telephone int32,password int16,location Location)User{
	user:=User{
		Fullname: name,
		Role:role,
		Email: email,
		Gender: gendre,
		Telephone: telephone,
		Password: password,
		location: location,
		Providers: []Provider{} ,
		Favorites: []Favorite{},
		Reclamations: []Reclamation{},
		Rotes: []Rote{},
		Notifications: []Notification{},
		

	}
	return user
}