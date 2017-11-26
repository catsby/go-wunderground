package wunderground

// Use to request the rawtide feature in Service.Request
var FRawTide = "rawtide"

// Use to request the tide feature in Service.Request
var FTide = "tide"

type RawTide struct {
	TideInfo   []TideInfo `json:"tideInfo"`
	RawTideObs []struct {
		Epoch  int     `json:"epoch"`
		Height float64 `json:"height"`
	} `json:"rawTideObs"`
	RawTideStats []struct {
		Maxheight float64 `json:"maxheight"`
		Minheight float64 `json:"minheight"`
	} `json:"rawTideStats"`
}

type Tide struct {
	TideInfo    []TideInfo `json:"tideInfo"`
	TideSummary []struct {
		Date struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
			Epoch  string `json:"epoch"`
		} `json:"date"`
		Utcdate struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
			Epoch  string `json:"epoch"`
		} `json:"utcdate"`
		Data struct {
			Height string `json:"height"`
			Type   string `json:"type"`
		} `json:"data"`
	} `json:"tideSummary"`
	TideSummaryStats []struct {
		Maxheight float64 `json:"maxheight"`
		Minheight float64 `json:"minheight"`
	} `json:"tideSummaryStats"`
}

type TideInfo struct {
	TideSite string `json:"tideSite"`
	Lat      string `json:"lat"`
	Lon      string `json:"lon"`
	Units    string `json:"units"`
	Type     string `json:"type"`
	Tzname   string `json:"tzname"`
}
