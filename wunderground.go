package wunderground

import (
	"encoding/json"
	"fmt"
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
	ar := &ApiResponse{}
	err := c.request("forecast", query, ar)

	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	return ar, err
}

func (c *Service) Conditions(query interface{}) (*ConditionsResponse, error) {
	cr := &ConditionsResponse{}
	err := c.request("conditions", query, cr)

	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	return cr, err
}

func (c *Service) request(path string, query, v interface{}) error {
	qs := fmt.Sprintf("/%s/q/%d.json", path, query)
	resp, err := c.client.Get(API_URL + c.ApiKey + qs)
	if err != nil {
		log.Fatal("whoops in request:", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return err
}
