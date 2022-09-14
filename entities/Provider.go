package entities

import (

	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Name        string `JSON:"name"`
	img         string  `JSON:"img"`
	Description string  `JSON:"description"`
	UserID      uint	 `JSON:"userid"`
	Types []Type `gorm:"many2many:Provider_Type"`
	Categorys []Category `gorm:"many2many:Provider_Category"`
}
func NewProvider (name string,img string, description string, userid uint)Provider{
	prov:=Provider{
		Name: name,
		Description: description,
		UserID: userid,
		Types:[]Type{},
		Categorys: []Category{},

	}
	return prov
}