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

	if len(os.Args) == 1 && search == "" {
		log.Fatal("No search term provided")
	} else if search == "" {
		search = os.Args[1]
	}

	key := os.Getenv("WUNDERGROUND_API_KEY")
	if len(key) == 0 {
		log.Fatal("No API key found")
	}

	client := wunderground.NewService(key)

	fmt.Printf("Getting weather for %v...\n\n", search)

	forecast, err := client.Forecast(search)

	if err != nil {
		panic(err)
	}

	conditions, err := client.Conditions(search)

	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("templates/list")
	fmt.Println("Forcast for", conditions.LocationName())
	err = t.Execute(os.Stdout, forecast.Forecast.TxtForecast.Days)

	if err != nil {
		panic(err)
	}

}
