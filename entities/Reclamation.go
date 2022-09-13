package entities

import "gorm.io/gorm"

type Reclamation struct {
	gorm.Model
	Description string
	UserID uint
}
func NewReclamatin (description string,userid uint)Reclamation{
	rec :=Reclamation{
		Description: description,
		UserID: userid,

	}
	return rec
}