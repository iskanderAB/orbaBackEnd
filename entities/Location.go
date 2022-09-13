package entities

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	AMD float64
	LAN float64
	UserID uint
}
func NewLocation(amd float64,lan float64 ,userid uint )Location{
	loc :=Location{
		AMD: amd,
		LAN: lan,
		UserID: userid,


	}
	return loc
}