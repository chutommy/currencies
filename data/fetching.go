package data

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// record is a cache structure which holds raw source data.
type record struct {
	name        string
	country     string
	description string
	changeP     string
	rate        string
	lastUpdate  string
}

// fetchData crawls the source website and returns the latest raw data.
func fetchData() ([]record, error) {

	// fetch the date
	doc, err := goquery.NewDocument("https://markets.businessinsider.com/currencies")
	if err != nil {
		return nil, fmt.Errorf("fetch source website: %w", err)
	}

	// prepare saving memory
	var records []record
	// loop over rows
	doc.Find(".row-hover").Each(func(i int, s *goquery.Selection) {
		children := s.Children()

		name := children.First()
		ctr := name.Next()
		dsc := ctr.Next().Next()
		chP := dsc.Next()
		prc := chP.Next()
		last := prc.Next()

		name = name.Children().First()
		dsc = dsc.Children().First()
		chP = chP.Children().First()

		n := strings.TrimSpace(name.Text())
		c := strings.TrimSpace(ctr.Text())
		d := strings.TrimSpace(dsc.Text())
		ch := strings.TrimSpace(chP.Text())
		p := strings.TrimSpace(prc.Text())
		l := strings.TrimSpace(last.Text())

		// add to the list
		r := record{
			name:        n,
			country:     c,
			description: d,
			changeP:     ch,
			rate:        p,
			lastUpdate:  l,
		}
		records = append(records, r)
	})

	return records, nil
}
