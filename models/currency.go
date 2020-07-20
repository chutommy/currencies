package models

import "time"

// Currency is a base structure of the currency.
type Currency struct {
	Name        string
	Country     string
	Description string
	ChangeP     float32
	Rate        float32
	LastUpdate  time.Time
}
