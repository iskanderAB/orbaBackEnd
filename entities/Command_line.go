package entities

import "gorm.io/gorm"

type Command_line struct {
	gorm.Model
	Quantity int16
	ProductID uint
	CommandID uint
}
func NewCommand_line(quantity int16,proudctid uint,commandid uint )Command_line{
	com_line:=Command_line{
		Quantity: quantity,
		ProductID: proudctid,
		CommandID: commandid,

	}
	return com_line
}