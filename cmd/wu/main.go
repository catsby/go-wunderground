// A clone of wu (github.com/sramsay/wu) which relies on the full go API.
package main

import (
	"flag"
	"fmt"
	wu "github.com/catsby/go-wunderground"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
)

var (
	apiKey     = flag.String("api_key", os.Getenv("WUNDERGROUND_API_KEY"), "API Key. May also be set via the environment: WUNDERGROUND_API_KEY")
	conditions = flag.Bool("conditions", false, "Reports the current weather conditions")
	alerts     = flag.Bool("alerts", false, "Reports any active weather alerts")
	lookup     = flag.Bool("lookup", false, "Lookup the codes for the weather stations in a particular area")
	astro      = flag.Bool("astro", false, "Reports sunrise, sunset, and lunar phase")
	forecast   = flag.Bool("forecast", false, "Reports the current (3-day) forecast")
	forecast10 = flag.Bool("forecast10", false, "Reports the current (7-day) forecast")
	almanac    = flag.Bool("almanac", false, "Reports average high, low and record temperatures")
	yesterday  = flag.Bool("yesterday", false, "Reports yesterday's weather data")
	history    = flag.String("history", "", "Reports historical data for a particular day --history=\"YYYYMMDD\"")
	planner    = flag.String("planner", "", "Reports historical data for a particular date range (30-day max) --planner=\"MMDDMMDD\"")
	tides      = flag.Bool("tides", false, "Reports tidal data (if available")
	version    = flag.Bool("version", false, "Print the version number")
	all        = flag.Bool("all", false, "Show all weather data")
	query_s    = flag.String("s", "",
		"Weather station: \"city, state|country\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG. Defaults to the user's condrc or auto lookup.")
)

// Tries very hard to figure out the user's target destination.
func normalizeLocation(cfg *Config) *wu.Query {
	var q wu.Query

	var loc string
	switch {
	case len(*query_s) > 0:
		loc = *query_s
	case flag.NArg() > 0:
		loc = strings.Join(flag.Args(), " ")
	case cfg != nil && len(cfg.Station) > 0:
		loc = cfg.Station
	default:
		q.Auto = true
		return &q
	}
	loc = strings.TrimSpace(loc)

	city := strings.SplitN(loc, ",", 2)
	if len(city) > 1 {
		// In practice, state and country are treated the same.
		q.City = strings.Replace(strings.TrimSpace(city[0]), " ", "_", -1)
		q.USState = strings.Replace(strings.TrimSpace(city[1]), " ", "_", -1)
	} else {
		q.User = loc
	}

	return &q
}

func getKey(cfg *Config) string {
	var res string
	switch {
	case len(*apiKey) > 0:
		res = *apiKey
	case cfg != nil && len(cfg.Key) > 0:
		res = cfg.Key
	}

	return res
}

func getFeatures() [][]string {
	// Some features aren't compatible. If the user requests such a combination, simply make two passes.
	var features []string
	var features2 []string
	if *alerts || *all {
		features = append(features, wu.FAlerts)
	}
	if *almanac || *all {
		features = append(features, wu.FAlmanac)
	}
	if *astro || *all {
		features = append(features, wu.FAstronomy)
	}
	if *conditions || *all {
		features = append(features, wu.FConditions)
	}
	if *forecast || *all {
		features = append(features, wu.FForecast)
	}
	if *forecast10 || *all {
		features = append(features, wu.FForecast10Day)
	}
	if len(*history) > 0 {
		features = append(features, wu.FHistoryUser(*history))
	}
	if *yesterday || *all {
		features = append(features, wu.FYesterday)
	}
	if len(*planner) > 0 {
		if len(*history) > 0 {
			features2 = append(features2, wu.FPlannerUser(*planner))
		} else {
			features = append(features, wu.FPlannerUser(*planner))
		}
	}
	if *tides || *all {
		features = append(features, wu.FTide)
	}
	if *lookup || *all {
		features = append(features, wu.FGeoLookup)
	}

	if len(features) == 0 {
		features = append(features, wu.FConditions)
	}

	res := [][]string{features}
	if len(features2) > 0 {
		res = append(res, features2)
	}

	return res
}

func main() {
	flag.Parse()

	var cfg *Config
	if user, err := user.Current(); err == nil {
		if cfg, err = ParseConfig(path.Join(user.HomeDir, ".condrc")); err != nil {
			log.Fatal(err)
		}
	}

	query := normalizeLocation(cfg)

	key := getKey(cfg)
	if len(key) == 0 {
		log.Fatal("API Key not found")
	}

	featureListList := getFeatures()

	client := wu.NewDevLimitedService(key)

	for _, f := range featureListList {
		ar, err := client.Request(f, query)
		if err != nil {
			log.Fatalf("failed during API query: %s", err)
		}

		fmt.Println(ar.ToString())
	}
}
