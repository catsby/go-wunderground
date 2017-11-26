package wunderground

type Response struct {
	Version        string `json:"version"`
	TermsofService string `json:"termsofService"`
	Features       struct {
		Forecast int `json:"forecast"`
	} `json:"features"`
}
