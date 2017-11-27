package wunderground

import (
	"fmt"
	"strconv"
)

// Use to request the astronomy feature in Service.Request
var FAstronomy = "astronomy"

type MoonPhase struct {
	PercentIlluminated string        `json:"percentIlluminated"`
	AgeOfMoon          string        `json:"ageOfMoon"`
	CurrentTime        MoonPhaseTime `json:"current_time"`
	Sunrise            MoonPhaseTime `json:"sunrise"`
	Sunset             MoonPhaseTime `json:"sunset"`
}

type MoonPhaseTime struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

func (m *MoonPhase) ToString() string {
	// Slightly more complicated, but close enough. Based on
	// https://www.spaceweatherlive.com/en/moon-phases-calendar
	var phase string
	var age, _ = strconv.Atoi(m.AgeOfMoon)
	switch {
	case age < 2:
		phase = "New moon"
	case age < 8:
		phase = "Waxing crescent"
	case age < 10:
		phase = "First quarter"
	case age < 15:
		phase = "Waxing gibbous"
	case age < 17:
		phase = "Full moon"
	case age < 23:
		phase = "Waning gibbous"
	case age < 25:
		phase = "Last quarter"
	case age < 29:
		phase = "Waning crescent"
	}

	return fmt.Sprintf("Moon Phase: %s (%s%% illuminated)\nSunrise   : %s:%s\nSunset    : %s:%s",
		phase, m.PercentIlluminated, m.Sunrise.Hour, m.Sunrise.Minute, m.Sunset.Hour, m.Sunset.Minute)
}
