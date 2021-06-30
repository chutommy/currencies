package server

import (
	"fmt"
	"strings"
	"time"

	currency "github.com/chutommy/currencies/protos/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (c *Currency) handleUpdates(interval time.Duration, url string) {

	updates, errs := c.ds.MonitorData(interval, url)

	// log errors
	go func() {
		for err := range errs {
			c.log.Printf("[error] monitor data: %s", err)
		}
	}()

	// update
	for range updates {
		c.log.Printf("[update] currency data updated")

		// iterate over clients
		for srv, reqs := range c.subscriptions {

			// iterate over subscriptions
			for _, req := range reqs {

				// handle request
				resp, err := c.handleGetCurrencyRequest(req)
				if err != nil {
					c.log.Printf("[error] unexpected request handle error: %v", err)

					// defines gRPC error
					gErr := status.Newf(
						codes.NotFound,
						"Currency \"%s\" was not found.", req.GetName(),
					)

					// send error message
					err = srv.Send(&currency.StreamingSubscribeResponse{
						Message: &currency.StreamingSubscribeResponse_Error{
							Error: gErr.Proto(),
						},
					})
					if err != nil {
						c.log.Printf("[error] failed to send response: %v", err)
					}

					continue
				}

				// send reponse
				err = srv.Send(&currency.StreamingSubscribeResponse{
					Message: &currency.StreamingSubscribeResponse_GetCurrencyResponse{
						GetCurrencyResponse: resp,
					},
				})
				if err != nil {
					c.log.Printf("[error] failed to send response: %v", err)

					continue
				}
			}
		}
	}
}
