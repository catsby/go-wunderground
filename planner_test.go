package wunderground

import (
	"encoding/json"
	"testing"
)

func TestPlannerUnmarshal(t *testing.T) {
	text := `
{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "planner": 1
  }
        }
                ,
        "trip": {
                "title": "Historical Summary for July 01 - July 31",
                "airport_code": "KSFO",
                "error": "",
                "period_of_record": {
                "date_start": {
                "date":{
        "epoch":"1341169200",
        "pretty":"12:00 PM PDT on July 01, 2012",
        "day":1,
        "month":7,
        "year":2012,
        "yday":182,
        "hour":12,
        "min":"00",
        "sec":0,
        "isdst":"1",
        "monthname":"July",
        "monthname_short":"Jul",
        "weekday_short":"Sun",
        "weekday":"Sunday",
        "ampm":"PM",
        "tz_short":"PDT",
        "tz_long":"America/Los_Angeles"
}
                },
                "date_end": {
                "date":{
        "epoch":"1501527600",
        "pretty":"12:00 PM PDT on July 31, 2017",
        "day":31,
        "month":7,
        "year":2017,
        "yday":211,
        "hour":12,
        "min":"00",
        "sec":0,
        "isdst":"1",
        "monthname":"July",
        "monthname_short":"Jul",
        "weekday_short":"Mon",
        "weekday":"Monday",
        "ampm":"PM",
        "tz_short":"PDT",
        "tz_long":"America/Los_Angeles"
}
                }
                },
                "temp_high": {
                "min": {
                "F": "63",
                "C": "17"
                },
                "avg": {
                "F": "72",
                "C": "22"
                },
                "max": {
                "F": "90",
                "C": "32"
                }
                },
                "temp_low": {
                "min": {
                "F": "51",
                "C": "11"
                },
                "avg": {
                "F": "56",
                "C": "14"
                },
                "max": {
                "F": "65",
                "C": "18"
                }
                },
                "precip": {
                "min": {
                "in": "0.00",
                "cm": "0.0"
                },
                "avg": {
                "in": "0.00",
                "cm": "0.0"
                },
                "max": {
                "in": "0.00",
                "cm": "0.0"
                }
                },
                "dewpoint_high": {
                "min": {
                "F": "49",
                "C": "9"
                },
                "avg": {
                "F": "55",
                "C": "13"
                },
                "max": {
                "F": "66",
                "C": "19"
                }
                },
                "dewpoint_low": {
                "min": {
                "F": "43",
                "C": "6"
                },
                "avg": {
                "F": "51",
                "C": "11"
                },
                "max": {
                "F": "57",
                "C": "14"
                }
                },
                "cloud_cover": {
                "cond": "mostly sunny"
                },
                "chance_of": {
                "tempoversixty": {
                "name": "Warm",
                "description": "temperature over 60&deg;F / 16&deg;C and below 90&deg;F / 32&deg;C",
                "percentage": "99"
                },
                "chanceofwindyday": {
                "name": "Windy",
                "description": "average wind over 10 mph / 15km/h",
                "percentage": "75"
                },
                "chanceofpartlycloudyday": {
                "name": "Partly Cloudy",
                "description": "",
                "percentage": "51"
                },
                "chanceofsunnycloudyday": {
                "name": "Sunny",
                "description": "",
                "percentage": "40"
                },
                "chanceofcloudyday": {
                "name": "Cloudy",
                "description": "",
                "percentage": "10"
                },
                "chanceofhumidday": {
                "name": "Humid",
                "description": "dew point over 60&deg;F / 16&deg;C",
                "percentage": "6"
                },
                "chanceoffogday": {
                "name": "Fog",
                "description": "",
                "percentage": "5"
                },
                "chanceofprecip": {
                "name": "Precipitation",
                "description": "",
                "percentage": "3"
                },
                "chanceofrainday": {
                "name": "Rain",
                "description": "",
                "percentage": "3"
                },
                "chanceofthunderday": {
                "name": "Thunderstorms",
                "description": "",
                "percentage": "1"
                },
                "tempoverninety": {
                "name": "Hot",
                "description": "temperature over 90&deg;F / 32&deg;C",
                "percentage": "1"
                },
                "chanceofsnowonground": {
                "name": "Ground Snow",
                "description": "",
                "percentage": "0"
                },
                "chanceoftornadoday": {
                "name": "Tornado",
                "description": "",
                "percentage": "0"
                },
                "chanceofsultryday": {
                "name": "Sweltering",
                "description": "dew point over 70&deg;F / 21&deg;C",
                "percentage": "0"
                },
                "tempbelowfreezing": {
                "name": "Freezing",
                "description": "temperature below 32&deg;F / 0&deg;C",
                "percentage": "0"
                },
                "tempoverfreezing": {
                "name": "Cool",
                "description": "temperature between 32&deg;F / 0&deg;C and 60&deg;F / 16&deg; C",
                "percentage": "0"
                },
                "chanceofhailday": {
                "name": "Hail",
                "description": "",
                "percentage": "0"
                },
                "chanceofsnowday": {
                "name": "Snow",
                "description": "",
                "percentage": "0"
                }
                }
        }
}
`

	var r ApiResponse
	if err := json.Unmarshal([]byte(text), &r); err != nil {
		t.Error(err)
	}
}
