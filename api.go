package wunderground

type ApiResponse struct {
	Response            `json:"response"`
	*CurrentObservation `json:"current_observation"`
	*Forecast           `json:"forecast"`
	*History            `json:"history"`
}

type Response struct {
	Version        string `json:"version"`
	TermsofService string `json:"termsofService"`
	Features       struct {
		Conditions int `json:"conditions"`
		Forecast   int `json:"forecast"`
		History    int `json:"history"`
	} `json:"features"`
}
