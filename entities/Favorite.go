package entities

import (

	"gorm.io/gorm"
)
type Favorite struct{
	gorm.Model
	UserID uint
	ProductID uint



}
func NewFavorite(userid uint,productid uint)Favorite{
	fav:= Favorite{
		UserID: userid,
		ProductID: productid,
	}
	return fav
}