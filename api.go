package wunderground

type ApiResponse struct {
	Response            `json:"response"`
	*Alerts             `json:"alerts"`
	*Almanac            `json:"almanac"`
	*CurrentObservation `json:"current_observation"`
	*Forecast           `json:"forecast"`
	*History            `json:"history"`
	*HourlyForecast     `json:"hourly_forecast"`
	*Location           `json:"location"`
	*MoonPhase          `json:"moon_phase"`
	*RawTide            `json:"rawtide"`
	*Satellite          `json:"satellite"`
	*Tide               `json:"tide"`
	*Trip               `json:"trip"`
	*WebCams            `json:"webcams"`
}

type Response struct {
	Version        string `json:"version"`
	TermsofService string `json:"termsofService"`
	Features       struct {
		Alerts        int `json:"alerts"`
		Almanac       int `json:"almanac"`
		Astronomy     int `json:"astronomy"`
		Conditions    int `json:"conditions"`
		Forecast      int `json:"forecast"`
		Forecast10Day int `json:"forecast10day"`
		GeoLookup     int `json:"geolookup"`
		History       int `json:"history"`
		Hourly        int `json:"hourly"`
		Hourly10Day   int `json:"hourly10day"`
		Planner       int `json:"planner"`
		RawTide       int `json:"rawtide"`
		Satellite     int `json:"satellite"`
		Tide          int `json:"tide"`
		WebCams       int `json:"webcams"`
		Yesterday     int `json:"yesterday"`
	} `json:"features"`
	*Error `json:"error"`
}

type Error struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
