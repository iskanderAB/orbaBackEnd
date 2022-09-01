package entities

import "time"

type Rote struct {
	IDUser     string
	Attribute  string
	IDProduct  string
	Created_at time.Time
	Rate int16
	Updated_at time.Time
}