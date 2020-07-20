package data

import (
	"fmt"
	"reflect"
	"time"

	"github.com/chutified/currencies/models"
)

// Service defines the data fetched from the website source.
type Service struct {
	Currencies map[string]*models.Currency
}

// New constructs a new data service.
func New() *Service {
	return &Service{
		Currencies: make(map[string]*models.Currency),
	}
}

// GetCurrency retieves data from the cache memory.
func (ds *Service) GetCurrency(name string) (*models.Currency, error) {

	// search
	c, ok := ds.Currencies[name]
	if !ok {
		return nil, fmt.Errorf("currency '%s' not found", name)
	}

	// success
	return c, nil
}

// Update updates the Currencies data.
func (ds *Service) Update() error {

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
	ds.Currencies = m
	return nil
}

// MonitorData monitors the updates and whenever new values are pulled, it sends a signal through the channel.
func (ds *Service) MonitorData(interval time.Duration) (chan struct{}, chan error) {

	// prepare channels
	update := make(chan struct{})
	errs := make(chan error)

	go func() {
		ticker := time.Tick(interval)
		for range ticker {

			// prepare maps
			cache := ds.Currencies
			err := ds.Update()
			if err != nil {
				errs <- fmt.Errorf("update currencies: %w", err)
				continue
			}

			// compare
			if !reflect.DeepEqual(ds.Currencies, cache) {

				// update
				update <- struct{}{}
			}
		}
	}()

	return update, errs
}
