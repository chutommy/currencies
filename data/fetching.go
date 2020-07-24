package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	models "github.com/chutified/currencies/models"
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
func fetchData(url string) ([]record, error) {

	// fetch the date
	doc, err := goquery.NewDocument(url)
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
		dsc := ctr.Next()
		chP := dsc.Next().Next()
		prc := chP.Next()
		last := prc.Next()

		// filter
		name = name.Children().First()
		dsc = dsc.Children().First()
		chP = chP.Children().First()

		lastTemp := last.Children().Last()
		if lastTemp.Text() != "" {
			last = lastTemp
		}

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

// parseRecords construct a map of currencies from je slice of the records.
func parseRecords(rs []record) (map[string]*models.Currency, error) {

	// prepare cache memory
	cs := make(map[string]*models.Currency)

	// add base
	cs["USD"] = &models.Currency{
		Name:        "USD",
		Country:     "United States of America",
		Description: "United States Dollar",
		ChangeP:     0,
		Rate:        1,
		LastUpdate:  time.Now(),
	}

	for _, r := range rs {

		// name
		var name string
		_, err := fmt.Sscanf(r.name, "USD/%s", &name)
		if err != nil {
			return nil, fmt.Errorf("unexpected name format: %w", err)
		}
		name = strings.ToUpper(name)

		// country
		country := r.country

		// description
		description := r.description

		// changeP
		chP := strings.ReplaceAll(r.changeP, ",", "")
		changeP, err := strconv.ParseFloat(chP, 32)
		if err != nil {
			return nil, fmt.Errorf("unexpected change in percentage: %w", err)
		}

		// rate
		p := strings.ReplaceAll(r.rate, ",", "")
		rate, err := strconv.ParseFloat(p, 32)
		if err != nil {
			return nil, fmt.Errorf("unexpected currency rate: %w", err)
		}

		// prepare layouts
		const timeLayout = "01/02/2006 03:04:05 PM UTC-0400"
		const dateLayout = "1/2/2006"
		// time
		lastUpdate, err := time.Parse(timeLayout, r.lastUpdate)
		if err != nil {

			// date
			lastUpdate, err = time.Parse(dateLayout, r.lastUpdate)
			if err != nil {
				return nil, fmt.Errorf("unexpected time layout: %w", err)
			}
		}

		// construct
		item := &models.Currency{
			Name:        name,
			Country:     country,
			Description: description,
			ChangeP:     float32(changeP),
			Rate:        float32(rate),
			LastUpdate:  lastUpdate,
		}
		cs[item.Name] = item
	}

	return cs, nil
}
