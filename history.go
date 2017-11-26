package wunderground

import (
	"fmt"
	"time"
)

type History struct {
	Date struct {
		Pretty string `json:"pretty"`
		Year   string `json:"year"`
		Mon    string `json:"mon"`
		Mday   string `json:"mday"`
		Hour   string `json:"hour"`
		Min    string `json:"min"`
		Tzname string `json:"tzname"`
	} `json:"date"`
	Utcdate struct {
		Pretty string `json:"pretty"`
		Year   string `json:"year"`
		Mon    string `json:"mon"`
		Mday   string `json:"mday"`
		Hour   string `json:"hour"`
		Min    string `json:"min"`
		Tzname string `json:"tzname"`
	} `json:"utcdate"`
	Observations []struct {
		Date struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"date"`
		Utcdate struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"utcdate"`
		Tempm      string `json:"tempm"`
		Tempi      string `json:"tempi"`
		Dewptm     string `json:"dewptm"`
		Dewpti     string `json:"dewpti"`
		Hum        string `json:"hum"`
		Wspdm      string `json:"wspdm"`
		Wspdi      string `json:"wspdi"`
		Wgustm     string `json:"wgustm"`
		Wgusti     string `json:"wgusti"`
		Wdird      string `json:"wdird"`
		Wdire      string `json:"wdire"`
		Vism       string `json:"vism"`
		Visi       string `json:"visi"`
		Pressurem  string `json:"pressurem"`
		Pressurei  string `json:"pressurei"`
		Windchillm string `json:"windchillm"`
		Windchilli string `json:"windchilli"`
		Heatindexm string `json:"heatindexm"`
		Heatindexi string `json:"heatindexi"`
		Precipm    string `json:"precipm"`
		Precipi    string `json:"precipi"`
		Conds      string `json:"conds"`
		Icon       string `json:"icon"`
		Fog        string `json:"fog"`
		Rain       string `json:"rain"`
		Snow       string `json:"snow"`
		Hail       string `json:"hail"`
		Thunder    string `json:"thunder"`
		Tornado    string `json:"tornado"`
		Metar      string `json:"metar"`
	} `json:"observations"`
	Dailysummary []struct {
		Date struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"date"`
		Fog                                string `json:"fog"`
		Rain                               string `json:"rain"`
		Snow                               string `json:"snow"`
		Snowfallm                          string `json:"snowfallm"`
		Snowfalli                          string `json:"snowfalli"`
		Monthtodatesnowfallm               string `json:"monthtodatesnowfallm"`
		Monthtodatesnowfalli               string `json:"monthtodatesnowfalli"`
		Since1Julsnowfallm                 string `json:"since1julsnowfallm"`
		Since1Julsnowfalli                 string `json:"since1julsnowfalli"`
		Snowdepthm                         string `json:"snowdepthm"`
		Snowdepthi                         string `json:"snowdepthi"`
		Hail                               string `json:"hail"`
		Thunder                            string `json:"thunder"`
		Tornado                            string `json:"tornado"`
		Meantempm                          string `json:"meantempm"`
		Meantempi                          string `json:"meantempi"`
		Meandewptm                         string `json:"meandewptm"`
		Meandewpti                         string `json:"meandewpti"`
		Meanpressurem                      string `json:"meanpressurem"`
		Meanpressurei                      string `json:"meanpressurei"`
		Meanwindspdm                       string `json:"meanwindspdm"`
		Meanwindspdi                       string `json:"meanwindspdi"`
		Meanwdire                          string `json:"meanwdire"`
		Meanwdird                          string `json:"meanwdird"`
		Meanvism                           string `json:"meanvism"`
		Meanvisi                           string `json:"meanvisi"`
		Humidity                           string `json:"humidity"`
		Maxtempm                           string `json:"maxtempm"`
		Maxtempi                           string `json:"maxtempi"`
		Mintempm                           string `json:"mintempm"`
		Mintempi                           string `json:"mintempi"`
		Maxhumidity                        string `json:"maxhumidity"`
		Minhumidity                        string `json:"minhumidity"`
		Maxdewptm                          string `json:"maxdewptm"`
		Maxdewpti                          string `json:"maxdewpti"`
		Mindewptm                          string `json:"mindewptm"`
		Mindewpti                          string `json:"mindewpti"`
		Maxpressurem                       string `json:"maxpressurem"`
		Maxpressurei                       string `json:"maxpressurei"`
		Minpressurem                       string `json:"minpressurem"`
		Minpressurei                       string `json:"minpressurei"`
		Maxwspdm                           string `json:"maxwspdm"`
		Maxwspdi                           string `json:"maxwspdi"`
		Minwspdm                           string `json:"minwspdm"`
		Minwspdi                           string `json:"minwspdi"`
		Maxvism                            string `json:"maxvism"`
		Maxvisi                            string `json:"maxvisi"`
		Minvism                            string `json:"minvism"`
		Minvisi                            string `json:"minvisi"`
		Gdegreedays                        string `json:"gdegreedays"`
		Heatingdegreedays                  string `json:"heatingdegreedays"`
		Coolingdegreedays                  string `json:"coolingdegreedays"`
		Precipm                            string `json:"precipm"`
		Precipi                            string `json:"precipi"`
		Precipsource                       string `json:"precipsource"`
		Heatingdegreedaysnormal            string `json:"heatingdegreedaysnormal"`
		Monthtodateheatingdegreedays       string `json:"monthtodateheatingdegreedays"`
		Monthtodateheatingdegreedaysnormal string `json:"monthtodateheatingdegreedaysnormal"`
		Since1Sepheatingdegreedays         string `json:"since1sepheatingdegreedays"`
		Since1Sepheatingdegreedaysnormal   string `json:"since1sepheatingdegreedaysnormal"`
		Since1Julheatingdegreedays         string `json:"since1julheatingdegreedays"`
		Since1Julheatingdegreedaysnormal   string `json:"since1julheatingdegreedaysnormal"`
		Coolingdegreedaysnormal            string `json:"coolingdegreedaysnormal"`
		Monthtodatecoolingdegreedays       string `json:"monthtodatecoolingdegreedays"`
		Monthtodatecoolingdegreedaysnormal string `json:"monthtodatecoolingdegreedaysnormal"`
		Since1Sepcoolingdegreedays         string `json:"since1sepcoolingdegreedays"`
		Since1Sepcoolingdegreedaysnormal   string `json:"since1sepcoolingdegreedaysnormal"`
		Since1Jancoolingdegreedays         string `json:"since1jancoolingdegreedays"`
		Since1Jancoolingdegreedaysnormal   string `json:"since1jancoolingdegreedaysnormal"`
	} `json:"dailysummary"`
}

// Use to request the history feature in Service.Request
// Note: WU currently supports only a single history feature per query.
func FHistory(date *time.Time) string {
	return fmt.Sprintf("history_%s", date.Format("20070102"))
}
