package entities

import (

	"gorm.io/gorm"
)
type Command struct{
	gorm.Model
	TotalPrix float64
	Status string
	Attribute string
	UserID uint
	Command_lines []Command_line
}
func NEwCommand(prix float64,status string,attribute string, userid uint,)Command{
	comm :=Command{
		TotalPrix: prix,
		Status: status,
		Attribute: attribute,
		UserID: userid,
		Command_lines: []Command_line{},
	}
	return comm

}