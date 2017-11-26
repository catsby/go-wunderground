package wunderground

import (
	"fmt"
	"strings"
)

type Query struct {
	// A United States location must set one of City/USState or five-digit USZip.
	USState string
	USZip   string

	// A non-United States location
	// A country name. If using a country code, see
	// https://www.wunderground.com/weather/api/d/docs?d=resources/country-to-iso-matching
	Country string

	City string

	Lat, Long string

	Airport string

	// A Personal Weather Station ID (e.g., KCASANFR70)
	PWS string

	// Automatically determine the location
	Auto bool

	// Use the location of the given IP
	IP string

	// A user-specified query string (e.g., "CA/San_Francisco").
	User string
}

// Return a string suitable for the Wunderground query API. Assumes JSON format.
func (q *Query) String() string {
	switch {
	case len(q.USState) > 0:
		return fmt.Sprintf("%s/%s.json",
			q.USState, strings.Replace(q.City, " ", "_", -1))
	case len(q.USZip) > 0:
		return q.USZip + ".json"
	case len(q.Country) > 0:
		return fmt.Sprintf("%s/%s.json",
			q.Country, strings.Replace(q.City, " ", "_", -1))
	case len(q.Lat) > 0:
		return fmt.Sprintf("%s,%s.json", q.Lat, q.Long)
	case len(q.Airport) > 0:
		return q.Airport + ".json"
	case len(q.PWS) > 0:
		return "pws:" + q.PWS + ".json"
	case len(q.IP) > 0:
		return "autoip.json?geo_ip=38.102.136.138"
	case q.Auto:
		return "autoip.json"
	case len(q.User) > 0:
		return q.User + ".json"
	default:
		return ""
	}
}

type Wunderground interface {
	Conditions(query *Query) (*ConditionsResponse, error)
	Forecast(query *Query) (*ForecastResponse, error)
}
