package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

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

func (f *Forecast) String() string {
	return fmt.Sprintf("%v\n", f.TxtForecast)
}

type Response struct {
	Version string `json:"version,omitempty"`
	TOS     string `json:"termsOfService,omitempty"`
}

type ApiResponse struct {
	Response Response `json:"response"`
	Forecast Forecast `json:"forecast"`
}

const API_URL = "https://api.wunderground.com/api/"

func main() {
	fmt.Println("Starting!")

	var pc int
	flag.IntVar(&pc, "postal", 65203, "Postal code to search. Default 65203")
	flag.Parse()

	key := os.Getenv("WUNDERGROUND_API_KEY")
	if len(key) == 0 {
		log.Fatal("No API key found")
	}

	fmt.Printf("\n\tGetting weather for %d\n\n", pc)
	query := fmt.Sprintf("/forecast/q/%d.json", pc)
	resp, err := http.Get(API_URL + key + query)
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
	t, err := template.ParseFiles("templates/list")
	fmt.Println("t: %v", t)
	fmt.Println("Forcast Days:\n")
	err = t.Execute(os.Stdout, api_response.Forecast.TxtForecast.Days)

	if err != nil {
		panic(err)
	}
}
