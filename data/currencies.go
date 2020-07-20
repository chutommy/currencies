package data

import (
	"fmt"

	"github.com/chutified/currencies/models"
)

// DataService defines the data fetched from the website source.
type DataService struct {
	currencies map[string]*models.Currency
}

// New constructs a new data service.
func New() *DataService {
	return &DataService{
		currencies: make(map[string]*models.Currency),
	}
}

// Update updates the currencies data.
func (ds *DataService) Update() error {

	// get data
	rs, err := fetchData()
	if err != nil {
		return fmt.Errorf("fetching data: %w", err)
	}
	m, err := parseRecords(rs)
	if err != nil {
		return fmt.Errorf("parsing records: %w", err)
	}

	// success
	ds.currencies = m
	return nil
}
