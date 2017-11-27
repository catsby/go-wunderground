package wunderground

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Use to request the planner feature in Service.Request
func FPlanner(start, end time.Time) string {
	return fmt.Sprintf("planner_%s%s", start.Format("0102"), end.Format("0102"))
}

// Use to request the planner feature in Service.Request
func FPlannerUser(userRange string) string {
	return "planner_" + userRange
}

type Trip struct {
	Title          string `json:"title"`
	AirportCode    string `json:"airport_code"`
	Error          string `json:"error"`
	PeriodOfRecord struct {
		DateStart TripDate `json:"date_start"`
		DateEnd   TripDate `json:"date_end"`
	} `json:"period_of_record"`
	TempHigh     TripTempRange `json:"temp_high"`
	TempLow      TripTempRange `json:"temp_low"`
	Precip       TripTempRange `json:"precip"`
	DewpointHigh TripTempRange `json:"dewpoint_high"`
	DewpointLow  TripTempRange `json:"dewpoint_low"`
	CloudCover   struct {
		Cond string `json:"cond"`
	} `json:"cloud_cover"`
	ChanceOf struct {
		TempOverSixty           TripChance `json:"tempoversixty"`
		ChanceOfWindyDay        TripChance `json:"chanceofwindyday"`
		ChanceOfPartlyCloudyDay TripChance `json:"chanceofpartlycloudyday"`
		ChanceOfSunnyCloudyDay  TripChance `json:"chanceofsunnycloudyday"`
		ChanceOfCloudyDay       TripChance `json:"chanceofcloudyday"`
		ChanceOfFogDay          TripChance `json:"chanceoffogday"`
		ChanceOfHumidDay        TripChance `json:"chanceofhumidday"`
		ChanceOfPrecip          TripChance `json:"chanceofprecip"`
		ChanceOfRainDay         TripChance `json:"chanceofrainday"`
		TempOverNinety          TripChance `json:"tempoverninety"`
		ChanceOfThunderDay      TripChance `json:"chanceofthunderday"`
		ChanceOfSnowOnGround    TripChance `json:"chanceofsnowonground"`
		ChanceOfTornadoDay      TripChance `json:"chanceoftornadoday"`
		ChanceOfSultryDay       TripChance `json:"chanceofsultryday"`
		TempBelowFreezing       TripChance `json:"tempbelowfreezing"`
		TempOverFreezing        TripChance `json:"tempoverfreezing"`
		ChanceOfHailDay         TripChance `json:"chanceofhailday"`
		ChanceOfSnowDay         TripChance `json:"chanceofsnowday"`
	} `json:"chance_of"`
}

type TripDate struct {
	TripDateInner `json:"date"`
}

type TripDateInner struct {
	Epoch        string `json:"epoch"`
	Pretty       string `json:"pretty"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	YDay         int    `json:"yday"`
	Hour         int    `json:"hour"`
	Min          string `json:"min"`
	Sec          int    `json:"sec"`
	IsDST        string `json:"isdst"`
	MonthName    string `json:"monthname"`
	WeekdayShort string `json:"weekday_short"`
	Weekday      string `json:"weekday"`
	AmPm         string `json:"ampm"`
	TzShort      string `json:"tz_short"`
	TzLong       string `json:"tz_long"`
}

type TripTempRange struct {
	Min Temperature `json:"min"`
	Avg Temperature `json:"avg"`
	Max Temperature `json:"max"`
}

type TripChance struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Percentage  string `json:"percentage"`
}

func (d *TripDateInner) ToDate() (time.Time, error) {
	mi, _ := strconv.Atoi(d.Min)
	loc, err := time.LoadLocation(d.TzLong)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(d.Year, time.Month(d.Month), d.Day, d.Hour, mi, d.Sec, 0 /*nsec*/, loc), nil
}

func (t *Trip) ToString() string {
	if len(t.Error) != 0 {
		return t.Error
	}

	var res []string
	ch := t.ChanceOf
	res = append(res, t.Title)
	res = append(res, "Station: "+t.AirportCode)
	res = append(res, "Chance of: ")

	res = append(res, "   Temps:")
	res = append(res, "      Over 90 F (32 C): "+ch.TempOverNinety.Percentage+"%")
	res = append(res, "      Between 60 F (15 C) and 90 F (32 C): "+ch.TempOverSixty.Percentage+"%")
	res = append(res, "      Between 32 F (0 C) and 60 F (16 C): "+ch.TempOverFreezing.Percentage+"%")
	res = append(res, "      Below 32 F (0 C): "+ch.TempBelowFreezing.Percentage+"%")

	res = append(res, "   Dewpoint above 70 F (21 C): "+ch.ChanceOfSultryDay.Percentage+"%")
	res = append(res, "   Dewpoint above 60 F (15 C): "+ch.ChanceOfHumidDay.Percentage+"%")
	res = append(res, "   Winds over 10 mph (15 km/h): "+ch.ChanceOfWindyDay.Percentage+"%")

	res = append(res, fmt.Sprintf("   %s day: %s%%", ch.ChanceOfSunnyCloudyDay.Name, ch.ChanceOfSunnyCloudyDay.Percentage))
	res = append(res, fmt.Sprintf("   %s day: %s%%", ch.ChanceOfCloudyDay.Name, ch.ChanceOfCloudyDay.Percentage))
	res = append(res, fmt.Sprintf("   %s day: %s%%", ch.ChanceOfPartlyCloudyDay.Name, ch.ChanceOfPartlyCloudyDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfPrecip.Name, ch.ChanceOfPrecip.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfFogDay.Name, ch.ChanceOfFogDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfRainDay.Name, ch.ChanceOfRainDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfThunderDay.Name, ch.ChanceOfThunderDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfTornadoDay.Name, ch.ChanceOfTornadoDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfHailDay.Name, ch.ChanceOfHailDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfSnowDay.Name, ch.ChanceOfSnowDay.Percentage))
	res = append(res, fmt.Sprintf("   %s: %s%%", ch.ChanceOfSnowOnGround.Name, ch.ChanceOfSnowOnGround.Percentage))

	return strings.Join(res, "\n")
}
