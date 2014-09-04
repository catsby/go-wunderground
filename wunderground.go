package wunderground

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

type ConditionsResponse struct {
	Response            Response            `json:"response"`
	CurrentObservastion CurrentObservastion `json:"current_observation"`
}

type CurrentObservastion struct {
	DisplayLocation struct {
		Full string `json:"full"`
	} `json:"display_location"`
}

func (co *ConditionsResponse) LocationName() string {
	return co.CurrentObservastion.DisplayLocation.Full
}

// Service represents your API.
type Service struct {
	ApiKey string
	client *http.Client
}

// NewService creates a Service using the given, if none is provided
// it uses http.DefaultClient.
func NewService(key string) *Service {
	return &Service{
		ApiKey: key,
		client: http.DefaultClient,
	}
}
func (c *Service) Forecast(query interface{}) (*ApiResponse, error) {
	// TODO: This doesn't belong here.
	// Should seperate this into a client/config struct, more like
	// cyberdelia/heroku-go or something.
	qs := fmt.Sprintf("/forecast/q/%d.json", query)
	resp, err := c.client.Get(API_URL + c.ApiKey + qs)
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

func (c *Service) Conditions(query interface{}) (*ConditionsResponse, error) {
	qs := fmt.Sprintf("/conditions/q/%d.json", query)
	resp, err := c.client.Get(API_URL + c.ApiKey + qs)
	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("something wrong in reading: ", err)
	}

	cr := &ConditionsResponse{}
	if err := json.Unmarshal(body, &cr); err != nil {
		log.Fatal("whoops in unmarshalling:", err)
	}

	return cr, err
}
