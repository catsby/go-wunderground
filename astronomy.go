package wunderground

import (
	"fmt"
)

// Use to request the astronomy feature in Service.Request
var FAstronomy = "astronomy"

type MoonPhase struct {
	PercentIlluminated string        `json:"percentIlluminated"`
	AgeOfMoon          string        `json:"ageOfMoon"`
	PhaseOfMoon        string        `json:"phaseofMoon"`
	Hemisphere         string        `json:"hemisphere"`
	CurrentTime        MoonPhaseTime `json:"current_time"`
	Sunrise            MoonPhaseTime `json:"sunrise"`
	Sunset             MoonPhaseTime `json:"sunset"`
}

type MoonPhaseTime struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

func (m *MoonPhase) ToString() string {
	return fmt.Sprintf("Moon Phase: %s (%s%% illuminated)\nSunrise   : %s:%s\nSunset    : %s:%s",
		m.PhaseOfMoon, m.PercentIlluminated, m.Sunrise.Hour, m.Sunrise.Minute, m.Sunset.Hour, m.Sunset.Minute)
}
