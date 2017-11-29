package wunderground

// Use to request the hourly feature in Service.Request
var FHourly = "hourly"

// Use to request the hourly10day feature in Service.Request
var FHourly10Day = "hourly10day"

type HourlyForecast []struct {
	FCTTIME struct {
		Hour                   string `json:"hour"`
		HourPadded             string `json:"hour_padded"`
		Min                    string `json:"min"`
		Sec                    string `json:"sec"`
		Year                   string `json:"year"`
		Mon                    string `json:"mon"`
		MonPadded              string `json:"mon_padded"`
		MonAbbrev              string `json:"mon_abbrev"`
		Mday                   string `json:"mday"`
		MdayPadded             string `json:"mday_padded"`
		Yday                   string `json:"yday"`
		Isdst                  string `json:"isdst"`
		Epoch                  string `json:"epoch"`
		Pretty                 string `json:"pretty"`
		Civil                  string `json:"civil"`
		MonthName              string `json:"month_name"`
		MonthNameAbbrev        string `json:"month_name_abbrev"`
		WeekdayName            string `json:"weekday_name"`
		WeekdayNameNight       string `json:"weekday_name_night"`
		WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
		WeekdayNameUnlang      string `json:"weekday_name_unlang"`
		WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
		Ampm                   string `json:"ampm"`
		Tz                     string `json:"tz"`
		Age                    string `json:"age"`
	} `json:"FCTTIME"`
	Temp struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"temp"`
	Dewpoint struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"dewpoint"`
	Condition string `json:"condition"`
	Icon      string `json:"icon"`
	IconURL   string `json:"icon_url"`
	Fctcode   string `json:"fctcode"`
	Sky       string `json:"sky"`
	Wspd      struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"wspd"`
	Wdir struct {
		Dir     string `json:"dir"`
		Degrees string `json:"degrees"`
	} `json:"wdir"`
	Wx        string `json:"wx"`
	Uvi       string `json:"uvi"`
	Humidity  string `json:"humidity"`
	Windchill struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"windchill"`
	Heatindex struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"heatindex"`
	Feelslike struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"feelslike"`
	Qpf struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"qpf"`
	Snow struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"snow"`
	Pop  string `json:"pop"`
	Mslp struct {
		English string `json:"english"`
		Metric  string `json:"metric"`
	} `json:"mslp"`
}
