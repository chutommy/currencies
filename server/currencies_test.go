package server

import (
	"bytes"
	"context"
	"log"
	"testing"

	data "github.com/chutommy/currencies/data"
	currency "github.com/chutommy/currencies/protos/currency"
	"gopkg.in/go-playground/assert.v1"
)

func TestCurrency(t *testing.T) {

	log := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	ds := data.New()
	err := ds.Update("https://markets.businessinsider.com/currencies")
	if err != nil {
		t.Fatal("could not update data")
	}

	// >>> New()
	c := New(log, ds)

	// >>> GetCurrency()
	t.Run("GetCurrency", func(t1 *testing.T) {

		t1.Run("ok", func(t2 *testing.T) {

			req := &currency.GetCurrencyRequest{
				Name: "USD",
			}
			_, err := c.GetCurrency(context.Background(), req)
			assert.Equal(t2, err, nil)
		})

		t1.Run("invalid currency", func(t2 *testing.T) {

			req := &currency.GetCurrencyRequest{
				Name: "invalid",
			}
			_, err := c.GetCurrency(context.Background(), req)
			assert.NotEqual(t2, err, nil)
		})
	})

	// >>> GetRate()
	t.Run("GetRate", func(t1 *testing.T) {

		t1.Run("ok", func(t2 *testing.T) {

			req := &currency.GetRateRequest{
				Base:        "USD",
				Destination: "EUR",
			}
			_, err := c.GetRate(context.Background(), req)
			assert.Equal(t2, err, nil)
		})

		t1.Run("invalid base currency", func(t2 *testing.T) {

			req := &currency.GetRateRequest{
				Base:        "invalid",
				Destination: "EUR",
			}
			_, err := c.GetRate(context.Background(), req)
			assert.NotEqual(t2, err, nil)
		})

		t1.Run("invalid destination currency", func(t2 *testing.T) {

			req := &currency.GetRateRequest{
				Base:        "USD",
				Destination: "invalid",
			}
			_, err := c.GetRate(context.Background(), req)
			assert.NotEqual(t2, err, nil)
		})
	})

	// >>> SubscribeCurrency
	// TODO
}
