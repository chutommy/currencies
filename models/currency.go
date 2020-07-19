package models

type Currency struct {
	Name        string
	Country     string
	Description string
	ChangeP     int
	Rate        float64
	LastUpdate  int64
}
