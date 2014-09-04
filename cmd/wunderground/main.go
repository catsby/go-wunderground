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
	var pc int
	flag.IntVar(&pc, "postal", 65203, "Postal code to search. Default 65203")
	flag.Parse()

	key := os.Getenv("WUNDERGROUND_API_KEY")
	if len(key) == 0 {
		log.Fatal("No API key found")
	}

	client := wunderground.NewService(key)

	fmt.Printf("Getting weather for %d...\n\n", pc)

	forecast, err := client.Forecast(pc)

	if err != nil {
		panic(err)
	}

	conditions, err := client.Conditions(pc)

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
