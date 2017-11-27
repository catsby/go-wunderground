//go:generate go-bindata templates/...
package main

import (
	"flag"
	"fmt"
	"log"

	wu "github.com/catsby/go-wunderground"

	"os"
	"text/template"
)

var (
	search = flag.String("search", "", "Search query. Can be postal (65203) or city ('Columbia, MO')")
	apiKey = flag.String("api_key", os.Getenv("WUNDERGROUND_API_KEY"), "API Key. May also be set via the environment: WUNDERGROUND_API_KEY")
)

func main() {
	flag.Parse()

	var query wu.Query
	if len(*search) == 0 && flag.NArg() > 0 {
		*search = flag.Arg(0)
	}
	if len(*search) == 0 {
		log.Fatal("No search term provided")
	}

	if len(*apiKey) == 0 {
		log.Fatal("No API key found")
	}

	// Load local resources before making API calls.
	t := template.New("templates/list")
	if _, err := t.Parse(string(MustAsset("templates/list"))); err != nil {
		log.Fatalf("failed to load output template: %s", err)
	}

	client := wu.NewDevLimitedService(*apiKey)

	fmt.Printf("Getting weather for %v...\n\n", *search)
	query.User = *search

	response, err := client.Request([]string{wu.FForecast, wu.FConditions}, &query)
	if err != nil {
		log.Fatalf("response failed: %s", err)
	}

	fmt.Println("Forcast for ", response.CurrentObservation.DisplayLocation.Full)
	if err := t.Execute(os.Stdout, response.Forecast.TxtForecast.ForecastDay); err != nil {
		log.Fatalf("failed in output template: %s", err)
	}
}
