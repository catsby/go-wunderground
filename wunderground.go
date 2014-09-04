package wunderground

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	return fmt.Sprintf("%v\n", f.Date)
}

type Response struct {
	Version string `json:"version,omitempty"`
	TOS     string `json:"termsOfService,omitempty"`
}

type ApiResponse struct {
	Response Response `json:"response"`
	Forecast struct {
		TxtForecast TxtForecast `json:"txt_forecast"`
	} `json:"forecast"`
}

func Get(query int) (*ApiResponse, error) {
	key := os.Getenv("WUNDERGROUND_API_KEY")
	if len(key) == 0 {
		log.Fatal("No API key found")
	}
	qs := fmt.Sprintf("/forecast/q/%d.json", query)
	resp, err := http.Get(API_URL + key + qs)
	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("something wrong in reading: ", err)
	}

	api_response := &ApiResponse{}
	if err := json.Unmarshal(body, &api_response); err != nil {
		log.Fatal("whoops in unmarshalling:", err)
	}

	return api_response, err
}
