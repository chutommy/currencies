package data

import (
	"testing"

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
