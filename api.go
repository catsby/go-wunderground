package wunderground

import (
	"strconv"
	"strings"
	"time"
)

type ApiResponse struct {
	Response            `json:"response"`
	*Alerts             `json:"alerts"`
	*Almanac            `json:"almanac"`
	*CurrentObservation `json:"current_observation"`
	*Forecast           `json:"forecast"`
	*History            `json:"history"`
	*HourlyForecast     `json:"hourly_forecast"`
	*Location           `json:"location"`
	*MoonPhase          `json:"moon_phase"`
	*RawTide            `json:"rawtide"`
	*Satellite          `json:"satellite"`
	*Tide               `json:"tide"`
	*Trip               `json:"trip"`
	*WebCams            `json:"webcams"`
}

func (a *ApiResponse) ToString() string {
	var res []string
	if a.Alerts != nil {
		res = append(res, a.Alerts.ToString())
	}
	if a.Almanac != nil {
		res = append(res, a.Almanac.ToString())
	}
	if a.CurrentObservation != nil {
		res = append(res, a.CurrentObservation.ToString())
	}
	if a.Forecast != nil {
		res = append(res, a.Forecast.ToString())
	}
	if a.History != nil {
		res = append(res, a.History.ToString())
	}
	if a.Location != nil {
		res = append(res, a.Location.ToString())
	}
	if a.MoonPhase != nil {
		res = append(res, a.MoonPhase.ToString())
	}
	if a.Tide != nil {
		res = append(res, a.Tide.ToString())
	}
	if a.Trip != nil {
		res = append(res, a.Trip.ToString())
	}

	return strings.Join(res, "\n")
}

type Response struct {
	Version        string `json:"version"`
	TermsofService string `json:"termsofService"`
	Features       struct {
		Alerts        int `json:"alerts"`
		Almanac       int `json:"almanac"`
		Astronomy     int `json:"astronomy"`
		Conditions    int `json:"conditions"`
		Forecast      int `json:"forecast"`
		Forecast10Day int `json:"forecast10day"`
		GeoLookup     int `json:"geolookup"`
		History       int `json:"history"`
		Hourly        int `json:"hourly"`
		Hourly10Day   int `json:"hourly10day"`
		Planner       int `json:"planner"`
		RawTide       int `json:"rawtide"`
		Satellite     int `json:"satellite"`
		Tide          int `json:"tide"`
		WebCams       int `json:"webcams"`
		Yesterday     int `json:"yesterday"`
	} `json:"features"`
	*Error `json:"error"`
}

type Error struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Temperature struct {
	F string `json:"F"`
	C string `json:"C"`
}

type PrettyDate struct {
	Pretty string `json:"pretty"`
	Year   string `json:"year"`
	Mon    string `json:"mon"`
	MDay   string `json:"mday"`
	Hour   string `json:"hour"`
	Min    string `json:"min"`
	TzName string `json:"tzname"`
	Epoch  string `json:"epoch"`
}

func (d *PrettyDate) ToDate() (time.Time, error) {
	ye, _ := strconv.Atoi(d.Year)
	mon, _ := strconv.Atoi(d.Mon)
	mo := time.Month(mon)
	da, _ := strconv.Atoi(d.MDay)
	ho, _ := strconv.Atoi(d.Hour)
	mi, _ := strconv.Atoi(d.Min)
	loc, err := time.LoadLocation(d.TzName)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(ye, mo, da, ho, mi, 0 /*sec*/, 0 /*nsec*/, loc), nil
}
