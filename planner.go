package wunderground

import (
	"fmt"
	"time"
)

// Use to request the planner feature in Service.Request
func FPlanner(start, end time.Time) string {
	return fmt.Sprintf("planner_%s%s", start.Format("0102"), end.Format("0102"))
}

type Trip struct {
	Title          string `json:"title"`
	AirportCode    string `json:"airport_code"`
	Error          string `json:"error"`
	PeriodOfRecord struct {
		DateStart struct {
			Date struct {
				Epoch        string `json:"epoch"`
				Pretty       string `json:"pretty"`
				Day          int    `json:"day"`
				Month        int    `json:"month"`
				Year         int    `json:"year"`
				Yday         int    `json:"yday"`
				Hour         int    `json:"hour"`
				Min          string `json:"min"`
				Sec          int    `json:"sec"`
				Isdst        string `json:"isdst"`
				Monthname    string `json:"monthname"`
				WeekdayShort string `json:"weekday_short"`
				Weekday      string `json:"weekday"`
				Ampm         string `json:"ampm"`
				TzShort      string `json:"tz_short"`
				TzLong       string `json:"tz_long"`
			} `json:"date"`
		} `json:"date_start"`
		DateEnd struct {
			Date struct {
				Epoch        string `json:"epoch"`
				Pretty       string `json:"pretty"`
				Day          int    `json:"day"`
				Month        int    `json:"month"`
				Year         int    `json:"year"`
				Yday         int    `json:"yday"`
				Hour         int    `json:"hour"`
				Min          string `json:"min"`
				Sec          int    `json:"sec"`
				Isdst        string `json:"isdst"`
				Monthname    string `json:"monthname"`
				WeekdayShort string `json:"weekday_short"`
				Weekday      string `json:"weekday"`
				Ampm         string `json:"ampm"`
				TzShort      string `json:"tz_short"`
				TzLong       string `json:"tz_long"`
			} `json:"date"`
		} `json:"date_end"`
	} `json:"period_of_record"`
	TempHigh struct {
		Min struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"min"`
		Avg struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"avg"`
		Max struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"max"`
	} `json:"temp_high"`
	TempLow struct {
		Min struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"min"`
		Avg struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"avg"`
		Max struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"max"`
	} `json:"temp_low"`
	Precip struct {
		Min struct {
			In string `json:"in"`
			Cm string `json:"cm"`
		} `json:"min"`
		Avg struct {
			In string `json:"in"`
			Cm string `json:"cm"`
		} `json:"avg"`
		Max struct {
			In string `json:"in"`
			Cm string `json:"cm"`
		} `json:"max"`
	} `json:"precip"`
	DewpointHigh struct {
		Min struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"min"`
		Avg struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"avg"`
		Max struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"max"`
	} `json:"dewpoint_high"`
	DewpointLow struct {
		Min struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"min"`
		Avg struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"avg"`
		Max struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"max"`
	} `json:"dewpoint_low"`
	CloudCover struct {
		Cond string `json:"cond"`
	} `json:"cloud_cover"`
	ChanceOf struct {
		Tempoversixty struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"tempoversixty"`
		Chanceofwindyday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofwindyday"`
		Chanceofpartlycloudyday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofpartlycloudyday"`
		Chanceofsunnycloudyday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofsunnycloudyday"`
		Chanceofcloudyday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofcloudyday"`
		Chanceoffogday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceoffogday"`
		Chanceofhumidday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofhumidday"`
		Chanceofprecip struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofprecip"`
		Chanceofrainday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofrainday"`
		Tempoverninety struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"tempoverninety"`
		Chanceofthunderday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofthunderday"`
		Chanceofsnowonground struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofsnowonground"`
		Chanceoftornadoday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceoftornadoday"`
		Chanceofsultryday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofsultryday"`
		Tempbelowfreezing struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"tempbelowfreezing"`
		Tempoverfreezing struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"tempoverfreezing"`
		Chanceofhailday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofhailday"`
		Chanceofsnowday struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Percentage  string `json:"percentage"`
		} `json:"chanceofsnowday"`
	} `json:"chance_of"`
}
