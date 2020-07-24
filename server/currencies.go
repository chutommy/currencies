package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	data "github.com/chutified/currencies/data"
	currency "github.com/chutified/currencies/protos/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		c.handleUpdates(6*time.Second, "https://markets.businessinsider.com/currencies")
	}()

	return c
}

// GetCurrency handles GetCurrency gRPC request calls.
func (c *Currency) GetCurrency(ctx context.Context, req *currency.GetCurrencyRequest) (*currency.GetCurrencyResponse, error) {

	// handle
	resp, err := c.handleGetCurrencyRequest(req)
	if err != nil {
		c.log.Printf("[error] handle error: %v", err)

		// defines gRPC error
		gErr := status.Newf(
			codes.NotFound,
			"Currency \"%s\" was not found.", req.GetName(),
		)

		return nil, gErr.Err()
	}
	// success
	c.log.Printf("[handle] GetCurrency call: %s", strings.ToUpper(req.GetName()))
	return resp, nil
}

// GetRate handles GetRate gRPC request calls.
func (c *Currency) GetRate(ctx context.Context, req *currency.GetRateRequest) (*currency.GetRateResponse, error) {

	// handle
	resp, err := c.handleGetRateRequest(req)
	if err != nil {
		c.log.Printf("[error] handle error: %v", err)

		// defines gRPC error
		gErr := status.Newf(
			codes.NotFound,
			"Currency was not found: %v.", err,
		)

		return nil, gErr.Err()
	}

	// success
	c.log.Printf("[handle] GetRate call, base: %s, destination: %s", strings.ToUpper(req.GetBase()), strings.ToUpper(req.GetDestination()))
	return resp, nil
}

// SubscribeCurrency handles SubscribeCurrency gRPC calls.
func (c *Currency) SubscribeCurrency(srv currency.Currency_SubscribeCurrencyServer) error {
	// handle requests
	for {

		// get request
		req, err := srv.Recv()
		if err == io.EOF {

			// cancel all subscriptions
			delete(c.subscriptions, srv)
			fmt.Println(c.subscriptions[srv])

			c.log.Printf("[exit] client closed SubscribeCurrency call")
			return nil
		}
		if err != nil {

			// cancel all subscriptions
			delete(c.subscriptions, srv)

			c.log.Printf("[error] invalid request format: %v", err)
			return fmt.Errorf("invalid request: %w", err)
		}
		name := strings.ToUpper(req.GetName())

		// validate request
		if _, ok := c.ds.Currencies[name]; !ok {
			c.log.Printf("[error] currency \"%s\" not found", name)

			// defines gRPC error
			gErr := status.Newf(
				codes.NotFound,
				"Currency \"%s\" was not found.", name,
			)

			// send error message
			err = srv.Send(&currency.StreamingSubscribeResponse{
				Message: &currency.StreamingSubscribeResponse_Error{
					Error: gErr.Proto(),
				},
			})
			if err != nil {
				c.log.Printf("[error] failed to send error message: %v", err)
			}

			continue
		}

		// create client folder if does not exit
		reqs, ok := c.subscriptions[srv]
		if !ok {
			c.subscriptions[srv] = []*currency.GetCurrencyRequest{}
		}

		// check duplicates
		var validErr *status.Status
		for _, r := range reqs {
			rname := strings.ToUpper(r.GetName())
			if rname == name {
				c.log.Printf("[error] the client has been already subscribed to %s", rname)

				// defines gRPC error
				validErr = status.Newf(
					codes.Canceled,
					"This client has already subscribed to the \"%s\".", name,
				)
				break
			}
		}
		// if validErr exists returns error and continue
		if validErr != nil {

			// send error message
			err := srv.Send(&currency.StreamingSubscribeResponse{
				Message: &currency.StreamingSubscribeResponse_Error{
					Error: validErr.Proto(),
				},
			})
			if err != nil {
				c.log.Printf("[error] failed to send response: %v", err)
			}

			continue
		}

		// success, append
		c.log.Printf("[success] client successfully subscribed to: %s", name)
		c.subscriptions[srv] = append(c.subscriptions[srv], req)
	}
}
