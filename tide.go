package wunderground

import (
	"fmt"
	"strings"
	"time"
)

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
		Date    PrettyDate `json:"date"`
		UTCDate PrettyDate `json:"utcdate"`
		Data    struct {
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

func (t *Tide) ToString() string {
	if len(t.TideSummary) == 0 {
		return "No tidal data available."
	}

	var res []string
	res = append(res, "Tidal data for "+t.TideInfo[0].TideSite)

	var prevDate time.Time
	for _, s := range t.TideSummary {
		date, _ := s.Date.ToDate()
		if date != prevDate {
			res = append(res, date.Format("Jan 2, 2006: "))
		}
		prevDate = date

		res = append(res, fmt.Sprintf("     %s at %s", s.Data.Type, date.Format("3:04 AM")))
	}

	return strings.Join(res, "\n")
}
