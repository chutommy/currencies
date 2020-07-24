package data

import (
	"fmt"
	"math"
	"reflect"
	"time"

	models "github.com/chutified/currencies/models"
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

// GetCurrency retieves data structure from the cache memory.
func (ds *Service) GetCurrency(name string) (*models.Currency, error) {

	// search
	c, ok := ds.Currencies[name]
	if !ok {
		return nil, fmt.Errorf("currency '%s' not found", name)
	}

	// success
	return c, nil
}

// GetRate returns the exchange rate between base and destination currency.
func (ds *Service) GetRate(base string, dest string) (float32, error) {

	// search
	b, ok := ds.Currencies[base]
	if !ok {
		return 0, fmt.Errorf("base currency '%s' not found", base)
	}
	d, ok := ds.Currencies[dest]
	if !ok {
		return 0, fmt.Errorf("destination currency '%s' not found", dest)
	}

	// result
	rate := d.Rate / b.Rate
	// round
	rate = float32(math.Round(float64(rate*10000)) / 10000)
	return rate, nil
}

// Update updates the Currencies data.
func (ds *Service) Update(url string) error {

	// get data
	rs, err := fetchData(url)
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
func (ds *Service) MonitorData(interval time.Duration, url string) (chan struct{}, chan error) {

	// prepare channels
	update := make(chan struct{})
	errs := make(chan error)

	go func() {
		ticker := time.Tick(interval)
		for range ticker {

			// prepare maps
			cache := ds.Currencies
			err := ds.Update(url)
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
