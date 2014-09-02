package wunderground

import "fmt"

// const API_URL = "http://localhost:4567/"
const API_URL = "https://api.wunderground.com/api/"

type ForecastDay struct {
	Period  int    `json:"period,omitempty"`
	Title   string `json:"title,omitempty"`
	FctText string `json:"fcttext,omitempty"`
}

type TxtForecast struct {
	Date string        `json:"date,omitempty"`
	Days []ForecastDay `json:"forecastday,omitempty"`
}

func (f *TxtForecast) String() string {
	fmt.Println("welp...")
	return fmt.Sprintf("%v\n", f.Date)
}

type Forecast struct {
	TxtForecast TxtForecast `json:"txt_forecast"`
}

type Response struct {
	Version string `json:"version,omitempty"`
	TOS     string `json:"termsOfService,omitempty"`
}

type ApiResponse struct {
	Response Response `json:"response"`
	Forecast Forecast `json:"forecast"`
}
