package entities

import "time"

type Options struct {
	ID          string
	Title       string
	Description string
	Prix        float64
	Time        time.Time
	
}