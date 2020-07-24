package data

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestUpdate(t *testing.T) {

	// >>> New()
	s := New()

	t.Run("ok", func(t1 *testing.T) {
		err := s.Update("https://markets.businessinsider.com/currencies")
		assert.Equal(t1, err, nil)
	})

	t.Run("invalid URL", func(t1 *testing.T) {
		err := s.Update("invalid_URL")
		assert.MatchRegex(t1, err.Error(), ".*fetching data.*")
	})

	t.Run("invalid source content", func(t1 *testing.T) {
		err := s.Update("https://markets.businessinsider.com/commodities")
		assert.MatchRegex(t1, err.Error(), ".*parsing records.*")
	})
}

func TestMonitorData(t *testing.T) {
	s := New()

	t.Run("ok", func(t1 *testing.T) {
		upds, errs := s.MonitorData(500*time.Millisecond, "https://markets.businessinsider.com/currencies")

		var err error
		go func() {
			err = <-errs
		}()
		go func() {
			_ = <-upds
		}()
		time.Sleep(1 * time.Second)

		assert.Equal(t1, err, nil)
	})

	t.Run("invalid URL", func(t1 *testing.T) {
		upds, errs := s.MonitorData(500*time.Millisecond, "invalid_URL")

		var err error
		go func() {
			err = <-errs
		}()
		go func() {
			_ = <-upds
		}()
		time.Sleep(1 * time.Second)

		assert.NotEqual(t1, err, nil)
	})
}

func TestGetCurrency(t *testing.T) {
	s := New()
	err := s.Update("https://markets.businessinsider.com/currencies")
	if err != nil {
		t.Fatal("could not update data")
	}

	t.Run("ok", func(t1 *testing.T) {
		_, err = s.GetCurrency("USD")
		assert.Equal(t1, err, nil)
	})

	t.Run("invalid currency", func(t1 *testing.T) {
		_, err = s.GetCurrency("invalid")
		assert.Equal(t1, err.Error(), "currency 'invalid' not found")
	})
}

func TestGetRate(t *testing.T) {
	s := New()
	err := s.Update("https://markets.businessinsider.com/currencies")
	if err != nil {
		t.Fatal("could not update data")
	}

	t.Run("ok", func(t1 *testing.T) {
		_, err = s.GetRate("EUR", "USD")
		assert.Equal(t1, err, nil)
	})

	t.Run("invalid base", func(t1 *testing.T) {
		_, err = s.GetRate("invalid", "USD")
		assert.Equal(t1, err.Error(), "base currency 'invalid' not found")
	})

	t.Run("invalid destination", func(t1 *testing.T) {
		_, err = s.GetRate("EUR", "invalid")
		assert.Equal(t1, err.Error(), "destination currency 'invalid' not found")
	})
}
