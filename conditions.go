package wunderground

import (
	"fmt"
	"strconv"
	"strings"
)

// Use to request the conditions feature in Service.Request
var FConditions = "conditions"

type CurrentObservation struct {
	Image struct {
		URL   string `json:"url"`
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"image"`
	DisplayLocation     ObservationLocation `json:"display_location"`
	ObservationLocation ObservationLocation `json:"observation_location"`
	Estimated           struct {
	} `json:"estimated"`
	StationID             string  `json:"station_id"`
	ObservationTime       string  `json:"observation_time"`
	ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
	ObservationEpoch      string  `json:"observation_epoch"`
	LocalTimeRfc822       string  `json:"local_time_rfc822"`
	LocalEpoch            string  `json:"local_epoch"`
	LocalTzShort          string  `json:"local_tz_short"`
	LocalTzLong           string  `json:"local_tz_long"`
	LocalTzOffset         string  `json:"local_tz_offset"`
	Weather               string  `json:"weather"`
	TemperatureString     string  `json:"temperature_string"`
	TempF                 float64 `json:"temp_f"`
	TempC                 float64 `json:"temp_c"`
	RelativeHumidity      string  `json:"relative_humidity"`
	WindString            string  `json:"wind_string"`
	WindDir               string  `json:"wind_dir"`
	WindDegrees           float64 `json:"wind_degrees"`
	WindMph               float64 `json:"wind_mph"`
	WindGustMph           string  `json:"wind_gust_mph"`
	WindKph               float64 `json:"wind_kph"`
	WindGustKph           string  `json:"wind_gust_kph"`
	PressureMb            string  `json:"pressure_mb"`
	PressureIn            string  `json:"pressure_in"`
	PressureTrend         string  `json:"pressure_trend"`
	DewpointString        string  `json:"dewpoint_string"`
	DewpointF             float64 `json:"dewpoint_f"`
	DewpointC             float64 `json:"dewpoint_c"`
	HeatIndexString       string  `json:"heat_index_string"`
	HeatIndexF            string  `json:"heat_index_f"`
	HeatIndexC            string  `json:"heat_index_c"`
	WindchillString       string  `json:"windchill_string"`
	WindchillF            string  `json:"windchill_f"`
	WindchillC            string  `json:"windchill_c"`
	FeelslikeString       string  `json:"feelslike_string"`
	FeelslikeF            string  `json:"feelslike_f"`
	FeelslikeC            string  `json:"feelslike_c"`
	VisibilityMi          string  `json:"visibility_mi"`
	VisibilityKm          string  `json:"visibility_km"`
	Solarradiation        string  `json:"solarradiation"`
	UV                    string  `json:"UV"`
	Precip1HrString       string  `json:"precip_1hr_string"`
	Precip1HrIn           string  `json:"precip_1hr_in"`
	Precip1HrMetric       string  `json:"precip_1hr_metric"`
	PrecipTodayString     string  `json:"precip_today_string"`
	PrecipTodayIn         string  `json:"precip_today_in"`
	PrecipTodayMetric     string  `json:"precip_today_metric"`
	Icon                  string  `json:"icon"`
	IconURL               string  `json:"icon_url"`
	ForecastURL           string  `json:"forecast_url"`
	HistoryURL            string  `json:"history_url"`
	ObURL                 string  `json:"ob_url"`
}

type ObservationLocation struct {
	Full           string `json:"full"`
	City           string `json:"city"`
	State          string `json:"state"`
	StateName      string `json:"state_name"`
	Country        string `json:"country"`
	CountryIso3166 string `json:"country_iso3166"`
	Zip            string `json:"zip"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Elevation      string `json:"elevation"`
}

func (c *CurrentObservation) ToString() string {
	var res []string

	res = append(res, fmt.Sprintf("Current conditions at %s (%s)\n%s",
		c.ObservationLocation.Full, c.StationID, c.ObservationTime))
	res = append(res, "   Temperature: "+c.TemperatureString)
	if c.HeatIndexString != "NA" {
		res = append(res, "   Heat Index: "+c.HeatIndexString)
	}
	res = append(res, "   Sky Conditions: "+c.Weather)
	res = append(res, "   Wind: "+c.WindString)

	var pTrend string
	switch c.PressureTrend {
	case "+":
		pTrend = "rising"
	case "-":
		pTrend = "falling"
	case "0":
		pTrend = "holding steady"
	}
	res = append(res, fmt.Sprintf("   Pressure: %s in (%s mb) and %s", c.PressureIn, c.PressureMb, pTrend))

	res = append(res, "   Relative humidity: "+c.RelativeHumidity)

	var dpHuman string
	switch {
	case c.DewpointF < 50.0:
		dpHuman = " (dry)"
	case c.DewpointF <= 54.0:
		dpHuman = " (very comfortable)"
	case c.DewpointF <= 59.0:
		dpHuman = " (comfortable)"
	case c.DewpointF <= 64.0:
		dpHuman = " (okay for most)"
	case c.DewpointF <= 69.0:
		dpHuman = " (somewhat uncomfortable)"
	case c.DewpointF <= 74.0:
		dpHuman = " (very humid)"
	case c.DewpointF <= 80.0:
		dpHuman = " (oppressive)"
	default:
		dpHuman = " (dangerously high)"
	}
	res = append(res, "   Dewpoint: "+c.DewpointString+dpHuman)

	if c.WindchillString != "NA" {
		res = append(res, "   Windchill: "+c.WindchillString)
	}

	res = append(res, fmt.Sprintf("   Visibility: %s miles", c.VisibilityMi))

	if precIn, _ := strconv.ParseFloat(c.PrecipTodayIn, 64); precIn > 0 {
		res = append(res, "   Precipitation today: "+c.PrecipTodayString)
	}

	return strings.Join(res, "\n")
}
