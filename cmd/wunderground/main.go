package main

import (
	"flag"
	"fmt"

	"github.com/catsby/go-wunderground"

	"os"
	"text/template"
)

func main() {
	var pc int
	flag.IntVar(&pc, "postal", 65203, "Postal code to search. Default 65203")
	flag.Parse()

	fmt.Printf("\n\tGetting weather for %d\n\n", pc)

	forecast, err := wunderground.Get(pc)

	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("templates/list")
	fmt.Println("Forcast Days:\n")
	err = t.Execute(os.Stdout, forecast.Forecast.TxtForecast.Days)

	if err != nil {
		panic(err)
	}

}
