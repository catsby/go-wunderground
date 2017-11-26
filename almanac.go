package wunderground

// Use to request the almanac feature in Service.Request
var FAlmanac = "almanac"

type Almanac struct {
	AirportCode string `json:"airport_code"`
	TempHigh    struct {
		Normal struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"normal"`
		Record struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"record"`
		Recordyear string `json:"recordyear"`
	} `json:"temp_high"`
	TempLow struct {
		Normal struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"normal"`
		Record struct {
			F string `json:"F"`
			C string `json:"C"`
		} `json:"record"`
		Recordyear string `json:"recordyear"`
	} `json:"temp_low"`
}
