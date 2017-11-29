package wunderground

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Use to request the history feature in Service.Request
// Note: WU currently supports only a single history feature per query.
func FHistory(date *time.Time) string {
	return fmt.Sprintf("history_%s", date.Format("20070102"))
}

// Use to request the history feature in Service.Request
// Note: WU currently supports only a single history feature per query.
func FHistoryUser(userDate string) string {
	return "history_" + userDate
}

// Use to request the yesterday feature in Service.Request
var FYesterday = "yesterday"

type History struct {
	Date         PrettyDate `json:"date"`
	Utcdate      PrettyDate `json:"utcdate"`
	Observations []struct {
		Date       PrettyDate `json:"date"`
		Utcdate    PrettyDate `json:"utcdate"`
		Tempm      string     `json:"tempm"`
		Tempi      string     `json:"tempi"`
		DewPtm     string     `json:"dewptm"`
		DewPti     string     `json:"dewpti"`
		Hum        string     `json:"hum"`
		WSpdm      string     `json:"wspdm"`
		WSpdi      string     `json:"wspdi"`
		Wgustm     string     `json:"wgustm"`
		Wgusti     string     `json:"wgusti"`
		WDird      string     `json:"wdird"`
		WDire      string     `json:"wdire"`
		Vism       string     `json:"vism"`
		Visi       string     `json:"visi"`
		Pressurem  string     `json:"pressurem"`
		Pressurei  string     `json:"pressurei"`
		Windchillm string     `json:"windchillm"`
		Windchilli string     `json:"windchilli"`
		Heatindexm string     `json:"heatindexm"`
		Heatindexi string     `json:"heatindexi"`
		Precipm    string     `json:"precipm"`
		Precipi    string     `json:"precipi"`
		Conds      string     `json:"conds"`
		Icon       string     `json:"icon"`
		Fog        string     `json:"fog"`
		Rain       string     `json:"rain"`
		Snow       string     `json:"snow"`
		Hail       string     `json:"hail"`
		Thunder    string     `json:"thunder"`
		Tornado    string     `json:"tornado"`
		Metar      string     `json:"metar"`
	} `json:"observations"`
	DailySummary []struct {
		Date                               PrettyDate `json:"date"`
		Fog                                string     `json:"fog"`
		Rain                               string     `json:"rain"`
		Snow                               string     `json:"snow"`
		Snowfallm                          string     `json:"snowfallm"`
		Snowfalli                          string     `json:"snowfalli"`
		MonthToDateSnowfallm               string     `json:"monthtodatesnowfallm"`
		MonthToDateSnowfalli               string     `json:"monthtodatesnowfalli"`
		Since1JulSnowfallm                 string     `json:"since1julsnowfallm"`
		Since1JulSnowfalli                 string     `json:"since1julsnowfalli"`
		SnowDepthm                         string     `json:"snowdepthm"`
		SnowDepthi                         string     `json:"snowdepthi"`
		Hail                               string     `json:"hail"`
		Thunder                            string     `json:"thunder"`
		Tornado                            string     `json:"tornado"`
		MeanTempm                          string     `json:"meantempm"`
		MeanTempi                          string     `json:"meantempi"`
		MeanDewPtm                         string     `json:"meandewptm"`
		MeanDewPti                         string     `json:"meandewpti"`
		MeanPressurem                      string     `json:"meanpressurem"`
		MeanPressurei                      string     `json:"meanpressurei"`
		MeanWindSpdm                       string     `json:"meanwindspdm"`
		MeanWindSpdi                       string     `json:"meanwindspdi"`
		MeanWDire                          string     `json:"meanwdire"`
		MeanWDird                          string     `json:"meanwdird"`
		MeanVism                           string     `json:"meanvism"`
		MeanVisi                           string     `json:"meanvisi"`
		Humidity                           string     `json:"humidity"`
		MaxTempm                           string     `json:"maxtempm"`
		MaxTempi                           string     `json:"maxtempi"`
		MinTempm                           string     `json:"mintempm"`
		MinTempi                           string     `json:"mintempi"`
		MaxHumidity                        string     `json:"maxhumidity"`
		MinHumidity                        string     `json:"minhumidity"`
		MaxDewPtm                          string     `json:"maxdewptm"`
		MaxDewPti                          string     `json:"maxdewpti"`
		MinDewPtm                          string     `json:"mindewptm"`
		MinDewPti                          string     `json:"mindewpti"`
		MaxPressurem                       string     `json:"maxpressurem"`
		MaxPressurei                       string     `json:"maxpressurei"`
		MinPressurem                       string     `json:"minpressurem"`
		MinPressurei                       string     `json:"minpressurei"`
		MaxWSpdm                           string     `json:"maxwspdm"`
		MaxWSpdi                           string     `json:"maxwspdi"`
		MinWSpdm                           string     `json:"minwspdm"`
		MinWSpdi                           string     `json:"minwspdi"`
		MaxVism                            string     `json:"maxvism"`
		MaxVisi                            string     `json:"maxvisi"`
		MinVism                            string     `json:"minvism"`
		MinVisi                            string     `json:"minvisi"`
		GDegreeDays                        string     `json:"gdegreedays"`
		HeatingDegreeDays                  string     `json:"heatingdegreedays"`
		CoolingDegreeDays                  string     `json:"coolingdegreedays"`
		Precipm                            string     `json:"precipm"`
		Precipi                            string     `json:"precipi"`
		PrecipSource                       string     `json:"precipsource"`
		HeatingDegreeDaysNormal            string     `json:"heatingdegreedaysnormal"`
		MonthToDateHeatingDegreeDays       string     `json:"monthtodateheatingdegreedays"`
		MonthToDateHeatingDegreeDaysNormal string     `json:"monthtodateheatingdegreedaysnormal"`
		Since1SepHeatingDegreeDays         string     `json:"since1sepheatingdegreedays"`
		Since1SepHeatingDegreeDaysNormal   string     `json:"since1sepheatingdegreedaysnormal"`
		Since1JulHeatingDegreeDays         string     `json:"since1julheatingdegreedays"`
		Since1JulHeatingDegreeDaysNormal   string     `json:"since1julheatingdegreedaysnormal"`
		CoolingDegreeDaysNormal            string     `json:"coolingDegreeDaysnormal"`
		MonthToDateCoolingDegreeDays       string     `json:"monthtodatecoolingdegreedays"`
		MonthToDateCoolingDegreeDaysNormal string     `json:"monthtodatecoolingdegreedaysnormal"`
		Since1SepCoolingDegreeDays         string     `json:"since1sepcoolingdegreedays"`
		Since1SepCoolingDegreeDaysNormal   string     `json:"since1sepcoolingdegreedaysnormal"`
		Since1JanCoolingDegreeDays         string     `json:"since1jancoolingdegreedays"`
		Since1JanCoolingDegreeDaysNormal   string     `json:"since1jancoolingdegreedaysnormal"`
	} `json:"dailysummary"`
}

func (h *History) ToString() string {
	if len(h.Observations) == 0 {
		return "No data available for specified date"
	}

	var res []string
	for _, hist := range h.DailySummary {
		var conditions []string
		if hist.Fog == "1" {
			conditions = append(conditions, "fog")
		}
		if hist.Rain == "1" {
			conditions = append(conditions, "rain")
		}
		if hist.Snow == "1" {
			conditions = append(conditions, "snow")
		}
		if hist.Hail == "1" {
			conditions = append(conditions, "hail")
		}
		if hist.Tornado == "1" {
			conditions = append(conditions, "tornado")
		}
		res = append(res, fmt.Sprintf("Weather summary for %s: %s",
			h.Date.Pretty, strings.Join(conditions, " ")))

		if hist.Snow == "1" && hist.MonthToDateSnowfalli != "" {
			res = append(res, "   Snow:")
			if hist.Snowfalli == "T" {
				res = append(res, "     trace")
			} else if fall, _ := strconv.ParseFloat(hist.Snowfalli, 64); fall >= 0.0 {
				res = append(res, fmt.Sprintf("     %s in (%s mm)",
					hist.Snowfalli, hist.Snowfallm))
				res = append(res, fmt.Sprintf("     Snow depth: %s in (%s mm)",
					hist.SnowDepthi, hist.SnowDepthm))
				res = append(res, fmt.Sprintf("     Month to date: %s in (%s mm)",
					hist.MonthToDateSnowfalli, hist.MonthToDateSnowfallm))
				res = append(res, fmt.Sprintf("     Since July 1st: %s in (%s mm)",
					hist.Since1JulSnowfalli, hist.Since1JulSnowfallm))
			}
		}

		if hist.Rain == "1" {
			if hist.Precipi == "T" {
				res = append(res, "   Precipitation: trace")
			} else if fall, _ := strconv.ParseFloat(hist.Precipi, 64); fall > 0.0 {
				res = append(res, fmt.Sprintf("   Precipitation: %s in (%s mm)",
					hist.Precipi, hist.Precipm))
			}
		}

		res = append(res, "   Temperature:")
		res = append(res, fmt.Sprintf("      Mean Temperature: %s F (%s C)",
			hist.MeanTempi, hist.MeanTempm))
		res = append(res, fmt.Sprintf("      Max Temperature: %s F (%s C)",
			hist.MaxTempi, hist.MaxTempm))
		res = append(res, fmt.Sprintf("      Min Temperature: %s F (%s C)",
			hist.MinTempi, hist.MinTempm))

		res = append(res, "   Degree Days:")
		if hist.HeatingDegreeDays != "" {
			var normal string
			if hist.HeatingDegreeDaysNormal != "" {
				normal = " (" + hist.HeatingDegreeDaysNormal + " days normal)"
			}
			res = append(res, "      Heating Degree Days: "+hist.HeatingDegreeDays+normal)
			if hist.HeatingDegreeDaysNormal != "" && hist.HeatingDegreeDaysNormal != "0" {
				res = append(res, fmt.Sprintf("         HDG month to date: %s (%s days normal)",
					hist.MonthToDateHeatingDegreeDays, hist.MonthToDateHeatingDegreeDaysNormal))
				if hist.Since1JulHeatingDegreeDaysNormal == "" {
					res = append(res, fmt.Sprintf("         HDG since Sept 1st: %s (%s days normal)",
						hist.Since1SepHeatingDegreeDays, hist.Since1SepHeatingDegreeDaysNormal))
				} else {
					res = append(res, fmt.Sprintf("         HDG since July 1st: %s (%s days normal)",
						hist.Since1JulHeatingDegreeDays, hist.Since1JulHeatingDegreeDaysNormal))
				}
			}
		}

		if hist.CoolingDegreeDays != "" {
			var normal string
			if hist.CoolingDegreeDaysNormal != "" {
				normal = " (" + hist.CoolingDegreeDaysNormal + " days normal)"
			}
			res = append(res, "      Cooling Degree Days: "+hist.CoolingDegreeDays+normal)
			if hist.CoolingDegreeDaysNormal != "" && hist.CoolingDegreeDaysNormal != "0" {
				res = append(res, fmt.Sprintf("         CDG month to date: %s (%s days normal)",
					hist.MonthToDateCoolingDegreeDays, hist.MonthToDateCoolingDegreeDaysNormal))
				if hist.Since1JanCoolingDegreeDaysNormal == "" {
					res = append(res, fmt.Sprintf("         CDG since Sept 1st: %s (%s days normal)",
						hist.Since1SepCoolingDegreeDays, hist.Since1SepCoolingDegreeDaysNormal))
				} else {
					res = append(res, fmt.Sprintf("         CDG since Jan 1st: %s (%s days normal)",
						hist.Since1JanCoolingDegreeDays, hist.Since1JanCoolingDegreeDaysNormal))
				}
			}
		}

		res = append(res, "   Moisture:")
		res = append(res, fmt.Sprintf("      Mean Dew Point: %s (%s C)", hist.MeanDewPti, hist.MeanDewPtm))
		res = append(res, fmt.Sprintf("      Max Dew Point: %s (%s C)", hist.MaxDewPti, hist.MaxDewPtm))
		res = append(res, fmt.Sprintf("      Min Dew Point: %s (%s C)", hist.MinDewPti, hist.MinDewPtm))

		if hist.Humidity != "" {
			res = append(res, "      Humidity: "+hist.Humidity+"%")
		}
		res = append(res, "      Max Humidity: "+hist.MaxHumidity+"%")
		res = append(res, "      Min Humidity: "+hist.MinHumidity+"%")

		res = append(res, "   Pressure:")
		res = append(res, fmt.Sprintf("      Mean Pressure: %s in (%s mb)", hist.MeanPressurei, hist.MeanPressurem))
		res = append(res, fmt.Sprintf("      Max Pressure: %s in (%s mb)", hist.MaxPressurei, hist.MaxPressurem))
		res = append(res, fmt.Sprintf("      Min Pressure: %s in (%s mb)", hist.MinPressurei, hist.MinPressurem))

		res = append(res, "   Wind:")
		res = append(res, fmt.Sprintf("      Mean Wind Speed: %s mph (%s kph)", hist.MeanWindSpdi, hist.MeanWindSpdm))
		res = append(res, fmt.Sprintf("      Max Wind Speed: %s mph (%s kph)", hist.MaxWSpdi, hist.MaxWSpdm))
		res = append(res, fmt.Sprintf("      Min Wind Speed: %s mph (%s kph)", hist.MinWSpdi, hist.MinWSpdm))
		meanWDir, _ := strconv.ParseFloat(hist.MeanWDird, 64)
		res = append(res, fmt.Sprintf("      Mean Wind Direction: %s° (%s)", hist.MeanWDird, compassEnglish(meanWDir)))

		res = append(res, "   Visibility:")
		res = append(res, fmt.Sprintf("      Mean Visibility %s mi (%s km)", hist.MeanVisi, hist.MeanVism))
		res = append(res, fmt.Sprintf("      Max Visibility %s mi (%s km)", hist.MaxVisi, hist.MaxVism))
		res = append(res, fmt.Sprintf("      Min Visibility %s mi (%s km)", hist.MinVisi, hist.MinVism))
	}

	return strings.Join(res, "\n")
}

// Convert a compass degree to a human value.
func compassEnglish(degree float64) string {
	LABELS := []string{
		"N", "NNE", "NE", "ENE",
		"E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW",
		"W", "WNW", "NW", "NNW"}
	RANGE := 360.0 / float64(len(LABELS))
	return LABELS[int((degree/RANGE)+0.5)%len(LABELS)]
}
