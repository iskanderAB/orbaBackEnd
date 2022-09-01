package entities

import "time" 
type Command struct{
	ID string
	TotalPrix float64
	Status string
	Attribute string
	Created_at time.Time
}