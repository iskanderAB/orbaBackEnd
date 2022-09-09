package entities

import "gorm.io/gorm"

type Reclamation struct {
	gorm.Model
	Description string
}