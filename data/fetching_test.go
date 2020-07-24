package data

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestParseRecords(t *testing.T) {

	reset := []record{{
		name:        "USD/USD",
		country:     "United States of America",
		description: "United States Dollar",
		changeP:     "0",
		rate:        "1",
		lastUpdate:  "01/02/2006 03:04:05 PM UTC-0400",
	}}

	tests := []struct {
		name   string
		action func([]record)
		err    string
	}{
		{
			name:   "ok",
			action: func(rs []record) {},
			err:    "",
		},
		{
			name:   "invalid name",
			action: func(rs []record) { rs[0].name = "invalid" },
			err:    "unexpected name format",
		},
		{
			name:   "invalid percentages",
			action: func(rs []record) { rs[0].changeP = "invalid" },
			err:    "unexpected change in percentage",
		},
		{
			name:   "invalid rate",
			action: func(rs []record) { rs[0].rate = "invalid" },
			err:    "unexpected currency rate",
		},
		{
			name:   "invalid time",
			action: func(rs []record) { rs[0].lastUpdate = "invalid" },
			err:    "unexpected time layout",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			// clone the slice
			rs := make([]record, len(reset))
			copy(rs, reset)

			// run the action
			test.action(rs)

			_, err := parseRecords(rs)
			if err != nil {

				// check error message
				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
			}
		})
	}
}
