package wunderground

import (
  "fmt"
  "log"
)

type ApiResponse struct {
	Response Response `json:"response"`
	Forecast struct {
		TxtForecast TxtForecast `json:"txt_forecast"`
	} `json:"forecast"`
}

type TxtForecast struct {
	Date string        `json:"date,omitempty"`
	Days []ForecastDay `json:"forecastday,omitempty"`
}

func (f *TxtForecast) String() string {
	return fmt.Sprintf("%v\n", f.Date)
}

type ForecastDay struct {
	Period  int    `json:"period,omitempty"`
	Title   string `json:"title,omitempty"`
	FctText string `json:"fcttext,omitempty"`
}

func (c *Service) Forecast(query interface{}) (*ApiResponse, error) {
	ar := &ApiResponse{}
	err := c.request("forecast", query, ar)

	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	return ar, err
}

