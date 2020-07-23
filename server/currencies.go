package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	data "github.com/chutified/currencies/data"
	currency "github.com/chutified/currencies/protos/currency"
)

// Currency is a server which serves the gRPC currency calls.
type Currency struct {
	log *log.Logger
	ds  *data.Service

	// subscriptions holds all subscriptions for each client
	subscriptions map[currency.Currency_SubscribeCurrencyServer][]*currency.GetCurrencyRequest
}

// New is a constructor of the currency server.
func New(log *log.Logger, ds *data.Service) *Currency {
	c := &Currency{
		log: log,
		ds:  ds,

		subscriptions: make(map[currency.Currency_SubscribeCurrencyServer][]*currency.GetCurrencyRequest),
	}

	// monitoring
	go func() {
		c.handleUpdates()
	}()

	return c
}

// GetCurrency handles GetCurrency gRPC request calls.
func (c *Currency) GetCurrency(ctx context.Context, req *currency.GetCurrencyRequest) (*currency.GetCurrencyResponse, error) {

	// handle
	resp, err := c.handleGetCurrencyRequest(req)
	if err != nil {
		return nil, err
	}
	// success
	c.log.Printf("[handle] GetCurrency call: %s", req.GetName())
	return resp, nil
}

// GetRate handles GetRate gRPC request calls.
func (c *Currency) GetRate(ctx context.Context, req *currency.GetRateRequest) (*currency.GetRateResponse, error) {

	// handle
	resp, err := c.handleGetRateRequest(req)
	if err != nil {
		return nil, err
	}
	// success
	c.log.Printf("[handle] GetRate call, base: %s, destination: %s", req.GetBase(), req.GetDestination())
	return resp, nil
}

// SubscribeCurrency handles SubscribeCurrency gRPC calls.
func (c *Currency) SubscribeCurrency(srv currency.Currency_SubscribeCurrencyServer) error {
	// handle requests
	for {

		// get request
		req, err := srv.Recv()
		if err == io.EOF {
			c.log.Printf("[exit] client closed SubscribeCurrency call")
			break
		}
		if err != nil {
			c.log.Printf("[error] invalid request format: %v", err)
			return fmt.Errorf("invalid request: %w", err)
		}
		name := req.GetName()

		// validate request
		if _, ok := c.ds.Currencies[name]; !ok {
			c.log.Printf("[error] currency %v not found", name)

			// TODO handle error the grpc way
			continue
		}

		// create client folder if does not exit
		reqs, ok := c.subscriptions[srv]
		if !ok {
			c.subscriptions[srv] = []*currency.GetCurrencyRequest{}
		}

		// check duplicates
		var validErr error
		for _, r := range reqs {
			if r.GetName() == name {
				c.log.Printf("[error] the client has been already subscribed to %s", r.GetName())

				// TODO handle error the grpc way
				validErr = errors.New("valid err")
				break
			}
		}
		if validErr != nil {

			// TODO handle error the grpc way
			continue
		}

		// success, append
		c.log.Printf("[success] client successfully subscribed to: %s", name)
		c.subscriptions[srv] = append(c.subscriptions[srv], req)
	}

	// exit
	return nil
}

func (c *Currency) handleUpdates() {

	updates, errs := c.ds.MonitorData(15 * time.Second)

	// log errors
	go func() {
		for err := range errs {
			c.log.Printf("[error] monitor data: %s", err)
		}
	}()

	// update
	for range updates {
		c.log.Printf("[update] currency data updated")

		// // iterate over clients
		// for _, clients := range c.subscriptions {

		//     // iterate over subscriptions
		//     for _, req := range clients {

		//     }
		// }
	}
}
