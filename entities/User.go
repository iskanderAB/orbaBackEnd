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


 
}
