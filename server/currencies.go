package server

import (
	"log"
	"time"

	"github.com/chutified/currencies/data"
)

type Currency struct {
	log *log.Logger
	ds  *data.Service
}

func New(log *log.Logger, ds *data.Service) *Currency {
	c := &Currency{
		log: log,
		ds:  ds,
	}

	return c
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

		// TODO
		// data got updated
	}
}
