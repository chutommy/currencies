package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chutified/currencies/models"
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

	// prepare cache memory
	var records []record
	// loop over rows
	doc.Find(".row-hover").Each(func(i int, s *goquery.Selection) {
		children := s.Children()

		// select
		name := children.First()
		ctr := name.Next()
		dsc := ctr.Next().Next()
		chP := dsc.Next()
		prc := chP.Next()
		last := prc.Next()

		// filter
		name = name.Children().First()
		dsc = dsc.Children().First()
		chP = chP.Children().First()

		// trim spaces
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

func parseRecords(rs []record) (map[string]models.Currency, error) {

	// prepare cache memory
	var cs map[string]models.Currency
	for _, r := range rs {

		// name
		var name string
		_, err := fmt.Sscanf(r.name, "USD/%s", &name)
		if err != nil {
			return nil, fmt.Errorf("unexpected name format %s", err)
		}

		// country
		country := r.country

		// description
		description := r.description

		// changeP
		strconv.ParseFloat(r.changeP, 32)
	}

	return cs, nil
}
