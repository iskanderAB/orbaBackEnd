package entities

import "gorm.io/gorm"

type Rote struct {
	gorm.Model
	Attribute  string
	Rate int16
	UserID uint
	ProductID uint
}
func NewRote(attribute string , rate int16, userid uint,productid uint)Rote{
	rote :=Rote{
		Attribute: attribute,
		Rate: rate,
		UserID: userid,
		ProductID: productid,

	}
	return rote
}