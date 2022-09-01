package entities

import "time"

type Product struct {
	ID          string
	Title       string
	Description string
	Prix        float32
	ProdTime        time.Time
}