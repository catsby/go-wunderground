package wunderground

import (
	"encoding/json"
	"testing"
)

func TestHistoryUnmarshal(t *testing.T) {
	text := `
{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "history": 1
  }
	}
		,
	"history": {
		"date": {
		"pretty": "April 6, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "12",
		"min": "00",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "April 6, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "19",
		"min": "00",
		"tzname": "UTC"
		},
		"observations": [
		{
		"date": {
		"pretty": "12:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "00",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "7:56 AM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "07",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.9", "tempi":"48.0","dewptm":"6.7", "dewpti":"44.1","hum":"86","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1020.6", "pressurei":"30.14","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 060756Z 00000KT 10SM FEW027 SCT200 09/07 A3014 RMK AO2 SLP206 T00890067 401440083" },
		{
		"date": {
		"pretty": "1:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "01",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "8:56 AM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "08",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"10.0", "tempi":"50.0","dewptm":"6.7", "dewpti":"44.1","hum":"80","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1020.6", "pressurei":"30.14","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 060856Z 00000KT 10SM FEW025 SCT200 10/07 A3014 RMK AO2 SLP206 T01000067 51010" },
		{
		"date": {
		"pretty": "2:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "02",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "9:56 AM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "09",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.9", "tempi":"48.0","dewptm":"6.7", "dewpti":"44.1","hum":"86","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"190","wdire":"South","vism":"16.1", "visi":"10.0","pressurem":"1020.6", "pressurei":"30.14","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 060956Z 19003KT 10SM FEW028 SCT200 09/07 A3014 RMK AO2 SLP206 T00890067" },
		{
		"date": {
		"pretty": "3:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "03",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "10:56 AM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "10",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.3", "tempi":"46.9","dewptm":"6.7", "dewpti":"44.1","hum":"90","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1020.6", "pressurei":"30.14","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061056Z 00000KT 10SM SCT044 SCT200 08/07 A3014 RMK AO2 SLP206 T00830067" },
		{
		"date": {
		"pretty": "4:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "04",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "11:56 AM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "11",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.3", "tempi":"46.9","dewptm":"6.7", "dewpti":"44.1","hum":"90","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"140","wdire":"SE","vism":"16.1", "visi":"10.0","pressurem":"1020.4", "pressurei":"30.14","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061156Z 14003KT 10SM BKN048 08/07 A3014 RMK AO2 SLP204 70024 T00830067 10100 20078 58002" },
		{
		"date": {
		"pretty": "5:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "05",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "12:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "12",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"7.8", "tempi":"46.0","dewptm":"5.6", "dewpti":"42.1","hum":"86","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"150","wdire":"SSE","vism":"16.1", "visi":"10.0","pressurem":"1020.6", "pressurei":"30.14","windchillm":"7.1", "windchilli":"44.8","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061256Z 15003KT 10SM BKN048 08/06 A3014 RMK AO2 SLP206 T00780056" },
		{
		"date": {
		"pretty": "6:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "06",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "1:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "13",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.3", "tempi":"46.9","dewptm":"6.1", "dewpti":"43.0","hum":"86","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1021.0", "pressurei":"30.15","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061356Z 00000KT 10SM FEW200 08/06 A3015 RMK AO2 SLP210 T00830061" },
		{
		"date": {
		"pretty": "7:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "07",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "2:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "14",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"8.9", "tempi":"48.0","dewptm":"7.2", "dewpti":"45.0","hum":"89","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1021.1", "pressurei":"30.16","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061456Z 00000KT 10SM FEW030 09/07 A3016 RMK AO2 SLP211 T00890072 51007" },
		{
		"date": {
		"pretty": "8:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "08",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "3:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "15",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"10.0", "tempi":"50.0","dewptm":"6.7", "dewpti":"44.1","hum":"80","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"60","wdire":"ENE","vism":"16.1", "visi":"10.0","pressurem":"1021.7", "pressurei":"30.17","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061556Z 06004KT 10SM FEW030 10/07 A3017 RMK AO2 SLP217 T01000067" },
		{
		"date": {
		"pretty": "9:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "09",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "4:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "16",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.1", "tempi":"52.0","dewptm":"7.2", "dewpti":"45.0","hum":"77","wspdm":"13.0", "wspdi":"8.1","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"100","wdire":"East","vism":"16.1", "visi":"10.0","pressurem":"1021.9", "pressurei":"30.18","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061656Z 10007KT 10SM FEW025 SCT060 11/07 A3018 RMK AO2 SLP219 T01110072" },
		{
		"date": {
		"pretty": "10:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "10",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "5:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "17",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.7", "tempi":"53.1","dewptm":"7.8", "dewpti":"46.0","hum":"77","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"80","wdire":"East","vism":"16.1", "visi":"10.0","pressurem":"1022.2", "pressurei":"30.19","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061756Z 08005KT 10SM BKN046 12/08 A3019 RMK AO2 SLP222 T01170078 10117 20078 51011" },
		{
		"date": {
		"pretty": "11:56 AM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "11",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "6:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "18",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"12.2", "tempi":"54.0","dewptm":"7.8", "dewpti":"46.0","hum":"75","wspdm":"11.1", "wspdi":"6.9","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"70","wdire":"ENE","vism":"16.1", "visi":"10.0","pressurem":"1021.8", "pressurei":"30.18","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061856Z 07006KT 10SM BKN048 12/08 A3018 RMK AO2 SLP218 T01220078" },
		{
		"date": {
		"pretty": "12:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "12",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "7:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "19",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"13.3", "tempi":"55.9","dewptm":"7.8", "dewpti":"46.0","hum":"69","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"40","wdire":"NE","vism":"16.1", "visi":"10.0","pressurem":"1021.3", "pressurei":"30.16","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 061956Z 04004KT 10SM FEW037 BKN044 13/08 A3016 RMK AO2 SLP213 T01330078" },
		{
		"date": {
		"pretty": "1:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "13",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "8:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "20",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"13.3", "tempi":"55.9","dewptm":"6.7", "dewpti":"44.1","hum":"64","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"40","wdire":"NE","vism":"16.1", "visi":"10.0","pressurem":"1020.9", "pressurei":"30.15","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 062056Z 04004KT 10SM BKN038 13/07 A3015 RMK AO2 SLP209 T01330067 57014" },
		{
		"date": {
		"pretty": "2:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "14",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "9:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "21",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"13.9", "tempi":"57.0","dewptm":"5.6", "dewpti":"42.1","hum":"57","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"300","wdire":"WNW","vism":"16.1", "visi":"10.0","pressurem":"1020.2", "pressurei":"30.13","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Overcast","icon":"cloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 062156Z 30005KT 10SM BKN042 OVC200 14/06 A3013 RMK AO2 SLP202 BINOVC T01390056" },
		{
		"date": {
		"pretty": "3:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "15",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "10:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "22",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"14.4", "tempi":"57.9","dewptm":"6.7", "dewpti":"44.1","hum":"60","wspdm":"14.8", "wspdi":"9.2","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"290","wdire":"WNW","vism":"16.1", "visi":"10.0","pressurem":"1019.4", "pressurei":"30.11","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 062256Z 29008KT 10SM SCT039 BKN200 14/07 A3011 RMK AO2 SLP194 T01440067" },
		{
		"date": {
		"pretty": "4:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "16",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "11:56 PM GMT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "23",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"14.4", "tempi":"57.9","dewptm":"5.0", "dewpti":"41.0","hum":"53","wspdm":"16.7", "wspdi":"10.4","wgustm":"25.9", "wgusti":"16.1","wdird":"270","wdire":"West","vism":"16.1", "visi":"10.0","pressurem":"1019.1", "pressurei":"30.10","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 062356Z 27009G14KT 10SM FEW035 BKN200 14/05 A3010 RMK AO2 SLP191 T01440050 10150 20117 56017 $" },
		{
		"date": {
		"pretty": "5:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "17",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "12:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "00",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"13.9", "tempi":"57.0","dewptm":"5.6", "dewpti":"42.1","hum":"57","wspdm":"18.5", "wspdi":"11.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"260","wdire":"West","vism":"16.1", "visi":"10.0","pressurem":"1018.6", "pressurei":"30.08","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070056Z 26010KT 10SM FEW040 BKN200 14/06 A3008 RMK AO2 SLP186 T01390056 $" },
		{
		"date": {
		"pretty": "6:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "18",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "1:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "01",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"12.2", "tempi":"54.0","dewptm":"5.6", "dewpti":"42.1","hum":"64","wspdm":"11.1", "wspdi":"6.9","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"270","wdire":"West","vism":"16.1", "visi":"10.0","pressurem":"1018.1", "pressurei":"30.07","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070156Z 27006KT 10SM FEW045 BKN200 12/06 A3007 RMK AO2 SLP181 T01220056 $" },
		{
		"date": {
		"pretty": "7:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "19",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "2:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "02",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.7", "tempi":"53.1","dewptm":"5.0", "dewpti":"41.0","hum":"64","wspdm":"11.1", "wspdi":"6.9","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"280","wdire":"West","vism":"16.1", "visi":"10.0","pressurem":"1017.5", "pressurei":"30.05","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070256Z 28006KT 10SM BKN200 12/05 A3005 RMK AO2 SLP175 T01170050 58017 $" },
		{
		"date": {
		"pretty": "8:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "20",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "3:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "03",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.7", "tempi":"53.1","dewptm":"5.0", "dewpti":"41.0","hum":"64","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"270","wdire":"West","vism":"16.1", "visi":"10.0","pressurem":"1017.7", "pressurei":"30.06","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070356Z 27004KT 10SM BKN200 12/05 A3005 RMK AO2 SLP177 T01170050 $" },
		{
		"date": {
		"pretty": "9:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "21",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "4:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "04",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.1", "tempi":"52.0","dewptm":"4.4", "dewpti":"39.9","hum":"63","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1017.6", "pressurei":"30.05","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070456Z 00000KT 10SM SCT200 11/04 A3005 RMK AO2 SLP176 T01110044 $" },
		{
		"date": {
		"pretty": "10:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "22",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "5:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "05",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"11.1", "tempi":"52.0","dewptm":"5.6", "dewpti":"42.1","hum":"69","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"300","wdire":"WNW","vism":"16.1", "visi":"10.0","pressurem":"1017.2", "pressurei":"30.04","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070556Z 30003KT 10SM SCT200 11/06 A3004 RMK AO2 SLP172 T01110056 10144 20106 58003 $" },
		{
		"date": {
		"pretty": "11:56 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "23",
		"min": "56",
		"tzname": "America/Los_Angeles"
		},
		"utcdate": {
		"pretty": "6:56 AM GMT on April 07, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "07",
		"hour": "06",
		"min": "56",
		"tzname": "UTC"
		},
		"tempm":"10.6", "tempi":"51.1","dewptm":"5.6", "dewpti":"42.1","hum":"71","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1016.6", "pressurei":"30.02","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 070656Z 00000KT 10SM SCT200 11/06 A3002 RMK AO2 SLP166 T01060056 $" }
		],
		"dailysummary": [
		{ "date": {
		"pretty": "12:00 PM PDT on April 06, 2006",
		"year": "2006",
		"mon": "04",
		"mday": "06",
		"hour": "12",
		"min": "00",
		"tzname": "America/Los_Angeles"
		},
		"fog":"0","rain":"0","snow":"0","snowfallm":"0.00", "snowfalli":"0.00","monthtodatesnowfallm":"", "monthtodatesnowfalli":"","since1julsnowfallm":"", "since1julsnowfalli":"","snowdepthm":"", "snowdepthi":"","hail":"0","thunder":"0","tornado":"0","meantempm":"11", "meantempi":"52","meandewptm":"6", "meandewpti":"43","meanpressurem":"1020", "meanpressurei":"30.12","meanwindspdm":"5", "meanwindspdi":"3","meanwdire":"","meanwdird":"304","meanvism":"16", "meanvisi":"10","humidity":"","maxtempm":"15", "maxtempi":"59","mintempm":"8", "mintempi":"46","maxhumidity":"90","minhumidity":"53","maxdewptm":"8", "maxdewpti":"46","mindewptm":"4", "mindewpti":"40","maxpressurem":"1022", "maxpressurei":"30.19","minpressurem":"1016", "minpressurei":"30.02","maxwspdm":"19", "maxwspdi":"12","minwspdm":"0", "minwspdi":"0","maxvism":"16", "maxvisi":"10","minvism":"16", "minvisi":"10","gdegreedays":"2","heatingdegreedays":"14","coolingdegreedays":"0","precipm":"0.00", "precipi":"0.00","precipsource":"","heatingdegreedaysnormal":"","monthtodateheatingdegreedays":"","monthtodateheatingdegreedaysnormal":"","since1sepheatingdegreedays":"","since1sepheatingdegreedaysnormal":"","since1julheatingdegreedays":"","since1julheatingdegreedaysnormal":"","coolingdegreedaysnormal":"","monthtodatecoolingdegreedays":"","monthtodatecoolingdegreedaysnormal":"","since1sepcoolingdegreedays":"","since1sepcoolingdegreedaysnormal":"","since1jancoolingdegreedays":"","since1jancoolingdegreedaysnormal":"" }
		]
	}
}
`
	var r ApiResponse
	if err := json.Unmarshal([]byte(text), &r); err != nil {
		t.Error(err)
	}
}

func TestCompassEnglish(t *testing.T) {
	tests := []struct {
		deg  float64
		want string
	}{
		{348.8, "N"},
		{10, "N"},
		{24, "NNE"},
	}
	for _, test := range tests {
		if val := compassEnglish(test.deg); val != test.want {
			t.Errorf("Failed test: %f, want=%s, got=%s", test.deg, test.want, val)
		}
	}
}
