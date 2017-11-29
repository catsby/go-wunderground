package wunderground

import (
	"strings"
)

// Use to request the forecast feature in Service.Request
var FForecast = "forecast"

// Use to request the forecast10day feature in Service.Request
var FForecast10Day = "forecast10day"

type Forecast struct {
	TxtForecast    `json:"txt_forecast"`
	SimpleForecast `json:"simpleforecast"`
}

type TxtForecast struct {
	Date        string `json:"date"`
	ForecastDay []struct {
		Period        int    `json:"period"`
		Icon          string `json:"icon"`
		IconURL       string `json:"icon_url"`
		Title         string `json:"title"`
		FctText       string `json:"fcttext"`
		FctTextMetric string `json:"fcttext_metric"`
		Pop           string `json:"pop"`
	} `json:"forecastday"`
}

type SimpleForecast struct {
	Forecastday []struct {
		Date struct {
			Epoch        string `json:"epoch"`
			Pretty       string `json:"pretty"`
			Day          int    `json:"day"`
			Month        int    `json:"month"`
			Year         int    `json:"year"`
			Yday         int    `json:"yday"`
			Hour         int    `json:"hour"`
			Min          string `json:"min"`
			Sec          int    `json:"sec"`
			Isdst        string `json:"isdst"`
			Monthname    string `json:"monthname"`
			WeekdayShort string `json:"weekday_short"`
			Weekday      string `json:"weekday"`
			Ampm         string `json:"ampm"`
			TzShort      string `json:"tz_short"`
			TzLong       string `json:"tz_long"`
		} `json:"date"`
		Period      int          `json:"period"`
		High        ForecastTemp `json:"high"`
		Low         ForecastTemp `json:"low"`
		Conditions  string       `json:"conditions"`
		Icon        string       `json:"icon"`
		IconURL     string       `json:"icon_url"`
		Skyicon     string       `json:"skyicon"`
		Pop         int          `json:"pop"`
		QpfAllday   ForecastQpf  `json:"qpf_allday"`
		QpfDay      ForecastQpf  `json:"qpf_day"`
		QpfNight    ForecastQpf  `json:"qpf_night"`
		SnowAllday  ForecastSnow `json:"snow_allday"`
		SnowDay     ForecastSnow `json:"snow_day"`
		SnowNight   ForecastSnow `json:"snow_night"`
		Maxwind     ForecastWind `json:"maxwind"`
		Avewind     ForecastWind `json:"avewind"`
		Avehumidity float64      `json:"avehumidity"`
		Maxhumidity float64      `json:"maxhumidity"`
		Minhumidity float64      `json:"minhumidity"`
	} `json:"forecastday"`
}

type ForecastTemp struct {
	Fahrenheit string `json:"fahrenheit"`
	Celsius    string `json:"celsius"`
}

type ForecastQpf struct {
	In float64 `json:"in"`
	Mm float64 `json:"mm"`
}

type ForecastSnow struct {
	In float64 `json:"in"`
	Cm float64 `json:"cm"`
}

type ForecastWind struct {
	Mph     float64 `json:"mph"`
	Kph     float64 `json:"kph"`
	Dir     string  `json:"dir"`
	Degrees float64 `json:"degrees"`
}

func (f *Forecast) ToString() string {
	var res []string
	res = append(res, "Issued at "+f.Date)
	for _, forecast := range f.TxtForecast.ForecastDay {
		res = append(res, forecast.Title+": "+forecast.FctText)
	}

	return strings.Join(res, "\n")
}
