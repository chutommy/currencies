package server

import (
	"fmt"
	"strings"

	"github.com/chutified/currencies/protos/currency"
)

func (c *Currency) handleGetCurrencyRequest(req *currency.GetCurrencyRequest) (*currency.GetCurrencyResponse, error) {

	// get name
	name := req.GetName()
	name = strings.ToUpper(name)

	// retrive data
	crc, err := c.ds.GetCurrency(name)
	if err != nil {
		return nil, fmt.Errorf("handling GetCurrency call: %w", err)
	}

	// success
	resp := &currency.GetCurrencyResponse{
		Name:        crc.Name,
		Country:     crc.Country,
		Description: crc.Description,
		Change:      crc.ChangeP,
		RateUSD:     crc.Rate,
		UpdatedAt:   crc.LastUpdate.String(),
	}
	return resp, nil
}

func (c *Currency) handleGetRateRequest(req *currency.GetRateRequest) (*currency.GetRateResponse, error) {

	// get values
	base := req.GetBase()
	dest := req.GetDestination()
	base = strings.ToUpper(base)
	dest = strings.ToUpper(dest)

	// retrieve data
	rate, err := c.ds.GetRate(base, dest)
	if err != nil {
		return nil, fmt.Errorf("call GetRate: %w", err)
	}

	// success
	resp := &currency.GetRateResponse{
		Rate: rate,
	}
	return resp, nil
}
