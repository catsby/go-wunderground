package wunderground

import (
	"testing"
)

func TestQueryString(t *testing.T) {
	var tests = []struct {
		q        *Query
		expected string
	}{
		{&Query{USState: "CA", City: "San Francisco"}, "CA/San_Francisco.json"},
		{&Query{USZip: "60290"}, "60290.json"},
		{&Query{Country: "Australia", City: "Wagga Wagga"}, "Australia/Wagga_Wagga.json"},
		{&Query{Lat: "38.8", Long: "-122.4"}, "38.8,-122.4.json"},
		{&Query{Airport: "KJFK"}, "KJFK.json"},
		{&Query{PWS: "KCASANFR70"}, "pws:KCASANFR70.json"},
		{&Query{Auto: true}, "autoip.json"},
		{&Query{IP: "38.102.136.138"}, "autoip.json?geo_ip=38.102.136.138"},
		{&Query{Auto: true, IP: "38.102.136.138"}, "autoip.json?geo_ip=38.102.136.138"},
		{&Query{User: "cats and dogs"}, "cats and dogs.json"},
	}

	for _, test := range tests {
		actual := test.q.String()
		if actual != test.expected {
			t.Errorf("%+v, actual=%s, expected=%s", *test.q, actual, test.expected)
		}
	}
}
