package server

import (
	"fmt"
	"strings"

	"github.com/chutified/currencies/protos/currency"
)

// handleGetCurrencyRequest handles request for the GetCurrecy calls.
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
