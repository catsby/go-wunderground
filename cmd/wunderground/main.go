//go:generate go-bindata templates/...
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/catsby/go-wunderground"

	"os"
	"text/template"
)

func main() {
	var search string
	flag.StringVar(&search, "search", "", "Search query. Can be postal (65203) or city ('Columbia, MO')")
	flag.Parse()

	var query wunderground.Query
	if len(os.Args) == 1 && search == "" {
		log.Fatal("No search term provided")
	} else if search == "" {
		search = os.Args[1]
	}
	query.User = search

	key := os.Getenv("WUNDERGROUND_API_KEY")
	if len(key) == 0 {
		log.Fatal("No API key found")
	}

	client := wunderground.NewDevLimitedService(key)

	fmt.Printf("Getting weather for %v...\n\n", search)

	forecast, err := client.Forecast(&query)

	if err != nil {
		panic(err)
	}

	conditions, err := client.Conditions(&query)

	if err != nil {
		panic(err)
	}

	t := template.New("templates/list")
	t, err = t.Parse(string(MustAsset("templates/list")))
	if err != nil {
		panic(err)
	}

	fmt.Println("Forcast for", conditions.LocationName())
	err = t.Execute(os.Stdout, forecast.Forecast.TxtForecast.Forecastday)
	if err != nil {
		panic(err)
	}

}
