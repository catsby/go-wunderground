package wunderground

import (
	"fmt"
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
	Forecastday []struct {
		Period        int    `json:"period"`
		Icon          string `json:"icon"`
		IconURL       string `json:"icon_url"`
		Title         string `json:"title"`
		Fcttext       string `json:"fcttext"`
		FcttextMetric string `json:"fcttext_metric"`
		Pop           string `json:"pop"`
	} `json:"forecastday"`
}

func (f *TxtForecast) String() string {
	return fmt.Sprintf("%v\n", f.Date)
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
		Period int `json:"period"`
		High   struct {
			Fahrenheit string `json:"fahrenheit"`
			Celsius    string `json:"celsius"`
		} `json:"high"`
		Low struct {
			Fahrenheit string `json:"fahrenheit"`
			Celsius    string `json:"celsius"`
		} `json:"low"`
		Conditions string `json:"conditions"`
		Icon       string `json:"icon"`
		IconURL    string `json:"icon_url"`
		Skyicon    string `json:"skyicon"`
		Pop        int    `json:"pop"`
		QpfAllday  struct {
			In float64 `json:"in"`
			Mm float64 `json:"mm"`
		} `json:"qpf_allday"`
		QpfDay struct {
			In float64 `json:"in"`
			Mm float64 `json:"mm"`
		} `json:"qpf_day"`
		QpfNight struct {
			In float64 `json:"in"`
			Mm float64 `json:"mm"`
		} `json:"qpf_night"`
		SnowAllday struct {
			In float64 `json:"in"`
			Cm float64 `json:"cm"`
		} `json:"snow_allday"`
		SnowDay struct {
			In float64 `json:"in"`
			Cm float64 `json:"cm"`
		} `json:"snow_day"`
		SnowNight struct {
			In float64 `json:"in"`
			Cm float64 `json:"cm"`
		} `json:"snow_night"`
		Maxwind struct {
			Mph     float64 `json:"mph"`
			Kph     float64 `json:"kph"`
			Dir     string  `json:"dir"`
			Degrees float64 `json:"degrees"`
		} `json:"maxwind"`
		Avewind struct {
			Mph     float64 `json:"mph"`
			Kph     float64 `json:"kph"`
			Dir     string  `json:"dir"`
			Degrees float64 `json:"degrees"`
		} `json:"avewind"`
		Avehumidity float64 `json:"avehumidity"`
		Maxhumidity float64 `json:"maxhumidity"`
		Minhumidity float64 `json:"minhumidity"`
	} `json:"forecastday"`
}
