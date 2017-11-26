package wunderground

// Use to request the astronomy feature in Service.Request
var FAstronomy = "astronomy"

type MoonPhase struct {
	PercentIlluminated string `json:"percentIlluminated"`
	AgeOfMoon          string `json:"ageOfMoon"`
	CurrentTime        struct {
		Hour   string `json:"hour"`
		Minute string `json:"minute"`
	} `json:"current_time"`
	Sunrise struct {
		Hour   string `json:"hour"`
		Minute string `json:"minute"`
	} `json:"sunrise"`
	Sunset struct {
		Hour   string `json:"hour"`
		Minute string `json:"minute"`
	} `json:"sunset"`
}
