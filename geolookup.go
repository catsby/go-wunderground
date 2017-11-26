package wunderground

// Use to request the geolookup feature in Service.Request
var FGeoLookup = "geolookup"

type Location struct {
	Type                  string `json:"type"`
	Country               string `json:"country"`
	CountryIso3166        string `json:"country_iso3166"`
	CountryName           string `json:"country_name"`
	State                 string `json:"state"`
	City                  string `json:"city"`
	TzShort               string `json:"tz_short"`
	TzLong                string `json:"tz_long"`
	Lat                   string `json:"lat"`
	Lon                   string `json:"lon"`
	Zip                   string `json:"zip"`
	Magic                 string `json:"magic"`
	Wmo                   string `json:"wmo"`
	L                     string `json:"l"`
	Requesturl            string `json:"requesturl"`
	Wuiurl                string `json:"wuiurl"`
	NearbyWeatherStations struct {
		Airport struct {
			Station []struct {
				City    string `json:"city"`
				State   string `json:"state"`
				Country string `json:"country"`
				Icao    string `json:"icao"`
				Lat     string `json:"lat"`
				Lon     string `json:"lon"`
			} `json:"station"`
		} `json:"airport"`
		Pws struct {
			Station []struct {
				Neighborhood string `json:"neighborhood"`
				City         string `json:"city"`
				State        string `json:"state"`
				Country      string `json:"country"`
				ID           string `json:"id"`
				DistanceKm   int    `json:"distance_km"`
				DistanceMi   int    `json:"distance_mi"`
			} `json:"station"`
		} `json:"pws"`
	} `json:"nearby_weather_stations"`
}
