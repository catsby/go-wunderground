package wunderground

import (
	"fmt"
)

// Use to request the almanac feature in Service.Request
var FAlmanac = "almanac"

type Almanac struct {
	AirportCode string      `json:"airport_code"`
	TempHigh    AlmanacTemp `json:"temp_high"`
	TempLow     AlmanacTemp `json:"temp_low"`
}

type AlmanacTemp struct {
	Normal     Temperature `json:"normal"`
	Record     Temperature `json:"record"`
	RecordYear string      `json:"recordyear"`
}

func (a *Almanac) ToString() string {
	return fmt.Sprintf("Normal high: %s\u00B0 F (%s\u00B0 C)\n", a.TempHigh.Normal.F, a.TempHigh.Normal.C) +
		fmt.Sprintf("Record high: %s\u00B0 F (%s\u00B0 C) [%s]\n", a.TempHigh.Record.F, a.TempHigh.Record.C, a.TempHigh.RecordYear) +
		fmt.Sprintf("Normal low : %s\u00B0 F (%s\u00B0 C)\n", a.TempLow.Normal.F, a.TempLow.Normal.C) +
		fmt.Sprintf("Record low : %s\u00B0 F (%s\u00B0 C) [%s]", a.TempLow.Record.F, a.TempLow.Record.C, a.TempLow.RecordYear)
}
