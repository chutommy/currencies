package data

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestUpdate(t *testing.T) {

	// >>> New()
	s := New()

	// >>> Update()
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
