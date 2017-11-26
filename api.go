package wunderground

type ApiResponse struct {
	Response            `json:"response"`
	*CurrentObservation `json:"current_observation"`
	*Forecast           `json:"forecast"`
}

type Response struct {
	Version        string `json:"version"`
	TermsofService string `json:"termsofService"`
	Features       struct {
		Conditions int `json:"conditions"`
		Forecast   int `json:"forecast"`
	} `json:"features"`
}
