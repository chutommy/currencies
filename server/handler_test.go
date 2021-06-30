package server

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"

	data "github.com/chutommy/currencies/data"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandleUpdates(t *testing.T) {
	ss := bytes.NewBufferString("")
	log := log.New(ss, "", log.LstdFlags)
	ds := data.New()
	err := ds.Update("https://markets.businessinsider.com/currencies")
	if err != nil {
		t.Fatal("could not update data")
	}
	c := &Currency{
		log: log,
		ds:  ds,
	}

	*ss = *(bytes.NewBufferString(""))
	t.Run("ok", func(t1 *testing.T) {
		// TODO add subscribtions

		go func() {
			c.handleUpdates(500*time.Millisecond, "https://markets.businessinsider.com/currencies")
		}()
		time.Sleep(1 * time.Second)

		assert.Equal(t1, strings.Contains(ss.String(), "[update]"), true)
	})

	*ss = *(bytes.NewBufferString(""))
	t.Run("invalid url", func(t1 *testing.T) {
		c := &Currency{
			log: log,
			ds:  ds,
		}

		go func() {
			c.handleUpdates(1*time.Millisecond, "invalid url")
		}()
		time.Sleep(1 * time.Second)

		assert.Equal(t1, strings.Contains(ss.String(), "[error]"), true)
	})

}
