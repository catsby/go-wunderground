package wunderground

import (
	"flag"
	"testing"
	"time"
)

var testApiKey = flag.String("test_api_key", "", "If set, enable tests that contact the live servers")

func TestFullRequest(t *testing.T) {
	if len(*testApiKey) == 0 {
		t.Skip("skipping live API tests")
	}

	client := NewDevLimitedService(*testApiKey)

	// Plan and history don't mix. The first request intentionally includes this error, and the next checks the Planner parsing.
	planStart := time.Date(2017, time.July, 1, 0, 0, 0, 0, time.UTC)
	planEnd := time.Date(2017, time.July, 31, 0, 0, 0, 0, time.UTC)

	ar, err := client.Request([]string{ FAlerts, FAlmanac, FAstronomy, FConditions, FForecast, FGeoLookup, FYesterday, FHourly, FPlanner(planStart, planEnd), FRawTide, FSatellite, FTide, FWebCams }, &Query{USState: "CA", City: "San Francisco"})
	if err != nil {
		t.Errorf("Failed to execute query: %s", err)
	}
	if ar.Alerts == nil {
		t.Errorf("failed to locate alerts")
	}
	if ar.Almanac == nil {
		t.Errorf("failed to locate almanac")
	}
	if ar.CurrentObservation == nil {
		t.Errorf("failed to locate observation")
	}
	if ar.Forecast == nil {
		t.Errorf("failed to locate forecast")
	}
	if ar.History == nil {
		t.Errorf("failed to locate history")
	}
	if ar.HourlyForecast == nil {
		t.Errorf("failed to locate hourly forecast")
	}
	if ar.Location == nil {
		t.Errorf("failed to locate location")
	}
	if ar.MoonPhase == nil {
		t.Errorf("failed to locate moon phase")
	}
	if ar.RawTide == nil {
		t.Errorf("failed to locate raw tide")
	}
	if ar.Satellite == nil {
		t.Errorf("failed to locate satellite")
	}
	if ar.Tide == nil {
		t.Errorf("failed to locate tide")
	}
	if ar.WebCams == nil {
		t.Errorf("failed to locate webcams")
	}

	ar, err = client.Request([]string{ FPlanner(planStart, planEnd) }, &Query{USState: "CA", City: "San Francisco"})
	if err != nil {
		t.Errorf("Failed to execute query: %s", err)
	}
	if ar.Trip == nil {
		t.Errorf("failed to locate trip")
	}
}
