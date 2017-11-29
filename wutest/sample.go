package wutest

import (
	wu "github.com/logank/go-wunderground"
	"encoding/json"
)

// Returns a new ApiResponse with everything but a planner set.
func GetTestApiResponse() *wu.ApiResponse{
	text := `
{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "alerts": 1
  ,
  "almanac": 1
  ,
  "astronomy": 1
  ,
  "conditions": 1
  ,
  "forecast": 1
  ,
  "geolookup": 1
  ,
  "yesterday": 1
  ,
  "hourly": 1
  ,
  "planner": 0
  ,
  "rawtide": 1
  ,
  "satellite": 1
  ,
  "tide": 1
  ,
  "webcams": 1
  }
                ,
        "error": {
                "type": "plannerwithhistory"
                ,"description": "Cannot load Travel Planner and History in the same request"
        }
        }
                ,       "location": {
                "type":"CITY",
                "country":"US",
                "country_iso3166":"US",
                "country_name":"USA",
                "state":"CA",
                "city":"San Francisco",
                "tz_short":"PST",
                "tz_long":"America/Los_Angeles",
                "lat":"37.77999878",
                "lon":"-122.41999817",
                "zip":"94102",
                "magic":"1",
                "wmo":"99999",
                "l":"/q/zmw:94102.1.99999",
                "requesturl":"US/CA/San_Francisco.html",
                "wuiurl":"https://www.wunderground.com/US/CA/San_Francisco.html",
                "nearby_weather_stations": {
                "airport": {
                "station": [
                { "city":"San Francisco", "state":"CA", "country":"US", "icao":"KSFO", "lat":"37.61999893", "lon":"-122.37000275" }
                ,{ "city":"Oakland", "state":"CA", "country":"US", "icao":"KOAK", "lat":"37.72000122", "lon":"-122.22000122" }
                ,{ "city":"Hayward", "state":"CA", "country":"US", "icao":"KHWD", "lat":"37.65999985", "lon":"-122.12000275" }
                ,{ "city":"Half Moon Bay", "state":"CA", "country":"US", "icao":"KHAF", "lat":"37.50999832", "lon":"-122.50000000" }
                ]
                }
                ,
                "pws": {
                "station": [
                {
                "neighborhood":"San Francisco",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1050",
                "lat":37.776947,
                "lon":-122.423325,
                "distance_km":0,
                "distance_mi":0
                },
                {
                "neighborhood":"Western Addition",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1139",
                "lat":37.781612,
                "lon":-122.431908,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"SOMA",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR131",
                "lat":37.778488,
                "lon":-122.408005,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"North Mission (Valencia and Market)",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR49",
                "lat":37.770599,
                "lon":-122.423500,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"San Francisco",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1105",
                "lat":37.781734,
                "lon":-122.432381,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"San Francisco",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1167",
                "lat":37.790237,
                "lon":-122.421356,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Nob Hill",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR575",
                "lat":37.792255,
                "lon":-122.415962,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Mission Dolores",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR371",
                "lat":37.767879,
                "lon":-122.424904,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Lower Haight",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR902",
                "lat":37.771694,
                "lon":-122.432228,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Mission Dolores",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1099",
                "lat":37.766209,
                "lon":-122.424431,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Western Addition",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1060",
                "lat":37.782047,
                "lon":-122.438248,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Mission (at Bar and Burrito)",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR142",
                "lat":37.765530,
                "lon":-122.422913,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Nob Hill",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR296",
                "lat":37.794266,
                "lon":-122.414932,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"The Castro",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR938",
                "lat":37.766399,
                "lon":-122.427544,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Mission Dolores",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR656",
                "lat":37.765980,
                "lon":-122.426208,
                "distance_km":1,
                "distance_mi":0
                },
                {
                "neighborhood":"Dolores Plaza",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR1051",
                "lat":37.765476,
                "lon":-122.424904,
                "distance_km":1,
                "distance_mi":1
                },
                {
                "neighborhood":"SoMa (7th&Brannan) (Prasad)",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR914",
                "lat":37.773605,
                "lon":-122.402252,
                "distance_km":1,
                "distance_mi":1
                },
                {
                "neighborhood":"Russian Hill at Broadway",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR989",
                "lat":37.795799,
                "lon":-122.418526,
                "distance_km":1,
                "distance_mi":1
                },
                {
                "neighborhood":"South of Market",
                "city":"San Francisco",
                "state":"CA",
                "country":"US",
                "id":"KCASANFR486",
                "lat":37.781174,
                "lon":-122.399551,
                "distance_km":1,
                "distance_mi":1
                }
                ]
                }
                }
        }
  ,     "current_observation": {
                "image": {
                "url":"http://icons.wxug.com/graphics/wu2/logo_130x80.png",
                "title":"Weather Underground",
                "link":"http://www.wunderground.com"
                },
                "display_location": {
                "full":"San Francisco, CA",
                "city":"San Francisco",
                "state":"CA",
                "state_name":"California",
                "country":"US",
                "country_iso3166":"US",
                "zip":"94102",
                "magic":"1",
                "wmo":"99999",
                "latitude":"37.77999878",
                "longitude":"-122.41999817",
                "elevation":"60.0"
                },
                "observation_location": {
                "full":"SOMA, San Francisco, California",
                "city":"SOMA, San Francisco",
                "state":"California",
                "country":"US",
                "country_iso3166":"US",
                "latitude":"37.778488",
                "longitude":"-122.408005",
                "elevation":"23 ft"
                },
                "estimated": {
                },
                "station_id":"KCASANFR131",
                "observation_time":"Last Updated on November 26, 7:24 PM PST",
                "observation_time_rfc822":"Sun, 26 Nov 2017 19:24:44 -0800",
                "observation_epoch":"1511753084",
                "local_time_rfc822":"Sun, 26 Nov 2017 19:24:46 -0800",
                "local_epoch":"1511753086",
                "local_tz_short":"PST",
                "local_tz_long":"America/Los_Angeles",
                "local_tz_offset":"-0800",
                "weather":"Overcast",
                "temperature_string":"58.0 F (14.4 C)",
                "temp_f":58.0,
                "temp_c":14.4,
                "relative_humidity":"85%",
                "wind_string":"From the SW at 2.0 MPH Gusting to 7.0 MPH",
                "wind_dir":"SW",
                "wind_degrees":220,
                "wind_mph":2.0,
                "wind_gust_mph":"7.0",
                "wind_kph":3.2,
                "wind_gust_kph":"11.3",
                "pressure_mb":"1013",
                "pressure_in":"29.93",
                "pressure_trend":"-",
                "dewpoint_string":"54 F (12 C)",
                "dewpoint_f":54,
                "dewpoint_c":12,
                "heat_index_string":"NA",
                "heat_index_f":"NA",
                "heat_index_c":"NA",
                "windchill_string":"NA",
                "windchill_f":"NA",
                "windchill_c":"NA",
                "feelslike_string":"58.0 F (14.4 C)",
                "feelslike_f":"58.0",
                "feelslike_c":"14.4",
                "visibility_mi":"10.0",
                "visibility_km":"16.1",
                "solarradiation":"0",
                "UV":"0.0","precip_1hr_string":"0.00 in ( 0 mm)",
                "precip_1hr_in":"0.00",
                "precip_1hr_metric":" 0",
                "precip_today_string":"0.00 in (0 mm)",
                "precip_today_in":"0.00",
                "precip_today_metric":"0",
                "icon":"cloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_cloudy.gif",
                "forecast_url":"http://www.wunderground.com/US/CA/San_Francisco.html",
                "history_url":"http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KCASANFR131",
                "ob_url":"http://www.wunderground.com/cgi-bin/findweather/getForecast?query=37.778488,-122.408005",
                "nowcast":""
        }
                ,       "satellite": {
                "image_url": "http://wublast.wunderground.com/cgi-bin/WUBLAST?lat=37.77999878&lon=-122.41999817&radius=75&width=300&height=300&key=sat_ir4_thumb&gtt=0&extension=png&proj=me&num=1&delay=25&timelabel=0&basemap=1&borders=1&theme=WUBLAST_WORLD&rand=1511753086&api_key=0000",
                "image_url_ir4": "http://wublast.wunderground.com/cgi-bin/WUBLAST?lat=37.77999878&lon=-122.41999817&radius=75&width=300&height=300&key=sat_ir4_thumb&gtt=0&extension=png&proj=me&num=1&delay=25&timelabel=0&basemap=1&borders=1&theme=WUBLAST_WORLD&rand=1511753086&api_key=0000",
                "image_url_vis": "http://wublast.wunderground.com/cgi-bin/WUBLAST?lat=37.77999878&lon=-122.41999817&radius=75&width=300&height=300&key=sat_vis_thumb&gtt=0&extension=png&proj=me&num=1&delay=25&timelabel=0&basemap=1&borders=1&theme=WUBLAST_WORLD&rand=1511753086&api_key=0000"
        }
                ,
        "forecast":{
                "txt_forecast": {
                "date":"6:29 PM PST",
                "forecastday": [
                {
                "period":0,
                "icon":"rain",
                "icon_url":"http://icons.wxug.com/i/c/k/rain.gif",
                "title":"Sunday",
                "fcttext":"Rain. Lows overnight in the low 50s.",
                "fcttext_metric":"Rain. Low 11C.",
                "pop":"80"
                }
                ,
                {
                "period":1,
                "icon":"nt_rain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
                "title":"Sunday Night",
                "fcttext":"Cloudy with periods of rain. Thunder possible. Low 52F. Winds WSW at 15 to 25 mph. Chance of rain 80%.",
                "fcttext_metric":"Rain. Thunder possible. Low 11C. Winds WSW at 25 to 40 km/h. Chance of rain 80%.",
                "pop":"80"
                }
                ,
                {
                "period":2,
                "icon":"clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "title":"Monday",
                "fcttext":"Mainly sunny. High 59F. Winds NW at 5 to 10 mph.",
                "fcttext_metric":"Mainly sunny. High near 15C. Winds NW at 10 to 15 km/h.",
                "pop":"10"
                }
                ,
                {
                "period":3,
                "icon":"nt_clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "title":"Monday Night",
                "fcttext":"Partly cloudy. Low 48F. Winds N at 5 to 10 mph.",
                "fcttext_metric":"A few clouds overnight. Low 9C. Winds NNE at 10 to 15 km/h.",
                "pop":"10"
                }
                ,
                {
                "period":4,
                "icon":"partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/partlycloudy.gif",
                "title":"Tuesday",
                "fcttext":"Intervals of clouds and sunshine. High 59F. Winds WNW at 5 to 10 mph.",
                "fcttext_metric":"Sunshine and clouds mixed. High around 15C. Winds WNW at 10 to 15 km/h.",
                "pop":"10"
                }
                ,
                {
                "period":5,
                "icon":"nt_clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "title":"Tuesday Night",
                "fcttext":"Mostly clear skies. Low 49F. Winds N at 5 to 10 mph.",
                "fcttext_metric":"Clear skies. Low 9C. Winds N at 10 to 15 km/h.",
                "pop":"10"
                }
                ,
                {
                "period":6,
                "icon":"clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "title":"Wednesday",
                "fcttext":"Sunny. High 63F. Winds NNE at 5 to 10 mph.",
                "fcttext_metric":"Sunny skies. High 18C. Winds NNE at 10 to 15 km/h.",
                "pop":"0"
                }
                ,
                {
                "period":7,
                "icon":"nt_clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "title":"Wednesday Night",
                "fcttext":"Clear skies. Low 48F. Winds light and variable.",
                "fcttext_metric":"A mostly clear sky. Low 9C. Winds light and variable.",
                "pop":"0"
                }
                ]
                },
                "simpleforecast": {
                "forecastday": [
                {"date":{
        "epoch":"1511751600",
        "pretty":"7:00 PM PST on November 26, 2017",
        "day":26,
        "month":11,
        "year":2017,
        "yday":329,
        "hour":19,
        "min":"00",
        "sec":0,
        "isdst":"0",
        "monthname":"November",
        "monthname_short":"Nov",
        "weekday_short":"Sun",
        "weekday":"Sunday",
        "ampm":"PM",
        "tz_short":"PST",
        "tz_long":"America/Los_Angeles"
},
                "period":1,
                "high": {
                "fahrenheit":"65",
                "celsius":"18"
                },
                "low": {
                "fahrenheit":"52",
                "celsius":"11"
                },
                "conditions":"Rain",
                "icon":"rain",
                "icon_url":"http://icons.wxug.com/i/c/k/rain.gif",
                "skyicon":"",
                "pop":80,
                "qpf_allday": {
                "in": 0.20,
                "mm": 5
                },
                "qpf_day": {
                "in": null,
                "mm": null
                },
                "qpf_night": {
                "in": 0.20,
                "mm": 5
                },
                "snow_allday": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_day": {
                "in": null,
                "cm": null
                },
                "snow_night": {
                "in": 0.0,
                "cm": 0.0
                },
                "maxwind": {
                "mph": 19,
                "kph": 31,
                "dir": "",
                "degrees": 0
                },
                "avewind": {
                "mph": 3,
                "kph": 5,
                "dir": "SSW",
                "degrees": 202
                },
                "avehumidity": 80,
                "maxhumidity": 86,
                "minhumidity": 77
                }
                ,
                {"date":{
        "epoch":"1511838000",
        "pretty":"7:00 PM PST on November 27, 2017",
        "day":27,
        "month":11,
        "year":2017,
        "yday":330,
        "hour":19,
        "min":"00",
        "sec":0,
        "isdst":"0",
        "monthname":"November",
        "monthname_short":"Nov",
        "weekday_short":"Mon",
        "weekday":"Monday",
        "ampm":"PM",
        "tz_short":"PST",
        "tz_long":"America/Los_Angeles"
},
                "period":2,
                "high": {
                "fahrenheit":"59",
                "celsius":"15"
                },
                "low": {
                "fahrenheit":"48",
                "celsius":"9"
                },
                "conditions":"Clear",
                "icon":"clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "skyicon":"",
                "pop":10,
                "qpf_allday": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_day": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_night": {
                "in": 0.00,
                "mm": 0
                },
                "snow_allday": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_day": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_night": {
                "in": 0.0,
                "cm": 0.0
                },
                "maxwind": {
                "mph": 10,
                "kph": 16,
                "dir": "NW",
                "degrees": 315
                },
                "avewind": {
                "mph": 7,
                "kph": 11,
                "dir": "NW",
                "degrees": 315
                },
                "avehumidity": 71,
                "maxhumidity": 83,
                "minhumidity": 60
                }
                ,
                {"date":{
        "epoch":"1511924400",
        "pretty":"7:00 PM PST on November 28, 2017",
        "day":28,
        "month":11,
        "year":2017,
        "yday":331,
        "hour":19,
        "min":"00",
        "sec":0,
        "isdst":"0",
        "monthname":"November",
        "monthname_short":"Nov",
        "weekday_short":"Tue",
        "weekday":"Tuesday",
        "ampm":"PM",
        "tz_short":"PST",
        "tz_long":"America/Los_Angeles"
},
                "period":3,
                "high": {
                "fahrenheit":"59",
                "celsius":"15"
                },
                "low": {
                "fahrenheit":"49",
                "celsius":"9"
                },
                "conditions":"Partly Cloudy",
                "icon":"partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/partlycloudy.gif",
                "skyicon":"",
                "pop":10,
                "qpf_allday": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_day": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_night": {
                "in": 0.00,
                "mm": 0
                },
                "snow_allday": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_day": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_night": {
                "in": 0.0,
                "cm": 0.0
                },
                "maxwind": {
                "mph": 10,
                "kph": 16,
                "dir": "WNW",
                "degrees": 290
                },
                "avewind": {
                "mph": 8,
                "kph": 13,
                "dir": "WNW",
                "degrees": 290
                },
                "avehumidity": 79,
                "maxhumidity": 89,
                "minhumidity": 68
                }
                ,
                {"date":{
        "epoch":"1512010800",
        "pretty":"7:00 PM PST on November 29, 2017",
        "day":29,
        "month":11,
        "year":2017,
        "yday":332,
        "hour":19,
        "min":"00",
        "sec":0,
        "isdst":"0",
        "monthname":"November",
        "monthname_short":"Nov",
        "weekday_short":"Wed",
        "weekday":"Wednesday",
        "ampm":"PM",
        "tz_short":"PST",
        "tz_long":"America/Los_Angeles"
},
                "period":4,
                "high": {
                "fahrenheit":"63",
                "celsius":"17"
                },
                "low": {
                "fahrenheit":"48",
                "celsius":"9"
                },
                "conditions":"Clear",
                "icon":"clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "skyicon":"",
                "pop":0,
                "qpf_allday": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_day": {
                "in": 0.00,
                "mm": 0
                },
                "qpf_night": {
                "in": 0.00,
                "mm": 0
                },
                "snow_allday": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_day": {
                "in": 0.0,
                "cm": 0.0
                },
                "snow_night": {
                "in": 0.0,
                "cm": 0.0
                },
                "maxwind": {
                "mph": 10,
                "kph": 16,
                "dir": "NNE",
                "degrees": 14
                },
                "avewind": {
                "mph": 7,
                "kph": 11,
                "dir": "NNE",
                "degrees": 14
                },
                "avehumidity": 65,
                "maxhumidity": 81,
                "minhumidity": 45
                }
                ]
                }
        }
                ,
        "hourly_forecast": [
                {
                "FCTTIME": {
                "hour": "20","hour_padded": "20","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "26","mday_padded": "26","yday": "329","isdst": "0","epoch": "1511755200","pretty": "8:00 PM PST on November 26, 2017","civil": "8:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Sunday","weekday_name_night": "Sunday Night","weekday_name_abbrev": "Sun","weekday_name_unlang": "Sunday","weekday_name_night_unlang": "Sunday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "53", "metric": "12"},
                "condition": "Mostly Cloudy",
                "icon": "mostlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_mostlycloudy.gif",
                "fctcode": "3",
                "sky": "79",
                "wspd": {"english": "12", "metric": "19"},
                "wdir": {"dir": "WSW", "degrees": "238"},
                "wx": "Mostly Cloudy",
                "uvi": "0",
                "humidity": "80",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "2",
                "mslp": {"english": "29.96", "metric": "1015"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "21","hour_padded": "21","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "26","mday_padded": "26","yday": "329","isdst": "0","epoch": "1511758800","pretty": "9:00 PM PST on November 26, 2017","civil": "9:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Sunday","weekday_name_night": "Sunday Night","weekday_name_abbrev": "Sun","weekday_name_unlang": "Sunday","weekday_name_night_unlang": "Sunday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "53", "metric": "12"},
                "condition": "Mostly Cloudy",
                "icon": "mostlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_mostlycloudy.gif",
                "fctcode": "3",
                "sky": "68",
                "wspd": {"english": "11", "metric": "18"},
                "wdir": {"dir": "SW", "degrees": "230"},
                "wx": "Mostly Cloudy",
                "uvi": "0",
                "humidity": "80",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "29.95", "metric": "1014"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "22","hour_padded": "22","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "26","mday_padded": "26","yday": "329","isdst": "0","epoch": "1511762400","pretty": "10:00 PM PST on November 26, 2017","civil": "10:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Sunday","weekday_name_night": "Sunday Night","weekday_name_abbrev": "Sun","weekday_name_unlang": "Sunday","weekday_name_night_unlang": "Sunday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "52", "metric": "11"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "59",
                "wspd": {"english": "13", "metric": "21"},
                "wdir": {"dir": "SW", "degrees": "229"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "79",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "0",
                "mslp": {"english": "29.95", "metric": "1014"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "23","hour_padded": "23","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "26","mday_padded": "26","yday": "329","isdst": "0","epoch": "1511766000","pretty": "11:00 PM PST on November 26, 2017","civil": "11:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Sunday","weekday_name_night": "Sunday Night","weekday_name_abbrev": "Sun","weekday_name_unlang": "Sunday","weekday_name_night_unlang": "Sunday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "58", "metric": "14"},
                "dewpoint": {"english": "51", "metric": "11"},
                "condition": "Chance of Rain",
                "icon": "chancerain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
                "fctcode": "12",
                "sky": "50",
                "wspd": {"english": "14", "metric": "23"},
                "wdir": {"dir": "WSW", "degrees": "239"},
                "wx": "Showers",
                "uvi": "0",
                "humidity": "77",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "58", "metric": "14"},
                "qpf": {"english": "0.08", "metric": "2"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "47",
                "mslp": {"english": "29.95", "metric": "1014"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "0","hour_padded": "00","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511769600","pretty": "12:00 AM PST on November 27, 2017","civil": "12:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "58", "metric": "14"},
                "dewpoint": {"english": "51", "metric": "11"},
                "condition": "Rain",
                "icon": "rain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
                "fctcode": "13",
                "sky": "100",
                "wspd": {"english": "19", "metric": "31"},
                "wdir": {"dir": "WSW", "degrees": "255"},
                "wx": "Rain",
                "uvi": "0",
                "humidity": "78",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "58", "metric": "14"},
                "qpf": {"english": "0.04", "metric": "1"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "73",
                "mslp": {"english": "29.96", "metric": "1015"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "1","hour_padded": "01","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511773200","pretty": "1:00 AM PST on November 27, 2017","civil": "1:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "48", "metric": "9"},
                "condition": "Rain",
                "icon": "rain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
                "fctcode": "13",
                "sky": "100",
                "wspd": {"english": "16", "metric": "26"},
                "wdir": {"dir": "W", "degrees": "266"},
                "wx": "Rain",
                "uvi": "0",
                "humidity": "77",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.03", "metric": "1"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "76",
                "mslp": {"english": "29.97", "metric": "1015"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "2","hour_padded": "02","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511776800","pretty": "2:00 AM PST on November 27, 2017","civil": "2:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Rain",
                "icon": "rain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
                "fctcode": "13",
                "sky": "100",
                "wspd": {"english": "15", "metric": "24"},
                "wdir": {"dir": "W", "degrees": "266"},
                "wx": "Rain",
                "uvi": "0",
                "humidity": "74",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.02", "metric": "1"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "67",
                "mslp": {"english": "29.97", "metric": "1015"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "3","hour_padded": "03","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511780400","pretty": "3:00 AM PST on November 27, 2017","civil": "3:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Chance of Rain",
                "icon": "chancerain",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
                "fctcode": "12",
                "sky": "62",
                "wspd": {"english": "15", "metric": "24"},
                "wdir": {"dir": "W", "degrees": "269"},
                "wx": "Showers",
                "uvi": "0",
                "humidity": "75",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.01", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "45",
                "mslp": {"english": "29.99", "metric": "1016"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "4","hour_padded": "04","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511784000","pretty": "4:00 AM PST on November 27, 2017","civil": "4:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Mostly Cloudy",
                "icon": "mostlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_mostlycloudy.gif",
                "fctcode": "3",
                "sky": "62",
                "wspd": {"english": "14", "metric": "23"},
                "wdir": {"dir": "W", "degrees": "275"},
                "wx": "Mostly Cloudy",
                "uvi": "0",
                "humidity": "73",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "22",
                "mslp": {"english": "30.01", "metric": "1016"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "5","hour_padded": "05","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511787600","pretty": "5:00 AM PST on November 27, 2017","civil": "5:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "54", "metric": "12"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "38",
                "wspd": {"english": "12", "metric": "19"},
                "wdir": {"dir": "W", "degrees": "281"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "75",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "54", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "15",
                "mslp": {"english": "30.03", "metric": "1017"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "6","hour_padded": "06","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511791200","pretty": "6:00 AM PST on November 27, 2017","civil": "6:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "53", "metric": "12"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "30",
                "wspd": {"english": "10", "metric": "16"},
                "wdir": {"dir": "WNW", "degrees": "291"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "74",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "53", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "4",
                "mslp": {"english": "30.06", "metric": "1018"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "7","hour_padded": "07","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511794800","pretty": "7:00 AM PST on November 27, 2017","civil": "7:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "53", "metric": "12"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "19",
                "wspd": {"english": "7", "metric": "11"},
                "wdir": {"dir": "WNW", "degrees": "297"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "75",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "53", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "5",
                "mslp": {"english": "30.08", "metric": "1019"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "8","hour_padded": "08","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511798400","pretty": "8:00 AM PST on November 27, 2017","civil": "8:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "53", "metric": "12"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "18",
                "wspd": {"english": "5", "metric": "8"},
                "wdir": {"dir": "WNW", "degrees": "301"},
                "wx": "Sunny",
                "uvi": "0",
                "humidity": "74",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "53", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "3",
                "mslp": {"english": "30.11", "metric": "1020"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "9","hour_padded": "09","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511802000","pretty": "9:00 AM PST on November 27, 2017","civil": "9:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "8",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "NW", "degrees": "313"},
                "wx": "Sunny",
                "uvi": "1",
                "humidity": "69",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "2",
                "mslp": {"english": "30.14", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "10","hour_padded": "10","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511805600","pretty": "10:00 AM PST on November 27, 2017","civil": "10:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "56", "metric": "13"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "5",
                "wspd": {"english": "7", "metric": "11"},
                "wdir": {"dir": "WNW", "degrees": "302"},
                "wx": "Sunny",
                "uvi": "2",
                "humidity": "68",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "56", "metric": "13"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "2",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "11","hour_padded": "11","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511809200","pretty": "11:00 AM PST on November 27, 2017","civil": "11:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "57", "metric": "14"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "8",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "WNW", "degrees": "303"},
                "wx": "Sunny",
                "uvi": "2",
                "humidity": "66",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "57", "metric": "14"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "30.18", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "12","hour_padded": "12","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511812800","pretty": "12:00 PM PST on November 27, 2017","civil": "12:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "57", "metric": "14"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "18",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "NW", "degrees": "312"},
                "wx": "Sunny",
                "uvi": "3",
                "humidity": "64",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "57", "metric": "14"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "13","hour_padded": "13","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511816400","pretty": "1:00 PM PST on November 27, 2017","civil": "1:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "58", "metric": "14"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "6",
                "wspd": {"english": "5", "metric": "8"},
                "wdir": {"dir": "NW", "degrees": "321"},
                "wx": "Sunny",
                "uvi": "2",
                "humidity": "63",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "58", "metric": "14"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "14","hour_padded": "14","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511820000","pretty": "2:00 PM PST on November 27, 2017","civil": "2:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "7",
                "wspd": {"english": "5", "metric": "8"},
                "wdir": {"dir": "NNW", "degrees": "331"},
                "wx": "Sunny",
                "uvi": "1",
                "humidity": "60",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "0",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "15","hour_padded": "15","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511823600","pretty": "3:00 PM PST on November 27, 2017","civil": "3:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "1",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "NW", "degrees": "317"},
                "wx": "Sunny",
                "uvi": "1",
                "humidity": "61",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "0",
                "mslp": {"english": "30.15", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "16","hour_padded": "16","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511827200","pretty": "4:00 PM PST on November 27, 2017","civil": "4:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "59", "metric": "15"},
                "dewpoint": {"english": "45", "metric": "7"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/clear.gif",
                "fctcode": "1",
                "sky": "1",
                "wspd": {"english": "7", "metric": "11"},
                "wdir": {"dir": "NW", "degrees": "319"},
                "wx": "Sunny",
                "uvi": "0",
                "humidity": "61",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "59", "metric": "15"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "30.15", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "17","hour_padded": "17","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511830800","pretty": "5:00 PM PST on November 27, 2017","civil": "5:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "57", "metric": "14"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "0",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "NNW", "degrees": "328"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "65",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "57", "metric": "14"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "1",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "18","hour_padded": "18","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511834400","pretty": "6:00 PM PST on November 27, 2017","civil": "6:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "56", "metric": "13"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "0",
                "wspd": {"english": "7", "metric": "11"},
                "wdir": {"dir": "NNW", "degrees": "337"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "70",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "56", "metric": "13"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "2",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "19","hour_padded": "19","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511838000","pretty": "7:00 PM PST on November 27, 2017","civil": "7:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "55", "metric": "13"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "0",
                "wspd": {"english": "6", "metric": "10"},
                "wdir": {"dir": "NNW", "degrees": "337"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "75",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "55", "metric": "13"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "3",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "20","hour_padded": "20","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511841600","pretty": "8:00 PM PST on November 27, 2017","civil": "8:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "54", "metric": "12"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "1",
                "wspd": {"english": "5", "metric": "8"},
                "wdir": {"dir": "NNW", "degrees": "328"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "77",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "54", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "4",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "21","hour_padded": "21","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511845200","pretty": "9:00 PM PST on November 27, 2017","civil": "9:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "53", "metric": "12"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "1",
                "wspd": {"english": "4", "metric": "6"},
                "wdir": {"dir": "NNW", "degrees": "345"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "79",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "53", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "5",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "22","hour_padded": "22","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511848800","pretty": "10:00 PM PST on November 27, 2017","civil": "10:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "53", "metric": "12"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "2",
                "wspd": {"english": "3", "metric": "5"},
                "wdir": {"dir": "NNW", "degrees": "336"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "81",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "53", "metric": "12"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "5",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "23","hour_padded": "23","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "27","mday_padded": "27","yday": "330","isdst": "0","epoch": "1511852400","pretty": "11:00 PM PST on November 27, 2017","civil": "11:00 PM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Monday","weekday_name_night": "Monday Night","weekday_name_abbrev": "Mon","weekday_name_unlang": "Monday","weekday_name_night_unlang": "Monday Night","ampm": "PM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "52", "metric": "11"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "3",
                "wspd": {"english": "3", "metric": "5"},
                "wdir": {"dir": "NW", "degrees": "322"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "83",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "52", "metric": "11"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "5",
                "mslp": {"english": "30.17", "metric": "1022"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "0","hour_padded": "00","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511856000","pretty": "12:00 AM PST on November 28, 2017","civil": "12:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "52", "metric": "11"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "4",
                "wspd": {"english": "1", "metric": "2"},
                "wdir": {"dir": "NNW", "degrees": "348"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "84",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "52", "metric": "11"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "5",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "1","hour_padded": "01","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511859600","pretty": "1:00 AM PST on November 28, 2017","civil": "1:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "52", "metric": "11"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "6",
                "wspd": {"english": "2", "metric": "3"},
                "wdir": {"dir": "N", "degrees": "8"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "85",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "52", "metric": "11"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "6",
                "mslp": {"english": "30.16", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "2","hour_padded": "02","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511863200","pretty": "2:00 AM PST on November 28, 2017","civil": "2:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "51", "metric": "11"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "10",
                "wspd": {"english": "2", "metric": "3"},
                "wdir": {"dir": "E", "degrees": "97"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "87",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "51", "metric": "11"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "6",
                "mslp": {"english": "30.15", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "3","hour_padded": "03","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511866800","pretty": "3:00 AM PST on November 28, 2017","civil": "3:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "50", "metric": "10"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "14",
                "wspd": {"english": "2", "metric": "3"},
                "wdir": {"dir": "E", "degrees": "86"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "89",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "50", "metric": "10"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "7",
                "mslp": {"english": "30.14", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "4","hour_padded": "04","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511870400","pretty": "4:00 AM PST on November 28, 2017","civil": "4:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "50", "metric": "10"},
                "dewpoint": {"english": "47", "metric": "8"},
                "condition": "Clear",
                "icon": "clear",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
                "fctcode": "1",
                "sky": "18",
                "wspd": {"english": "2", "metric": "3"},
                "wdir": {"dir": "S", "degrees": "172"},
                "wx": "Clear",
                "uvi": "0",
                "humidity": "89",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "50", "metric": "10"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "7",
                "mslp": {"english": "30.14", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "5","hour_padded": "05","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511874000","pretty": "5:00 AM PST on November 28, 2017","civil": "5:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "49", "metric": "9"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "36",
                "wspd": {"english": "3", "metric": "5"},
                "wdir": {"dir": "SSE", "degrees": "160"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "89",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "49", "metric": "9"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "8",
                "mslp": {"english": "30.14", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "6","hour_padded": "06","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511877600","pretty": "6:00 AM PST on November 28, 2017","civil": "6:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "49", "metric": "9"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "35",
                "wspd": {"english": "3", "metric": "5"},
                "wdir": {"dir": "S", "degrees": "187"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "88",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "49", "metric": "9"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "7",
                "mslp": {"english": "30.14", "metric": "1021"}
                }
                ,
                {
                "FCTTIME": {
                "hour": "7","hour_padded": "07","min": "00","min_unpadded": "0","sec": "0","year": "2017","mon": "11","mon_padded": "11","mon_abbrev": "Nov","mday": "28","mday_padded": "28","yday": "331","isdst": "0","epoch": "1511881200","pretty": "7:00 AM PST on November 28, 2017","civil": "7:00 AM","month_name": "November","month_name_abbrev": "Nov","weekday_name": "Tuesday","weekday_name_night": "Tuesday Night","weekday_name_abbrev": "Tue","weekday_name_unlang": "Tuesday","weekday_name_night_unlang": "Tuesday Night","ampm": "AM","tz": "","age": "","UTCDATE": ""
                },
                "temp": {"english": "50", "metric": "10"},
                "dewpoint": {"english": "46", "metric": "8"},
                "condition": "Partly Cloudy",
                "icon": "partlycloudy",
                "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
                "fctcode": "2",
                "sky": "33",
                "wspd": {"english": "3", "metric": "5"},
                "wdir": {"dir": "WSW", "degrees": "250"},
                "wx": "Partly Cloudy",
                "uvi": "0",
                "humidity": "88",
                "windchill": {"english": "-9999", "metric": "-9999"},
                "heatindex": {"english": "-9999", "metric": "-9999"},
                "feelslike": {"english": "50", "metric": "10"},
                "qpf": {"english": "0.0", "metric": "0"},
                "snow": {"english": "0.0", "metric": "0"},
                "pop": "6",
                "mslp": {"english": "30.15", "metric": "1021"}
                }
        ]
                ,       "moon_phase": {
                "percentIlluminated":"55",
                "ageOfMoon":"8",
                "phaseofMoon":"First Quarter",
                "hemisphere":"North",
                "current_time": {
                "hour":"19",
                "minute":"24"
                },
                "sunrise": {
                "hour":"7",
                "minute":"01"
                },
                "sunset": {
                "hour":"16",
                "minute":"52"
                },
                "moonrise": {
                "hour":"12",
                "minute":"56"
                },
                "moonset": {
                "hour":"",
                "minute":""
                }
        },
        "sun_phase": {
                "sunrise": {
                "hour":"7",
                "minute":"01"
                },
                "sunset": {
                "hour":"16",
                "minute":"52"
                }
        }
                ,"query_zone": "006",
	"alerts": [{
	  "type": "HEA",
	  "description": "Heat Advisory",
	  "date": "11:14 am CDT on July 3, 2012",
	  "date_epoch": "1341332040",
	  "expires": "7:00 AM CDT on July 07, 2012",
	  "expires_epoch": "1341662400",
	  "message": "\u000A...Heat advisory remains in effect until 7 am CDT Saturday...\u000A\u000A* temperature...heat indices of 100 to 105 are expected each \u000A afternoon...as Max temperatures climb into the mid to upper \u000A 90s...combined with dewpoints in the mid 60s to around 70. \u000A Heat indices will remain in the 75 to 80 degree range at \u000A night. \u000A\u000A* Impacts...the hot and humid weather will lead to an increased \u000A risk of heat related stress and illnesses. \u000A\u000APrecautionary/preparedness actions...\u000A\u000AA heat advisory means that a period of hot temperatures is\u000Aexpected. The combination of hot temperatures and high humidity\u000Awill combine to create a situation in which heat illnesses are\u000Apossible. Drink plenty of fluids...stay in an air-conditioned\u000Aroom...stay out of the sun...and check up on relatives...pets...\u000Aneighbors...and livestock.\u000A\u000ATake extra precautions if you work or spend time outside. Know\u000Athe signs and symptoms of heat exhaustion and heat stroke. Anyone\u000Aovercome by heat should be moved to a cool and shaded location.\u000AHeat stroke is an emergency...call 9 1 1.\u000A\u000A\u000A\u000AMjb\u000A\u000A\u000A",
	  "phenomena": "HT",
	  "significance": "Y",
	  "ZONES": [
	  {
	  "state":"UT",
	  "ZONE":"001"
	  }
	  ],
	  "StormBased": {
	  "vertices":[
	  {
	  "lat":"38.87",
	  "lon":"-87.13"
	  }
	  ,
	  {
	  "lat":"38.89",
	  "lon":"-87.13"
	  }
	  ,
	  {
	  "lat":"38.91",
	  "lon":"-87.11"
	  }
	  ,
	  {
	  "lat":"38.98",
	  "lon":"-86.93"
	  }
	  ,
	  {
	  "lat":"38.87",
	  "lon":"-86.69"
	  }
	  ,
	  {
	  "lat":"38.75",
	  "lon":"-86.3"
	  }
	  ,
	  {
	  "lat":"38.84",
	  "lon":"-87.16"
	  }
	  ],
	  "Vertex_count":7,
	  "stormInfo": {
	  "time_epoch": 1363464360,
	  "Motion_deg": 243,
	  "Motion_spd": 18,
	  "position_lat":38.90,
	  "position_lon":-86.96
	  }
	  }
	  }]
                ,
        "history": {
                "date": {
                "pretty": "November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "12",
                "min": "00",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "20",
                "min": "00",
                "tzname": "UTC"
                },
                "observations": [
                {
                "date": {
                "pretty": "12:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "00",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "8:56 AM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "08",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"15.6", "tempi":"60.1","dewptm":"12.2", "dewpti":"54.0","hum":"80","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"170","wdire":"South","vism":"16.1", "visi":"10.0","pressurem":"1015.5", "pressurei":"29.99","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"1","rain":"0","snow":"1","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 250856Z 17005KT 10SM BKN150 16/12 A2999 RMK AO2 SLP155 T01560122 58015 $" },
                {
                "date": {
                "pretty": "1:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "01",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "9:56 AM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "09",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"15.6", "tempi":"60.1","dewptm":"12.2", "dewpti":"54.0","hum":"80","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1015.3", "pressurei":"29.99","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"1","thunder":"0","tornado":"0","metar":"METAR KSFO 250956Z 00000KT 10SM SCT150 16/12 A2998 RMK AO2 SLP153 T01560122 $" },
                {
                "date": {
                "pretty": "2:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "02",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "10:56 AM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "10",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"15.6", "tempi":"60.1","dewptm":"13.3", "dewpti":"55.9","hum":"86","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"50","wdire":"NE","vism":"16.1", "visi":"10.0","pressurem":"1014.8", "pressurei":"29.97","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251056Z 05005KT 10SM FEW150 16/13 A2997 RMK AO2 SLP148 T01560133 $" },
                {
                "date": {
                "pretty": "3:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "03",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "11:56 AM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "11",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"13.9", "tempi":"57.0","dewptm":"12.2", "dewpti":"54.0","hum":"89","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"160","wdire":"SSE","vism":"16.1", "visi":"10.0","pressurem":"1014.5", "pressurei":"29.96","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251156Z 16004KT 10SM FEW150 14/12 A2996 RMK AO2 SLP145 T01390122 10172 20139 58011 $" },
                {
                "date": {
                "pretty": "4:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "04",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "12:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "12",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"14.4", "tempi":"57.9","dewptm":"11.7", "dewpti":"53.1","hum":"84","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"130","wdire":"SE","vism":"16.1", "visi":"10.0","pressurem":"1013.9", "pressurei":"29.94","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251256Z 13004KT 10SM FEW150 14/12 A2994 RMK AO2 SLP139 T01440117 $" },
                {
                "date": {
                "pretty": "5:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "05",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "1:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "13",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"14.4", "tempi":"57.9","dewptm":"11.7", "dewpti":"53.1","hum":"84","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"120","wdire":"ESE","vism":"16.1", "visi":"10.0","pressurem":"1014.0", "pressurei":"29.95","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251356Z 12004KT 10SM FEW180 14/12 A2995 RMK AO2 SLP140 T01440117 $" },
                {
                "date": {
                "pretty": "6:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "06",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "2:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "14",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"13.9", "tempi":"57.0","dewptm":"11.7", "dewpti":"53.1","hum":"87","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1014.5", "pressurei":"29.96","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251456Z 00000KT 10SM SCT200 14/12 A2996 RMK AO2 SLP145 T01390117 53000 $" },
                {
                "date": {
                "pretty": "7:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "07",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "3:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "15",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"14.4", "tempi":"57.9","dewptm":"12.2", "dewpti":"54.0","hum":"87","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"180","wdire":"South","vism":"16.1", "visi":"10.0","pressurem":"1014.9", "pressurei":"29.97","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251556Z 18003KT 10SM FEW180 BKN200 14/12 A2997 RMK AO2 SLP149 T01440122 $" },
                {
                "date": {
                "pretty": "8:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "08",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "4:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "16",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"16.1", "tempi":"61.0","dewptm":"12.2", "dewpti":"54.0","hum":"78","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1015.3", "pressurei":"29.99","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Overcast","icon":"cloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251656Z 00000KT 10SM FEW150 SCT180 OVC200 16/12 A2998 RMK AO2 SLP153 T01610122 $" },
                {
                "date": {
                "pretty": "9:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "09",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "5:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "17",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"16.7", "tempi":"62.1","dewptm":"12.2", "dewpti":"54.0","hum":"75","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"70","wdire":"ENE","vism":"16.1", "visi":"10.0","pressurem":"1015.3", "pressurei":"29.99","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Overcast","icon":"cloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251756Z 07003KT 10SM FEW150 SCT170 OVC200 17/12 A2998 RMK AO2 SLP153 T01670122 10172 20133 51008 $" },
                {
                "date": {
                "pretty": "10:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "10",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "6:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "18",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"16.7", "tempi":"62.1","dewptm":"13.3", "dewpti":"55.9","hum":"80","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"60","wdire":"ENE","vism":"16.1", "visi":"10.0","pressurem":"1014.8", "pressurei":"29.97","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Overcast","icon":"cloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251856Z 06003KT 10SM FEW150 SCT170 OVC200 17/13 A2997 RMK AO2 SLP148 T01670133 $" },
                {
                "date": {
                "pretty": "11:56 AM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "11",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "7:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "19",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.2", "tempi":"63.0","dewptm":"13.9", "dewpti":"57.0","hum":"81","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"80","wdire":"East","vism":"16.1", "visi":"10.0","pressurem":"1014.0", "pressurei":"29.95","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 251956Z 08004KT 10SM FEW150 SCT170 BKN200 17/14 A2995 RMK AO2 SLP140 T01720139 $" },
                {
                "date": {
                "pretty": "12:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "12",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "8:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "20",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.2", "tempi":"63.0","dewptm":"12.8", "dewpti":"55.0","hum":"75","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"350","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1014.2", "pressurei":"29.95","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 252056Z 35004KT 10SM FEW160 SCT180 BKN200 17/13 A2995 RMK AO2 SLP142 T01720128 55011 $" },
                {
                "date": {
                "pretty": "1:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "13",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "9:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "21",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"18.9", "tempi":"66.0","dewptm":"13.9", "dewpti":"57.0","hum":"73","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1013.9", "pressurei":"29.94","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 252156Z 00000KT 10SM FEW160 SCT180 BKN200 19/14 A2994 RMK AO2 SLP139 T01890139 $" },
                {
                "date": {
                "pretty": "2:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "14",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "10:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "22",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"18.3", "tempi":"64.9","dewptm":"13.9", "dewpti":"57.0","hum":"75","wspdm":"16.7", "wspdi":"10.4","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"10","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1013.9", "pressurei":"29.94","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 252256Z 01009KT 10SM BKN150 BKN200 18/14 A2994 RMK AO2 SLP139 T01830139 $" },
                {
                "date": {
                "pretty": "3:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "15",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "11:56 PM GMT on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "23",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.8", "tempi":"64.0","dewptm":"12.8", "dewpti":"55.0","hum":"73","wspdm":"5.6", "wspdi":"3.5","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"120","wdire":"ESE","vism":"16.1", "visi":"10.0","pressurem":"1013.6", "pressurei":"29.94","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 252356Z 12003KT 10SM BKN150 BKN200 18/13 A2993 RMK AO2 SLP136 T01780128 10194 20161 58006 $" },
                {
                "date": {
                "pretty": "4:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "16",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "12:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "00",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"18.3", "tempi":"64.9","dewptm":"12.2", "dewpti":"54.0","hum":"68","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1013.7", "pressurei":"29.94","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260056Z 00000KT 10SM SCT150 BKN180 18/12 A2994 RMK AO2 SLP137 T01830122 $" },
                {
                "date": {
                "pretty": "5:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "17",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "1:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "01",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"18.3", "tempi":"64.9","dewptm":"12.2", "dewpti":"54.0","hum":"68","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"330","wdire":"NNW","vism":"16.1", "visi":"10.0","pressurem":"1014.1", "pressurei":"29.95","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260156Z 33004KT 10SM BKN150 18/12 A2995 RMK AO2 SLP141 T01830122 $" },
                {
                "date": {
                "pretty": "6:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "18",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "2:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "02",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.2", "tempi":"63.0","dewptm":"12.8", "dewpti":"55.0","hum":"75","wspdm":"7.4", "wspdi":"4.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"30","wdire":"NNE","vism":"16.1", "visi":"10.0","pressurem":"1014.1", "pressurei":"29.95","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260256Z 03004KT 10SM BKN150 17/13 A2995 RMK AO2 SLP141 T01720128 51005 $" },
                {
                "date": {
                "pretty": "7:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "19",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "3:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "03",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.2", "tempi":"63.0","dewptm":"12.8", "dewpti":"55.0","hum":"75","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1014.3", "pressurei":"29.96","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260356Z 00000KT 10SM BKN150 17/13 A2995 RMK AO2 SLP143 T01720128 $" },
                {
                "date": {
                "pretty": "8:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "20",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "4:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "04",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"17.2", "tempi":"63.0","dewptm":"12.8", "dewpti":"55.0","hum":"75","wspdm":"0.0", "wspdi":"0.0","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"0","wdire":"North","vism":"16.1", "visi":"10.0","pressurem":"1014.6", "pressurei":"29.96","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260456Z 00000KT 10SM BKN150 17/13 A2996 RMK AO2 SLP146 T01720128 $" },
                {
                "date": {
                "pretty": "9:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "21",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "5:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "05",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"16.7", "tempi":"62.1","dewptm":"13.3", "dewpti":"55.9","hum":"80","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"50","wdire":"NE","vism":"16.1", "visi":"10.0","pressurem":"1014.8", "pressurei":"29.97","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260556Z 05005KT 10SM BKN150 17/13 A2997 RMK AO2 SLP148 T01670133 10189 20161 53007 $" },
                {
                "date": {
                "pretty": "10:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "22",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "6:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "06",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"20.0", "tempi":"68.0","dewptm":"10.6", "dewpti":"51.1","hum":"55","wspdm":"31.5", "wspdi":"19.6","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"200","wdire":"SSW","vism":"16.1", "visi":"10.0","pressurem":"1015.5", "pressurei":"29.99","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260656Z 20017KT 10SM SCT200 20/11 A2999 RMK AO2 SLP155 T02000106 $" },
                {
                "date": {
                "pretty": "11:56 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "23",
                "min": "56",
                "tzname": "America/Los_Angeles"
                },
                "utcdate": {
                "pretty": "7:56 AM GMT on November 26, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "26",
                "hour": "07",
                "min": "56",
                "tzname": "UTC"
                },
                "tempm":"19.4", "tempi":"66.9","dewptm":"9.4", "dewpti":"48.9","hum":"52","wspdm":"9.3", "wspdi":"5.8","wgustm":"-9999.0", "wgusti":"-9999.0","wdird":"150","wdire":"SSE","vism":"16.1", "visi":"10.0","pressurem":"1014.9", "pressurei":"29.97","windchillm":"-999", "windchilli":"-999","heatindexm":"-9999", "heatindexi":"-9999","precipm":"-9999.00", "precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR KSFO 260756Z 15005KT 10SM SCT200 19/09 A2997 RMK AO2 SLP149 T01940094 402000133 $" }
                ],
                "dailysummary": [
                { "date": {
                "pretty": "12:00 PM PST on November 25, 2017",
                "year": "2017",
                "mon": "11",
                "mday": "25",
                "hour": "12",
                "min": "00",
                "tzname": "America/Los_Angeles"
                },
                "fog":"1","rain":"0","snow":"1","snowfallm":"0.00", "snowfalli":"0.00","monthtodatesnowfallm":"25.40", "monthtodatesnowfalli":"1.00","since1julsnowfallm":"0.00", "since1julsnowfalli":"0.00","snowdepthm":"", "snowdepthi":"","hail":"0","thunder":"0","tornado":"0","meantempm":"17", "meantempi":"62","meandewptm":"12", "meandewpti":"54","meanpressurem":"1014", "meanpressurei":"29.96","meanwindspdm":"7", "meanwindspdi":"4","meanwdire":"","meanwdird":"93","meanvism":"16", "meanvisi":"10","humidity":"","maxtempm":"20", "maxtempi":"68","mintempm":"13", "mintempi":"56","maxhumidity":"86","minhumidity":"55","maxdewptm":"14", "maxdewpti":"57","mindewptm":"9", "mindewpti":"49","maxpressurem":"1015", "maxpressurei":"29.99","minpressurem":"1014", "minpressurei":"29.94","maxwspdm":"34", "maxwspdi":"21","minwspdm":"0", "minwspdi":"0","maxvism":"16", "maxvisi":"10","minvism":"16", "minvisi":"10","gdegreedays":"12","heatingdegreedays":"3","coolingdegreedays":"0","precipm":"0.00", "precipi":"0.00","precipsource":"","heatingdegreedaysnormal":"11","monthtodateheatingdegreedays":"158","monthtodateheatingdegreedaysnormal":"217","since1sepheatingdegreedays":"234","since1sepheatingdegreedaysnormal":"391","since1julheatingdegreedays":"294","since1julheatingdegreedaysnormal":"504","coolingdegreedaysnormal":"0","monthtodatecoolingdegreedays":"0","monthtodatecoolingdegreedaysnormal":"0","since1sepcoolingdegreedays":"214","since1sepcoolingdegreedaysnormal":"66","since1jancoolingdegreedays":"346","since1jancoolingdegreedaysnormal":"163" }
                ]
        }
                ,
        "almanac": {
                "airport_code": "KSFO",
                "temp_high": {
                "normal": {
                "F": "60",
                "C": "16"
                }
                ,
                "record": {
                "F": "76",
                "C": "24"
                },
                "recordyear": "1959"
                },
                "temp_low": {
                "normal": {
                "F": "47",
                "C": "8"
                }
                ,
                "record": {
                "F": "36",
                "C": "2"
                },
                "recordyear": "1952"
                }
        }
                ,
        "tide": {
                "tideInfo": [
                {
                "tideSite":"North Point, Pier 41, San Francisco, San Francisco Bay, California",
                "lat":"37.81",
                "lon":"-122.413",
                "units":"feet",
                "type":"tide",
                "tzname":"America/Los_Angeles"
                }
                ],
                "tideSummary": [
                {
                "date": {
  "pretty": "11:15 PM PST on November 26, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "26",
  "hour": "23",
  "min": "15",
  "tzname": "America/Los_Angeles",
                "epoch":"1511766957"
  },
  "utcdate": {
  "pretty": "7:15 AM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "07",
  "min": "15",
  "tzname": "UTC",
                "epoch":"1511766957"
  },
                "data": {
                "height":"0.78 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "12:14 AM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "00",
  "min": "14",
  "tzname": "America/Los_Angeles",
                "epoch":"1511770490"
  },
  "utcdate": {
  "pretty": "8:14 AM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "08",
  "min": "14",
  "tzname": "UTC",
                "epoch":"1511770490"
  },
                "data": {
                "height":"",
                "type":"Moonset"
                }
                },
                {
                "date": {
  "pretty": "6:38 AM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "06",
  "min": "38",
  "tzname": "America/Los_Angeles",
                "epoch":"1511793508"
  },
  "utcdate": {
  "pretty": "2:38 PM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "14",
  "min": "38",
  "tzname": "UTC",
                "epoch":"1511793508"
  },
                "data": {
                "height":"5.31 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "7:02 AM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "07",
  "min": "02",
  "tzname": "America/Los_Angeles",
                "epoch":"1511794969"
  },
  "utcdate": {
  "pretty": "3:02 PM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "15",
  "min": "02",
  "tzname": "UTC",
                "epoch":"1511794969"
  },
                "data": {
                "height":"",
                "type":"Sunrise"
                }
                },
                {
                "date": {
  "pretty": "12:43 PM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "12",
  "min": "43",
  "tzname": "America/Los_Angeles",
                "epoch":"1511815416"
  },
  "utcdate": {
  "pretty": "8:43 PM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "20",
  "min": "43",
  "tzname": "UTC",
                "epoch":"1511815416"
  },
                "data": {
                "height":"2.34 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "1:28 PM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "13",
  "min": "28",
  "tzname": "America/Los_Angeles",
                "epoch":"1511818139"
  },
  "utcdate": {
  "pretty": "9:28 PM GMT on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "21",
  "min": "28",
  "tzname": "UTC",
                "epoch":"1511818139"
  },
                "data": {
                "height":"",
                "type":"Moonrise"
                }
                },
                {
                "date": {
  "pretty": "4:51 PM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "16",
  "min": "51",
  "tzname": "America/Los_Angeles",
                "epoch":"1511830306"
  },
  "utcdate": {
  "pretty": "12:51 AM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "00",
  "min": "51",
  "tzname": "UTC",
                "epoch":"1511830306"
  },
                "data": {
                "height":"",
                "type":"Sunset"
                }
                },
                {
                "date": {
  "pretty": "6:12 PM PST on November 27, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "27",
  "hour": "18",
  "min": "12",
  "tzname": "America/Los_Angeles",
                "epoch":"1511835122"
  },
  "utcdate": {
  "pretty": "2:12 AM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "02",
  "min": "12",
  "tzname": "UTC",
                "epoch":"1511835122"
  },
                "data": {
                "height":"4.33 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "12:10 AM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "00",
  "min": "10",
  "tzname": "America/Los_Angeles",
                "epoch":"1511856631"
  },
  "utcdate": {
  "pretty": "8:10 AM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "08",
  "min": "10",
  "tzname": "UTC",
                "epoch":"1511856631"
  },
                "data": {
                "height":"0.98 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "1:15 AM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "01",
  "min": "15",
  "tzname": "America/Los_Angeles",
                "epoch":"1511860541"
  },
  "utcdate": {
  "pretty": "9:15 AM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "09",
  "min": "15",
  "tzname": "UTC",
                "epoch":"1511860541"
  },
                "data": {
                "height":"",
                "type":"Moonset"
                }
                },
                {
                "date": {
  "pretty": "7:03 AM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "07",
  "min": "03",
  "tzname": "America/Los_Angeles",
                "epoch":"1511881436"
  },
  "utcdate": {
  "pretty": "3:03 PM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "15",
  "min": "03",
  "tzname": "UTC",
                "epoch":"1511881436"
  },
                "data": {
                "height":"",
                "type":"Sunrise"
                }
                },
                {
                "date": {
  "pretty": "7:17 AM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "07",
  "min": "17",
  "tzname": "America/Los_Angeles",
                "epoch":"1511882240"
  },
  "utcdate": {
  "pretty": "3:17 PM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "15",
  "min": "17",
  "tzname": "UTC",
                "epoch":"1511882240"
  },
                "data": {
                "height":"5.63 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "1:32 PM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "13",
  "min": "32",
  "tzname": "America/Los_Angeles",
                "epoch":"1511904722"
  },
  "utcdate": {
  "pretty": "9:32 PM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "21",
  "min": "32",
  "tzname": "UTC",
                "epoch":"1511904722"
  },
                "data": {
                "height":"1.71 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "2:01 PM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "14",
  "min": "01",
  "tzname": "America/Los_Angeles",
                "epoch":"1511906518"
  },
  "utcdate": {
  "pretty": "10:01 PM GMT on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "22",
  "min": "01",
  "tzname": "UTC",
                "epoch":"1511906518"
  },
                "data": {
                "height":"",
                "type":"Moonrise"
                }
                },
                {
                "date": {
  "pretty": "4:51 PM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "16",
  "min": "51",
  "tzname": "America/Los_Angeles",
                "epoch":"1511916688"
  },
  "utcdate": {
  "pretty": "12:51 AM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "00",
  "min": "51",
  "tzname": "UTC",
                "epoch":"1511916688"
  },
                "data": {
                "height":"",
                "type":"Sunset"
                }
                },
                {
                "date": {
  "pretty": "7:27 PM PST on November 28, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "28",
  "hour": "19",
  "min": "27",
  "tzname": "America/Los_Angeles",
                "epoch":"1511926027"
  },
  "utcdate": {
  "pretty": "3:27 AM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "03",
  "min": "27",
  "tzname": "UTC",
                "epoch":"1511926027"
  },
                "data": {
                "height":"4.40 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "1:02 AM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "01",
  "min": "02",
  "tzname": "America/Los_Angeles",
                "epoch":"1511946134"
  },
  "utcdate": {
  "pretty": "9:02 AM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "09",
  "min": "02",
  "tzname": "UTC",
                "epoch":"1511946134"
  },
                "data": {
                "height":"1.18 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "2:18 AM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "02",
  "min": "18",
  "tzname": "America/Los_Angeles",
                "epoch":"1511950733"
  },
  "utcdate": {
  "pretty": "10:18 AM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "10",
  "min": "18",
  "tzname": "UTC",
                "epoch":"1511950733"
  },
                "data": {
                "height":"",
                "type":"Moonset"
                }
                },
                {
                "date": {
  "pretty": "7:04 AM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "07",
  "min": "04",
  "tzname": "America/Los_Angeles",
                "epoch":"1511967895"
  },
  "utcdate": {
  "pretty": "3:04 PM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "15",
  "min": "04",
  "tzname": "UTC",
                "epoch":"1511967895"
  },
                "data": {
                "height":"",
                "type":"Sunrise"
                }
                },
                {
                "date": {
  "pretty": "7:54 AM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "07",
  "min": "54",
  "tzname": "America/Los_Angeles",
                "epoch":"1511970896"
  },
  "utcdate": {
  "pretty": "3:54 PM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "15",
  "min": "54",
  "tzname": "UTC",
                "epoch":"1511970896"
  },
                "data": {
                "height":"5.98 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "2:15 PM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "14",
  "min": "15",
  "tzname": "America/Los_Angeles",
                "epoch":"1511993754"
  },
  "utcdate": {
  "pretty": "10:15 PM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "22",
  "min": "15",
  "tzname": "UTC",
                "epoch":"1511993754"
  },
                "data": {
                "height":"1.00 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "2:35 PM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "14",
  "min": "35",
  "tzname": "America/Los_Angeles",
                "epoch":"1511994931"
  },
  "utcdate": {
  "pretty": "10:35 PM GMT on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "22",
  "min": "35",
  "tzname": "UTC",
                "epoch":"1511994931"
  },
                "data": {
                "height":"",
                "type":"Moonrise"
                }
                },
                {
                "date": {
  "pretty": "4:51 PM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "16",
  "min": "51",
  "tzname": "America/Los_Angeles",
                "epoch":"1512003072"
  },
  "utcdate": {
  "pretty": "12:51 AM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "00",
  "min": "51",
  "tzname": "UTC",
                "epoch":"1512003072"
  },
                "data": {
                "height":"",
                "type":"Sunset"
                }
                },
                {
                "date": {
  "pretty": "8:35 PM PST on November 29, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "29",
  "hour": "20",
  "min": "35",
  "tzname": "America/Los_Angeles",
                "epoch":"1512016536"
  },
  "utcdate": {
  "pretty": "4:35 AM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "04",
  "min": "35",
  "tzname": "UTC",
                "epoch":"1512016536"
  },
                "data": {
                "height":"4.58 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "1:51 AM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "01",
  "min": "51",
  "tzname": "America/Los_Angeles",
                "epoch":"1512035491"
  },
  "utcdate": {
  "pretty": "9:51 AM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "09",
  "min": "51",
  "tzname": "UTC",
                "epoch":"1512035491"
  },
                "data": {
                "height":"1.40 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "3:24 AM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "03",
  "min": "24",
  "tzname": "America/Los_Angeles",
                "epoch":"1512041081"
  },
  "utcdate": {
  "pretty": "11:24 AM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "11",
  "min": "24",
  "tzname": "UTC",
                "epoch":"1512041081"
  },
                "data": {
                "height":"",
                "type":"Moonset"
                }
                },
                {
                "date": {
  "pretty": "7:05 AM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "07",
  "min": "05",
  "tzname": "America/Los_Angeles",
                "epoch":"1512054354"
  },
  "utcdate": {
  "pretty": "3:05 PM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "15",
  "min": "05",
  "tzname": "UTC",
                "epoch":"1512054354"
  },
                "data": {
                "height":"",
                "type":"Sunrise"
                }
                },
                {
                "date": {
  "pretty": "8:32 AM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "08",
  "min": "32",
  "tzname": "America/Los_Angeles",
                "epoch":"1512059568"
  },
  "utcdate": {
  "pretty": "4:32 PM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "16",
  "min": "32",
  "tzname": "UTC",
                "epoch":"1512059568"
  },
                "data": {
                "height":"6.34 ft",
                "type":"High Tide"
                }
                },
                {
                "date": {
  "pretty": "2:58 PM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "14",
  "min": "58",
  "tzname": "America/Los_Angeles",
                "epoch":"1512082691"
  },
  "utcdate": {
  "pretty": "10:58 PM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "22",
  "min": "58",
  "tzname": "UTC",
                "epoch":"1512082691"
  },
                "data": {
                "height":"0.27 ft",
                "type":"Low Tide"
                }
                },
                {
                "date": {
  "pretty": "3:11 PM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "15",
  "min": "11",
  "tzname": "America/Los_Angeles",
                "epoch":"1512083479"
  },
  "utcdate": {
  "pretty": "11:11 PM GMT on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "23",
  "min": "11",
  "tzname": "UTC",
                "epoch":"1512083479"
  },
                "data": {
                "height":"",
                "type":"Moonrise"
                }
                },
                {
                "date": {
  "pretty": "4:50 PM PST on November 30, 2017",
  "year": "2017",
  "mon": "11",
  "mday": "30",
  "hour": "16",
  "min": "50",
  "tzname": "America/Los_Angeles",
                "epoch":"1512089459"
  },
  "utcdate": {
  "pretty": "12:50 AM GMT on December 01, 2017",
  "year": "2017",
  "mon": "12",
  "mday": "01",
  "hour": "00",
  "min": "50",
  "tzname": "UTC",
                "epoch":"1512089459"
  },
                "data": {
                "height":"",
                "type":"Sunset"
                }
                }
                ],
                "tideSummaryStats": [
                {
  "maxheight":6.340000,
  "minheight":0.000000
                }
  ]
        }
                ,
        "rawtide": {
                "tideInfo": [
                {
                "tideSite":"North Point, Pier 41, San Francisco, San Francisco Bay, California",
                "lat":"37.81",
                "lon":"-122.413",
                "units":"feet",
                "type":"tide",
                "tzname":"America/Los_Angeles"
                }
                ],
                "rawTideObs": [
                {
                "epoch":1511753086,
                "height":3.261059
                },
                {
                "epoch":1511753386,
                "height":3.188894
                },
                {
                "epoch":1511753686,
                "height":3.115586
                },
                {
                "epoch":1511753986,
                "height":3.041269
                },
                {
                "epoch":1511754286,
                "height":2.966082
                },
                {
                "epoch":1511754586,
                "height":2.890166
                },
                {
                "epoch":1511754886,
                "height":2.813664
                },
                {
                "epoch":1511755186,
                "height":2.736719
                },
                {
                "epoch":1511755486,
                "height":2.659478
                },
                {
                "epoch":1511755786,
                "height":2.582087
                },
                {
                "epoch":1511756086,
                "height":2.504696
                },
                {
                "epoch":1511756386,
                "height":2.427450
                },
                {
                "epoch":1511756686,
                "height":2.350497
                },
                {
                "epoch":1511756986,
                "height":2.273985
                },
                {
                "epoch":1511757286,
                "height":2.198059
                },
                {
                "epoch":1511757586,
                "height":2.122863
                },
                {
                "epoch":1511757886,
                "height":2.048541
                },
                {
                "epoch":1511758186,
                "height":1.975232
                },
                {
                "epoch":1511758486,
                "height":1.903074
                },
                {
                "epoch":1511758786,
                "height":1.832202
                },
                {
                "epoch":1511759086,
                "height":1.762749
                },
                {
                "epoch":1511759386,
                "height":1.694841
                },
                {
                "epoch":1511759686,
                "height":1.628605
                },
                {
                "epoch":1511759986,
                "height":1.564159
                },
                {
                "epoch":1511760286,
                "height":1.501620
                },
                {
                "epoch":1511760586,
                "height":1.441099
                },
                {
                "epoch":1511760886,
                "height":1.382702
                },
                {
                "epoch":1511761186,
                "height":1.326530
                },
                {
                "epoch":1511761486,
                "height":1.272681
                },
                {
                "epoch":1511761786,
                "height":1.221243
                },
                {
                "epoch":1511762086,
                "height":1.172301
                },
                {
                "epoch":1511762386,
                "height":1.125936
                },
                {
                "epoch":1511762686,
                "height":1.082220
                },
                {
                "epoch":1511762986,
                "height":1.041221
                },
                {
                "epoch":1511763286,
                "height":1.003001
                },
                {
                "epoch":1511763586,
                "height":0.967614
                },
                {
                "epoch":1511763886,
                "height":0.935111
                },
                {
                "epoch":1511764186,
                "height":0.905534
                },
                {
                "epoch":1511764486,
                "height":0.878922
                },
                {
                "epoch":1511764786,
                "height":0.855306
                },
                {
                "epoch":1511765086,
                "height":0.834711
                },
                {
                "epoch":1511765386,
                "height":0.817158
                },
                {
                "epoch":1511765686,
                "height":0.802660
                },
                {
                "epoch":1511765986,
                "height":0.791227
                },
                {
                "epoch":1511766286,
                "height":0.782861
                },
                {
                "epoch":1511766586,
                "height":0.777560
                },
                {
                "epoch":1511766886,
                "height":0.775317
                },
                {
                "epoch":1511767186,
                "height":0.776120
                },
                {
                "epoch":1511767486,
                "height":0.779951
                },
                {
                "epoch":1511767786,
                "height":0.786789
                },
                {
                "epoch":1511768086,
                "height":0.796607
                },
                {
                "epoch":1511768386,
                "height":0.809375
                },
                {
                "epoch":1511768686,
                "height":0.825058
                },
                {
                "epoch":1511768986,
                "height":0.843618
                },
                {
                "epoch":1511769286,
                "height":0.865013
                },
                {
                "epoch":1511769586,
                "height":0.889197
                },
                {
                "epoch":1511769886,
                "height":0.916121
                },
                {
                "epoch":1511770186,
                "height":0.945734
                },
                {
                "epoch":1511770486,
                "height":0.977982
                },
                {
                "epoch":1511770786,
                "height":1.012806
                },
                {
                "epoch":1511771086,
                "height":1.050148
                },
                {
                "epoch":1511771386,
                "height":1.089946
                },
                {
                "epoch":1511771686,
                "height":1.132135
                },
                {
                "epoch":1511771986,
                "height":1.176650
                },
                {
                "epoch":1511772286,
                "height":1.223423
                },
                {
                "epoch":1511772586,
                "height":1.272386
                },
                {
                "epoch":1511772886,
                "height":1.323467
                },
                {
                "epoch":1511773186,
                "height":1.376595
                },
                {
                "epoch":1511773486,
                "height":1.431697
                },
                {
                "epoch":1511773786,
                "height":1.488698
                },
                {
                "epoch":1511774086,
                "height":1.547525
                },
                {
                "epoch":1511774386,
                "height":1.608101
                },
                {
                "epoch":1511774686,
                "height":1.670349
                },
                {
                "epoch":1511774986,
                "height":1.734192
                },
                {
                "epoch":1511775286,
                "height":1.799552
                },
                {
                "epoch":1511775586,
                "height":1.866352
                },
                {
                "epoch":1511775886,
                "height":1.934511
                },
                {
                "epoch":1511776186,
                "height":2.003950
                },
                {
                "epoch":1511776486,
                "height":2.074590
                },
                {
                "epoch":1511776786,
                "height":2.146349
                },
                {
                "epoch":1511777086,
                "height":2.219147
                },
                {
                "epoch":1511777386,
                "height":2.292903
                },
                {
                "epoch":1511777686,
                "height":2.367534
                },
                {
                "epoch":1511777986,
                "height":2.442958
                },
                {
                "epoch":1511778286,
                "height":2.519092
                },
                {
                "epoch":1511778586,
                "height":2.595854
                },
                {
                "epoch":1511778886,
                "height":2.673159
                },
                {
                "epoch":1511779186,
                "height":2.750922
                },
                {
                "epoch":1511779486,
                "height":2.829060
                },
                {
                "epoch":1511779786,
                "height":2.907486
                },
                {
                "epoch":1511780086,
                "height":2.986115
                },
                {
                "epoch":1511780386,
                "height":3.064859
                },
                {
                "epoch":1511780686,
                "height":3.143632
                },
                {
                "epoch":1511780986,
                "height":3.222346
                },
                {
                "epoch":1511781286,
                "height":3.300912
                },
                {
                "epoch":1511781586,
                "height":3.379241
                },
                {
                "epoch":1511781886,
                "height":3.457243
                },
                {
                "epoch":1511782186,
                "height":3.534829
                },
                {
                "epoch":1511782486,
                "height":3.611908
                },
                {
                "epoch":1511782786,
                "height":3.688388
                },
                {
                "epoch":1511783086,
                "height":3.764178
                },
                {
                "epoch":1511783386,
                "height":3.839186
                },
                {
                "epoch":1511783686,
                "height":3.913319
                },
                {
                "epoch":1511783986,
                "height":3.986486
                },
                {
                "epoch":1511784286,
                "height":4.058594
                },
                {
                "epoch":1511784586,
                "height":4.129551
                },
                {
                "epoch":1511784886,
                "height":4.199263
                },
                {
                "epoch":1511785186,
                "height":4.267639
                },
                {
                "epoch":1511785486,
                "height":4.334588
                },
                {
                "epoch":1511785786,
                "height":4.400019
                },
                {
                "epoch":1511786086,
                "height":4.463840
                },
                {
                "epoch":1511786386,
                "height":4.525963
                },
                {
                "epoch":1511786686,
                "height":4.586299
                },
                {
                "epoch":1511786986,
                "height":4.644762
                },
                {
                "epoch":1511787286,
                "height":4.701266
                },
                {
                "epoch":1511787586,
                "height":4.755726
                },
                {
                "epoch":1511787886,
                "height":4.808063
                },
                {
                "epoch":1511788186,
                "height":4.858195
                },
                {
                "epoch":1511788486,
                "height":4.906045
                },
                {
                "epoch":1511788786,
                "height":4.951539
                },
                {
                "epoch":1511789086,
                "height":4.994604
                },
                {
                "epoch":1511789386,
                "height":5.035173
                },
                {
                "epoch":1511789686,
                "height":5.073180
                },
                {
                "epoch":1511789986,
                "height":5.108562
                },
                {
                "epoch":1511790286,
                "height":5.141261
                },
                {
                "epoch":1511790586,
                "height":5.171224
                },
                {
                "epoch":1511790886,
                "height":5.198400
                },
                {
                "epoch":1511791186,
                "height":5.222744
                },
                {
                "epoch":1511791486,
                "height":5.244215
                },
                {
                "epoch":1511791786,
                "height":5.262777
                },
                {
                "epoch":1511792086,
                "height":5.278399
                },
                {
                "epoch":1511792386,
                "height":5.291055
                },
                {
                "epoch":1511792686,
                "height":5.300724
                },
                {
                "epoch":1511792986,
                "height":5.307391
                },
                {
                "epoch":1511793286,
                "height":5.311047
                },
                {
                "epoch":1511793586,
                "height":5.311689
                },
                {
                "epoch":1511793886,
                "height":5.309318
                },
                {
                "epoch":1511794186,
                "height":5.303942
                },
                {
                "epoch":1511794486,
                "height":5.295576
                },
                {
                "epoch":1511794786,
                "height":5.284239
                },
                {
                "epoch":1511795086,
                "height":5.269959
                },
                {
                "epoch":1511795386,
                "height":5.252766
                },
                {
                "epoch":1511795686,
                "height":5.232699
                },
                {
                "epoch":1511795986,
                "height":5.209802
                },
                {
                "epoch":1511796286,
                "height":5.184126
                },
                {
                "epoch":1511796586,
                "height":5.155727
                },
                {
                "epoch":1511796886,
                "height":5.124666
                },
                {
                "epoch":1511797186,
                "height":5.091011
                },
                {
                "epoch":1511797486,
                "height":5.054835
                },
                {
                "epoch":1511797786,
                "height":5.016216
                },
                {
                "epoch":1511798086,
                "height":4.975238
                },
                {
                "epoch":1511798386,
                "height":4.931990
                },
                {
                "epoch":1511798686,
                "height":4.886564
                },
                {
                "epoch":1511798986,
                "height":4.839059
                },
                {
                "epoch":1511799286,
                "height":4.789576
                },
                {
                "epoch":1511799586,
                "height":4.738223
                },
                {
                "epoch":1511799886,
                "height":4.685108
                },
                {
                "epoch":1511800186,
                "height":4.630347
                },
                {
                "epoch":1511800486,
                "height":4.574054
                },
                {
                "epoch":1511800786,
                "height":4.516350
                },
                {
                "epoch":1511801086,
                "height":4.457357
                },
                {
                "epoch":1511801386,
                "height":4.397199
                },
                {
                "epoch":1511801686,
                "height":4.336003
                },
                {
                "epoch":1511801986,
                "height":4.273895
                },
                {
                "epoch":1511802286,
                "height":4.211006
                },
                {
                "epoch":1511802586,
                "height":4.147466
                },
                {
                "epoch":1511802886,
                "height":4.083403
                },
                {
                "epoch":1511803186,
                "height":4.018950
                },
                {
                "epoch":1511803486,
                "height":3.954237
                },
                {
                "epoch":1511803786,
                "height":3.889394
                },
                {
                "epoch":1511804086,
                "height":3.824549
                },
                {
                "epoch":1511804386,
                "height":3.759832
                },
                {
                "epoch":1511804686,
                "height":3.695369
                },
                {
                "epoch":1511804986,
                "height":3.631284
                },
                {
                "epoch":1511805286,
                "height":3.567700
                },
                {
                "epoch":1511805586,
                "height":3.504739
                },
                {
                "epoch":1511805886,
                "height":3.442518
                },
                {
                "epoch":1511806186,
                "height":3.381151
                },
                {
                "epoch":1511806486,
                "height":3.320752
                },
                {
                "epoch":1511806786,
                "height":3.261427
                },
                {
                "epoch":1511807086,
                "height":3.203283
                },
                {
                "epoch":1511807386,
                "height":3.146421
                },
                {
                "epoch":1511807686,
                "height":3.090938
                },
                {
                "epoch":1511807986,
                "height":3.036926
                },
                {
                "epoch":1511808286,
                "height":2.984475
                },
                {
                "epoch":1511808586,
                "height":2.933670
                },
                {
                "epoch":1511808886,
                "height":2.884589
                },
                {
                "epoch":1511809186,
                "height":2.837310
                },
                {
                "epoch":1511809486,
                "height":2.791901
                },
                {
                "epoch":1511809786,
                "height":2.748429
                },
                {
                "epoch":1511810086,
                "height":2.706955
                },
                {
                "epoch":1511810386,
                "height":2.667535
                },
                {
                "epoch":1511810686,
                "height":2.630221
                },
                {
                "epoch":1511810986,
                "height":2.595058
                },
                {
                "epoch":1511811286,
                "height":2.562088
                },
                {
                "epoch":1511811586,
                "height":2.531348
                },
                {
                "epoch":1511811886,
                "height":2.502870
                },
                {
                "epoch":1511812186,
                "height":2.476680
                },
                {
                "epoch":1511812486,
                "height":2.452802
                },
                {
                "epoch":1511812786,
                "height":2.431253
                },
                {
                "epoch":1511813086,
                "height":2.412047
                },
                {
                "epoch":1511813386,
                "height":2.395192
                },
                {
                "epoch":1511813686,
                "height":2.380692
                },
                {
                "epoch":1511813986,
                "height":2.368549
                },
                {
                "epoch":1511814286,
                "height":2.358757
                },
                {
                "epoch":1511814586,
                "height":2.351310
                },
                {
                "epoch":1511814886,
                "height":2.346195
                },
                {
                "epoch":1511815186,
                "height":2.343397
                },
                {
                "epoch":1511815486,
                "height":2.342896
                },
                {
                "epoch":1511815786,
                "height":2.344670
                },
                {
                "epoch":1511816086,
                "height":2.348692
                },
                {
                "epoch":1511816386,
                "height":2.354932
                },
                {
                "epoch":1511816686,
                "height":2.363359
                },
                {
                "epoch":1511816986,
                "height":2.373935
                },
                {
                "epoch":1511817286,
                "height":2.386623
                },
                {
                "epoch":1511817586,
                "height":2.401380
                },
                {
                "epoch":1511817886,
                "height":2.418162
                },
                {
                "epoch":1511818186,
                "height":2.436922
                },
                {
                "epoch":1511818486,
                "height":2.457611
                },
                {
                "epoch":1511818786,
                "height":2.480175
                },
                {
                "epoch":1511819086,
                "height":2.504561
                },
                {
                "epoch":1511819386,
                "height":2.530711
                },
                {
                "epoch":1511819686,
                "height":2.558568
                },
                {
                "epoch":1511819986,
                "height":2.588068
                },
                {
                "epoch":1511820286,
                "height":2.619149
                },
                {
                "epoch":1511820586,
                "height":2.651746
                },
                {
                "epoch":1511820886,
                "height":2.685791
                },
                {
                "epoch":1511821186,
                "height":2.721214
                },
                {
                "epoch":1511821486,
                "height":2.757944
                },
                {
                "epoch":1511821786,
                "height":2.795908
                },
                {
                "epoch":1511822086,
                "height":2.835031
                },
                {
                "epoch":1511822386,
                "height":2.875236
                },
                {
                "epoch":1511822686,
                "height":2.916444
                },
                {
                "epoch":1511822986,
                "height":2.958576
                },
                {
                "epoch":1511823286,
                "height":3.001548
                },
                {
                "epoch":1511823586,
                "height":3.045277
                },
                {
                "epoch":1511823886,
                "height":3.089678
                },
                {
                "epoch":1511824186,
                "height":3.134664
                },
                {
                "epoch":1511824486,
                "height":3.180146
                },
                {
                "epoch":1511824786,
                "height":3.226034
                },
                {
                "epoch":1511825086,
                "height":3.272237
                },
                {
                "epoch":1511825386,
                "height":3.318662
                },
                {
                "epoch":1511825686,
                "height":3.365214
                },
                {
                "epoch":1511825986,
                "height":3.411799
                },
                {
                "epoch":1511826286,
                "height":3.458319
                },
                {
                "epoch":1511826586,
                "height":3.504677
                },
                {
                "epoch":1511826886,
                "height":3.550774
                },
                {
                "epoch":1511827186,
                "height":3.596510
                },
                {
                "epoch":1511827486,
                "height":3.641785
                },
                {
                "epoch":1511827786,
                "height":3.686497
                },
                {
                "epoch":1511828086,
                "height":3.730546
                },
                {
                "epoch":1511828386,
                "height":3.773830
                },
                {
                "epoch":1511828686,
                "height":3.816245
                },
                {
                "epoch":1511828986,
                "height":3.857691
                },
                {
                "epoch":1511829286,
                "height":3.898064
                },
                {
                "epoch":1511829586,
                "height":3.937264
                },
                {
                "epoch":1511829886,
                "height":3.975190
                },
                {
                "epoch":1511830186,
                "height":4.011740
                },
                {
                "epoch":1511830486,
                "height":4.046816
                },
                {
                "epoch":1511830786,
                "height":4.080320
                },
                {
                "epoch":1511831086,
                "height":4.112155
                },
                {
                "epoch":1511831386,
                "height":4.142226
                },
                {
                "epoch":1511831686,
                "height":4.170440
                },
                {
                "epoch":1511831986,
                "height":4.196708
                },
                {
                "epoch":1511832286,
                "height":4.220940
                },
                {
                "epoch":1511832586,
                "height":4.243052
                },
                {
                "epoch":1511832886,
                "height":4.262961
                },
                {
                "epoch":1511833186,
                "height":4.280589
                },
                {
                "epoch":1511833486,
                "height":4.295860
                },
                {
                "epoch":1511833786,
                "height":4.308703
                },
                {
                "epoch":1511834086,
                "height":4.319052
                },
                {
                "epoch":1511834386,
                "height":4.326844
                },
                {
                "epoch":1511834686,
                "height":4.332021
                },
                {
                "epoch":1511834986,
                "height":4.334530
                },
                {
                "epoch":1511835286,
                "height":4.334324
                },
                {
                "epoch":1511835586,
                "height":4.331361
                },
                {
                "epoch":1511835886,
                "height":4.325605
                },
                {
                "epoch":1511836186,
                "height":4.317025
                },
                {
                "epoch":1511836486,
                "height":4.305598
                },
                {
                "epoch":1511836786,
                "height":4.291306
                },
                {
                "epoch":1511837086,
                "height":4.274137
                },
                {
                "epoch":1511837386,
                "height":4.254087
                },
                {
                "epoch":1511837686,
                "height":4.231159
                },
                {
                "epoch":1511837986,
                "height":4.205362
                },
                {
                "epoch":1511838286,
                "height":4.176712
                },
                {
                "epoch":1511838586,
                "height":4.145234
                },
                {
                "epoch":1511838886,
                "height":4.110957
                },
                {
                "epoch":1511839186,
                "height":4.073920
                },
                {
                "epoch":1511839486,
                "height":4.034169
                },
                {
                "epoch":1511839786,
                "height":3.991754
                },
                {
                "epoch":1511840086,
                "height":3.946737
                },
                {
                "epoch":1511840386,
                "height":3.899183
                },
                {
                "epoch":1511840686,
                "height":3.849167
                },
                {
                "epoch":1511840986,
                "height":3.796769
                },
                {
                "epoch":1511841286,
                "height":3.742076
                },
                {
                "epoch":1511841586,
                "height":3.685181
                },
                {
                "epoch":1511841886,
                "height":3.626186
                },
                {
                "epoch":1511842186,
                "height":3.565196
                },
                {
                "epoch":1511842486,
                "height":3.502323
                },
                {
                "epoch":1511842786,
                "height":3.437686
                },
                {
                "epoch":1511843086,
                "height":3.371406
                },
                {
                "epoch":1511843386,
                "height":3.303612
                },
                {
                "epoch":1511843686,
                "height":3.234437
                },
                {
                "epoch":1511843986,
                "height":3.164018
                },
                {
                "epoch":1511844286,
                "height":3.092496
                },
                {
                "epoch":1511844586,
                "height":3.020014
                },
                {
                "epoch":1511844886,
                "height":2.946721
                },
                {
                "epoch":1511845186,
                "height":2.872768
                },
                {
                "epoch":1511845486,
                "height":2.798306
                },
                {
                "epoch":1511845786,
                "height":2.723490
                },
                {
                "epoch":1511846086,
                "height":2.648477
                },
                {
                "epoch":1511846386,
                "height":2.573423
                },
                {
                "epoch":1511846686,
                "height":2.498486
                },
                {
                "epoch":1511846986,
                "height":2.423824
                },
                {
                "epoch":1511847286,
                "height":2.349594
                },
                {
                "epoch":1511847586,
                "height":2.275952
                },
                {
                "epoch":1511847886,
                "height":2.203055
                },
                {
                "epoch":1511848186,
                "height":2.131055
                },
                {
                "epoch":1511848486,
                "height":2.060104
                },
                {
                "epoch":1511848786,
                "height":1.990352
                },
                {
                "epoch":1511849086,
                "height":1.921944
                },
                {
                "epoch":1511849386,
                "height":1.855024
                },
                {
                "epoch":1511849686,
                "height":1.789730
                },
                {
                "epoch":1511849986,
                "height":1.726198
                },
                {
                "epoch":1511850286,
                "height":1.664558
                },
                {
                "epoch":1511850586,
                "height":1.604936
                },
                {
                "epoch":1511850886,
                "height":1.547454
                },
                {
                "epoch":1511851186,
                "height":1.492227
                },
                {
                "epoch":1511851486,
                "height":1.439366
                },
                {
                "epoch":1511851786,
                "height":1.388973
                },
                {
                "epoch":1511852086,
                "height":1.341149
                },
                {
                "epoch":1511852386,
                "height":1.295985
                },
                {
                "epoch":1511852686,
                "height":1.253566
                },
                {
                "epoch":1511852986,
                "height":1.213972
                },
                {
                "epoch":1511853286,
                "height":1.177276
                },
                {
                "epoch":1511853586,
                "height":1.143544
                },
                {
                "epoch":1511853886,
                "height":1.112834
                },
                {
                "epoch":1511854186,
                "height":1.085200
                },
                {
                "epoch":1511854486,
                "height":1.060687
                },
                {
                "epoch":1511854786,
                "height":1.039333
                },
                {
                "epoch":1511855086,
                "height":1.021171
                },
                {
                "epoch":1511855386,
                "height":1.006227
                },
                {
                "epoch":1511855686,
                "height":0.994518
                },
                {
                "epoch":1511855986,
                "height":0.986058
                },
                {
                "epoch":1511856286,
                "height":0.980851
                },
                {
                "epoch":1511856586,
                "height":0.978899
                },
                {
                "epoch":1511856886,
                "height":0.980195
                },
                {
                "epoch":1511857186,
                "height":0.984726
                },
                {
                "epoch":1511857486,
                "height":0.992474
                },
                {
                "epoch":1511857786,
                "height":1.003416
                },
                {
                "epoch":1511858086,
                "height":1.017522
                },
                {
                "epoch":1511858386,
                "height":1.034761
                },
                {
                "epoch":1511858686,
                "height":1.055091
                },
                {
                "epoch":1511858986,
                "height":1.078471
                },
                {
                "epoch":1511859286,
                "height":1.104852
                },
                {
                "epoch":1511859586,
                "height":1.134184
                },
                {
                "epoch":1511859886,
                "height":1.166409
                },
                {
                "epoch":1511860186,
                "height":1.201469
                },
                {
                "epoch":1511860486,
                "height":1.239302
                },
                {
                "epoch":1511860786,
                "height":1.279841
                },
                {
                "epoch":1511861086,
                "height":1.323018
                },
                {
                "epoch":1511861386,
                "height":1.368762
                },
                {
                "epoch":1511861686,
                "height":1.416999
                },
                {
                "epoch":1511861986,
                "height":1.467653
                },
                {
                "epoch":1511862286,
                "height":1.520646
                },
                {
                "epoch":1511862586,
                "height":1.575898
                },
                {
                "epoch":1511862886,
                "height":1.633329
                },
                {
                "epoch":1511863186,
                "height":1.692855
                },
                {
                "epoch":1511863486,
                "height":1.754392
                },
                {
                "epoch":1511863786,
                "height":1.817856
                },
                {
                "epoch":1511864086,
                "height":1.883159
                },
                {
                "epoch":1511864386,
                "height":1.950215
                },
                {
                "epoch":1511864686,
                "height":2.018937
                },
                {
                "epoch":1511864986,
                "height":2.089235
                },
                {
                "epoch":1511865286,
                "height":2.161021
                },
                {
                "epoch":1511865586,
                "height":2.234205
                },
                {
                "epoch":1511865886,
                "height":2.308698
                },
                {
                "epoch":1511866186,
                "height":2.384408
                },
                {
                "epoch":1511866486,
                "height":2.461244
                },
                {
                "epoch":1511866786,
                "height":2.539117
                },
                {
                "epoch":1511867086,
                "height":2.617934
                },
                {
                "epoch":1511867386,
                "height":2.697604
                },
                {
                "epoch":1511867686,
                "height":2.778033
                },
                {
                "epoch":1511867986,
                "height":2.859130
                },
                {
                "epoch":1511868286,
                "height":2.940802
                },
                {
                "epoch":1511868586,
                "height":3.022955
                },
                {
                "epoch":1511868886,
                "height":3.105495
                },
                {
                "epoch":1511869186,
                "height":3.188328
                },
                {
                "epoch":1511869486,
                "height":3.271358
                },
                {
                "epoch":1511869786,
                "height":3.354491
                },
                {
                "epoch":1511870086,
                "height":3.437630
                },
                {
                "epoch":1511870386,
                "height":3.520678
                },
                {
                "epoch":1511870686,
                "height":3.603539
                },
                {
                "epoch":1511870986,
                "height":3.686113
                },
                {
                "epoch":1511871286,
                "height":3.768303
                },
                {
                "epoch":1511871586,
                "height":3.850008
                },
                {
                "epoch":1511871886,
                "height":3.931129
                },
                {
                "epoch":1511872186,
                "height":4.011565
                },
                {
                "epoch":1511872486,
                "height":4.091214
                },
                {
                "epoch":1511872786,
                "height":4.169974
                },
                {
                "epoch":1511873086,
                "height":4.247742
                },
                {
                "epoch":1511873386,
                "height":4.324415
                },
                {
                "epoch":1511873686,
                "height":4.399890
                },
                {
                "epoch":1511873986,
                "height":4.474063
                },
                {
                "epoch":1511874286,
                "height":4.546829
                },
                {
                "epoch":1511874586,
                "height":4.618084
                },
                {
                "epoch":1511874886,
                "height":4.687724
                },
                {
                "epoch":1511875186,
                "height":4.755644
                },
                {
                "epoch":1511875486,
                "height":4.821742
                },
                {
                "epoch":1511875786,
                "height":4.885914
                },
                {
                "epoch":1511876086,
                "height":4.948057
                },
                {
                "epoch":1511876386,
                "height":5.008071
                },
                {
                "epoch":1511876686,
                "height":5.065855
                },
                {
                "epoch":1511876986,
                "height":5.121310
                },
                {
                "epoch":1511877286,
                "height":5.174339
                },
                {
                "epoch":1511877586,
                "height":5.224848
                },
                {
                "epoch":1511877886,
                "height":5.272742
                },
                {
                "epoch":1511878186,
                "height":5.317934
                },
                {
                "epoch":1511878486,
                "height":5.360333
                },
                {
                "epoch":1511878786,
                "height":5.399857
                },
                {
                "epoch":1511879086,
                "height":5.436425
                },
                {
                "epoch":1511879386,
                "height":5.469959
                },
                {
                "epoch":1511879686,
                "height":5.500386
                },
                {
                "epoch":1511879986,
                "height":5.527637
                },
                {
                "epoch":1511880286,
                "height":5.551648
                },
                {
                "epoch":1511880586,
                "height":5.572360
                },
                {
                "epoch":1511880886,
                "height":5.589719
                },
                {
                "epoch":1511881186,
                "height":5.603677
                },
                {
                "epoch":1511881486,
                "height":5.614190
                },
                {
                "epoch":1511881786,
                "height":5.621221
                },
                {
                "epoch":1511882086,
                "height":5.624742
                },
                {
                "epoch":1511882386,
                "height":5.624727
                },
                {
                "epoch":1511882686,
                "height":5.621159
                },
                {
                "epoch":1511882986,
                "height":5.614029
                },
                {
                "epoch":1511883286,
                "height":5.603333
                },
                {
                "epoch":1511883586,
                "height":5.589075
                },
                {
                "epoch":1511883886,
                "height":5.571268
                },
                {
                "epoch":1511884186,
                "height":5.549929
                },
                {
                "epoch":1511884486,
                "height":5.525086
                },
                {
                "epoch":1511884786,
                "height":5.496774
                },
                {
                "epoch":1511885086,
                "height":5.465033
                },
                {
                "epoch":1511885386,
                "height":5.429914
                },
                {
                "epoch":1511885686,
                "height":5.391474
                },
                {
                "epoch":1511885986,
                "height":5.349778
                },
                {
                "epoch":1511886286,
                "height":5.304899
                },
                {
                "epoch":1511886586,
                "height":5.256916
                },
                {
                "epoch":1511886886,
                "height":5.205916
                },
                {
                "epoch":1511887186,
                "height":5.151995
                },
                {
                "epoch":1511887486,
                "height":5.095252
                },
                {
                "epoch":1511887786,
                "height":5.035797
                },
                {
                "epoch":1511888086,
                "height":4.973742
                },
                {
                "epoch":1511888386,
                "height":4.909210
                },
                {
                "epoch":1511888686,
                "height":4.842325
                },
                {
                "epoch":1511888986,
                "height":4.773220
                },
                {
                "epoch":1511889286,
                "height":4.702032
                },
                {
                "epoch":1511889586,
                "height":4.628903
                },
                {
                "epoch":1511889886,
                "height":4.553979
                },
                {
                "epoch":1511890186,
                "height":4.477411
                },
                {
                "epoch":1511890486,
                "height":4.399351
                },
                {
                "epoch":1511890786,
                "height":4.319959
                },
                {
                "epoch":1511891086,
                "height":4.239392
                },
                {
                "epoch":1511891386,
                "height":4.157815
                },
                {
                "epoch":1511891686,
                "height":4.075389
                },
                {
                "epoch":1511891986,
                "height":3.992281
                },
                {
                "epoch":1511892286,
                "height":3.908657
                },
                {
                "epoch":1511892586,
                "height":3.824684
                },
                {
                "epoch":1511892886,
                "height":3.740527
                },
                {
                "epoch":1511893186,
                "height":3.656352
                },
                {
                "epoch":1511893486,
                "height":3.572326
                },
                {
                "epoch":1511893786,
                "height":3.488609
                },
                {
                "epoch":1511894086,
                "height":3.405365
                },
                {
                "epoch":1511894386,
                "height":3.322753
                },
                {
                "epoch":1511894686,
                "height":3.240927
                },
                {
                "epoch":1511894986,
                "height":3.160042
                },
                {
                "epoch":1511895286,
                "height":3.080247
                },
                {
                "epoch":1511895586,
                "height":3.001687
                },
                {
                "epoch":1511895886,
                "height":2.924503
                },
                {
                "epoch":1511896186,
                "height":2.848832
                },
                {
                "epoch":1511896486,
                "height":2.774805
                },
                {
                "epoch":1511896786,
                "height":2.702547
                },
                {
                "epoch":1511897086,
                "height":2.632180
                },
                {
                "epoch":1511897386,
                "height":2.563818
                },
                {
                "epoch":1511897686,
                "height":2.497570
                },
                {
                "epoch":1511897986,
                "height":2.433539
                },
                {
                "epoch":1511898286,
                "height":2.371819
                },
                {
                "epoch":1511898586,
                "height":2.312502
                },
                {
                "epoch":1511898886,
                "height":2.255670
                },
                {
                "epoch":1511899186,
                "height":2.201399
                },
                {
                "epoch":1511899486,
                "height":2.149759
                },
                {
                "epoch":1511899786,
                "height":2.100814
                },
                {
                "epoch":1511900086,
                "height":2.054617
                },
                {
                "epoch":1511900386,
                "height":2.011221
                },
                {
                "epoch":1511900686,
                "height":1.970666
                },
                {
                "epoch":1511900986,
                "height":1.932989
                },
                {
                "epoch":1511901286,
                "height":1.898219
                },
                {
                "epoch":1511901586,
                "height":1.866380
                },
                {
                "epoch":1511901886,
                "height":1.837488
                },
                {
                "epoch":1511902186,
                "height":1.811554
                },
                {
                "epoch":1511902486,
                "height":1.788583
                },
                {
                "epoch":1511902786,
                "height":1.768573
                },
                {
                "epoch":1511903086,
                "height":1.751519
                },
                {
                "epoch":1511903386,
                "height":1.737407
                },
                {
                "epoch":1511903686,
                "height":1.726221
                },
                {
                "epoch":1511903986,
                "height":1.717939
                },
                {
                "epoch":1511904286,
                "height":1.712533
                },
                {
                "epoch":1511904586,
                "height":1.709973
                },
                {
                "epoch":1511904886,
                "height":1.710223
                },
                {
                "epoch":1511905186,
                "height":1.713244
                },
                {
                "epoch":1511905486,
                "height":1.718992
                },
                {
                "epoch":1511905786,
                "height":1.727420
                },
                {
                "epoch":1511906086,
                "height":1.738479
                },
                {
                "epoch":1511906386,
                "height":1.752115
                },
                {
                "epoch":1511906686,
                "height":1.768272
                },
                {
                "epoch":1511906986,
                "height":1.786893
                },
                {
                "epoch":1511907286,
                "height":1.807915
                },
                {
                "epoch":1511907586,
                "height":1.831275
                },
                {
                "epoch":1511907886,
                "height":1.856909
                },
                {
                "epoch":1511908186,
                "height":1.884750
                },
                {
                "epoch":1511908486,
                "height":1.914727
                },
                {
                "epoch":1511908786,
                "height":1.946771
                },
                {
                "epoch":1511909086,
                "height":1.980811
                },
                {
                "epoch":1511909386,
                "height":2.016772
                },
                {
                "epoch":1511909686,
                "height":2.054581
                },
                {
                "epoch":1511909986,
                "height":2.094161
                },
                {
                "epoch":1511910286,
                "height":2.135436
                },
                {
                "epoch":1511910586,
                "height":2.178329
                },
                {
                "epoch":1511910886,
                "height":2.222761
                },
                {
                "epoch":1511911186,
                "height":2.268651
                },
                {
                "epoch":1511911486,
                "height":2.315921
                },
                {
                "epoch":1511911786,
                "height":2.364488
                },
                {
                "epoch":1511912086,
                "height":2.414271
                },
                {
                "epoch":1511912386,
                "height":2.465185
                },
                {
                "epoch":1511912686,
                "height":2.517148
                },
                {
                "epoch":1511912986,
                "height":2.570074
                },
                {
                "epoch":1511913286,
                "height":2.623877
                },
                {
                "epoch":1511913586,
                "height":2.678470
                },
                {
                "epoch":1511913886,
                "height":2.733765
                },
                {
                "epoch":1511914186,
                "height":2.789674
                },
                {
                "epoch":1511914486,
                "height":2.846105
                },
                {
                "epoch":1511914786,
                "height":2.902969
                },
                {
                "epoch":1511915086,
                "height":2.960172
                },
                {
                "epoch":1511915386,
                "height":3.017621
                },
                {
                "epoch":1511915686,
                "height":3.075220
                },
                {
                "epoch":1511915986,
                "height":3.132875
                },
                {
                "epoch":1511916286,
                "height":3.190487
                },
                {
                "epoch":1511916586,
                "height":3.247959
                },
                {
                "epoch":1511916886,
                "height":3.305190
                },
                {
                "epoch":1511917186,
                "height":3.362079
                },
                {
                "epoch":1511917486,
                "height":3.418526
                },
                {
                "epoch":1511917786,
                "height":3.474426
                },
                {
                "epoch":1511918086,
                "height":3.529676
                },
                {
                "epoch":1511918386,
                "height":3.584170
                },
                {
                "epoch":1511918686,
                "height":3.637804
                },
                {
                "epoch":1511918986,
                "height":3.690471
                },
                {
                "epoch":1511919286,
                "height":3.742065
                },
                {
                "epoch":1511919586,
                "height":3.792477
                },
                {
                "epoch":1511919886,
                "height":3.841602
                },
                {
                "epoch":1511920186,
                "height":3.889331
                },
                {
                "epoch":1511920486,
                "height":3.935558
                },
                {
                "epoch":1511920786,
                "height":3.980177
                },
                {
                "epoch":1511921086,
                "height":4.023082
                },
                {
                "epoch":1511921386,
                "height":4.064168
                },
                {
                "epoch":1511921686,
                "height":4.103333
                },
                {
                "epoch":1511921986,
                "height":4.140473
                },
                {
                "epoch":1511922286,
                "height":4.175491
                },
                {
                "epoch":1511922586,
                "height":4.208288
                },
                {
                "epoch":1511922886,
                "height":4.238769
                },
                {
                "epoch":1511923186,
                "height":4.266842
                },
                {
                "epoch":1511923486,
                "height":4.292419
                },
                {
                "epoch":1511923786,
                "height":4.315415
                },
                {
                "epoch":1511924086,
                "height":4.335748
                },
                {
                "epoch":1511924386,
                "height":4.353342
                },
                {
                "epoch":1511924686,
                "height":4.368123
                },
                {
                "epoch":1511924986,
                "height":4.380027
                },
                {
                "epoch":1511925286,
                "height":4.388989
                },
                {
                "epoch":1511925586,
                "height":4.394955
                },
                {
                "epoch":1511925886,
                "height":4.397874
                },
                {
                "epoch":1511926186,
                "height":4.397702
                },
                {
                "epoch":1511926486,
                "height":4.394403
                },
                {
                "epoch":1511926786,
                "height":4.387947
                },
                {
                "epoch":1511927086,
                "height":4.378310
                },
                {
                "epoch":1511927386,
                "height":4.365477
                },
                {
                "epoch":1511927686,
                "height":4.349439
                },
                {
                "epoch":1511927986,
                "height":4.330199
                },
                {
                "epoch":1511928286,
                "height":4.307762
                },
                {
                "epoch":1511928586,
                "height":4.282146
                },
                {
                "epoch":1511928886,
                "height":4.253375
                },
                {
                "epoch":1511929186,
                "height":4.221482
                },
                {
                "epoch":1511929486,
                "height":4.186510
                },
                {
                "epoch":1511929786,
                "height":4.148507
                },
                {
                "epoch":1511930086,
                "height":4.107532
                },
                {
                "epoch":1511930386,
                "height":4.063653
                },
                {
                "epoch":1511930686,
                "height":4.016945
                },
                {
                "epoch":1511930986,
                "height":3.967491
                },
                {
                "epoch":1511931286,
                "height":3.915383
                },
                {
                "epoch":1511931586,
                "height":3.860722
                },
                {
                "epoch":1511931886,
                "height":3.803615
                },
                {
                "epoch":1511932186,
                "height":3.744176
                },
                {
                "epoch":1511932486,
                "height":3.682529
                },
                {
                "epoch":1511932786,
                "height":3.618802
                },
                {
                "epoch":1511933086,
                "height":3.553133
                },
                {
                "epoch":1511933386,
                "height":3.485662
                },
                {
                "epoch":1511933686,
                "height":3.416539
                },
                {
                "epoch":1511933986,
                "height":3.345917
                },
                {
                "epoch":1511934286,
                "height":3.273955
                },
                {
                "epoch":1511934586,
                "height":3.200816
                },
                {
                "epoch":1511934886,
                "height":3.126667
                },
                {
                "epoch":1511935186,
                "height":3.051679
                },
                {
                "epoch":1511935486,
                "height":2.976027
                },
                {
                "epoch":1511935786,
                "height":2.899886
                },
                {
                "epoch":1511936086,
                "height":2.823436
                },
                {
                "epoch":1511936386,
                "height":2.746856
                },
                {
                "epoch":1511936686,
                "height":2.670328
                },
                {
                "epoch":1511936986,
                "height":2.594032
                },
                {
                "epoch":1511937286,
                "height":2.518150
                },
                {
                "epoch":1511937586,
                "height":2.442862
                },
                {
                "epoch":1511937886,
                "height":2.368346
                },
                {
                "epoch":1511938186,
                "height":2.294779
                },
                {
                "epoch":1511938486,
                "height":2.222337
                },
                {
                "epoch":1511938786,
                "height":2.151190
                },
                {
                "epoch":1511939086,
                "height":2.081507
                },
                {
                "epoch":1511939386,
                "height":2.013452
                },
                {
                "epoch":1511939686,
                "height":1.947184
                },
                {
                "epoch":1511939986,
                "height":1.882859
                },
                {
                "epoch":1511940286,
                "height":1.820625
                },
                {
                "epoch":1511940586,
                "height":1.760628
                },
                {
                "epoch":1511940886,
                "height":1.703003
                },
                {
                "epoch":1511941186,
                "height":1.647883
                },
                {
                "epoch":1511941486,
                "height":1.595391
                },
                {
                "epoch":1511941786,
                "height":1.545646
                },
                {
                "epoch":1511942086,
                "height":1.498756
                },
                {
                "epoch":1511942386,
                "height":1.454825
                },
                {
                "epoch":1511942686,
                "height":1.413947
                },
                {
                "epoch":1511942986,
                "height":1.376209
                },
                {
                "epoch":1511943286,
                "height":1.341689
                },
                {
                "epoch":1511943586,
                "height":1.310457
                },
                {
                "epoch":1511943886,
                "height":1.282578
                },
                {
                "epoch":1511944186,
                "height":1.258103
                },
                {
                "epoch":1511944486,
                "height":1.237080
                },
                {
                "epoch":1511944786,
                "height":1.219545
                },
                {
                "epoch":1511945086,
                "height":1.205529
                },
                {
                "epoch":1511945386,
                "height":1.195053
                },
                {
                "epoch":1511945686,
                "height":1.188130
                },
                {
                "epoch":1511945986,
                "height":1.184768
                },
                {
                "epoch":1511946286,
                "height":1.184962
                },
                {
                "epoch":1511946586,
                "height":1.188706
                },
                {
                "epoch":1511946886,
                "height":1.195983
                },
                {
                "epoch":1511947186,
                "height":1.206769
                },
                {
                "epoch":1511947486,
                "height":1.221035
                },
                {
                "epoch":1511947786,
                "height":1.238746
                },
                {
                "epoch":1511948086,
                "height":1.259858
                },
                {
                "epoch":1511948386,
                "height":1.284325
                },
                {
                "epoch":1511948686,
                "height":1.312092
                },
                {
                "epoch":1511948986,
                "height":1.343102
                },
                {
                "epoch":1511949286,
                "height":1.377291
                },
                {
                "epoch":1511949586,
                "height":1.414591
                },
                {
                "epoch":1511949886,
                "height":1.454931
                },
                {
                "epoch":1511950186,
                "height":1.498235
                },
                {
                "epoch":1511950486,
                "height":1.544424
                },
                {
                "epoch":1511950786,
                "height":1.593415
                },
                {
                "epoch":1511951086,
                "height":1.645124
                },
                {
                "epoch":1511951386,
                "height":1.699461
                },
                {
                "epoch":1511951686,
                "height":1.756337
                },
                {
                "epoch":1511951986,
                "height":1.815660
                },
                {
                "epoch":1511952286,
                "height":1.877335
                },
                {
                "epoch":1511952586,
                "height":1.941268
                },
                {
                "epoch":1511952886,
                "height":2.007360
                },
                {
                "epoch":1511953186,
                "height":2.075514
                },
                {
                "epoch":1511953486,
                "height":2.145631
                },
                {
                "epoch":1511953786,
                "height":2.217611
                },
                {
                "epoch":1511954086,
                "height":2.291354
                },
                {
                "epoch":1511954386,
                "height":2.366759
                },
                {
                "epoch":1511954686,
                "height":2.443725
                },
                {
                "epoch":1511954986,
                "height":2.522149
                },
                {
                "epoch":1511955286,
                "height":2.601931
                },
                {
                "epoch":1511955586,
                "height":2.682968
                },
                {
                "epoch":1511955886,
                "height":2.765159
                },
                {
                "epoch":1511956186,
                "height":2.848402
                },
                {
                "epoch":1511956486,
                "height":2.932593
                },
                {
                "epoch":1511956786,
                "height":3.017632
                },
                {
                "epoch":1511957086,
                "height":3.103414
                },
                {
                "epoch":1511957386,
                "height":3.189839
                },
                {
                "epoch":1511957686,
                "height":3.276802
                },
                {
                "epoch":1511957986,
                "height":3.364201
                },
                {
                "epoch":1511958286,
                "height":3.451933
                },
                {
                "epoch":1511958586,
                "height":3.539893
                },
                {
                "epoch":1511958886,
                "height":3.627977
                },
                {
                "epoch":1511959186,
                "height":3.716079
                },
                {
                "epoch":1511959486,
                "height":3.804096
                },
                {
                "epoch":1511959786,
                "height":3.891919
                },
                {
                "epoch":1511960086,
                "height":3.979443
                },
                {
                "epoch":1511960386,
                "height":4.066558
                },
                {
                "epoch":1511960686,
                "height":4.153157
                },
                {
                "epoch":1511960986,
                "height":4.239129
                },
                {
                "epoch":1511961286,
                "height":4.324364
                },
                {
                "epoch":1511961586,
                "height":4.408749
                },
                {
                "epoch":1511961886,
                "height":4.492172
                },
                {
                "epoch":1511962186,
                "height":4.574520
                },
                {
                "epoch":1511962486,
                "height":4.655676
                },
                {
                "epoch":1511962786,
                "height":4.735527
                },
                {
                "epoch":1511963086,
                "height":4.813954
                },
                {
                "epoch":1511963386,
                "height":4.890841
                },
                {
                "epoch":1511963686,
                "height":4.966069
                },
                {
                "epoch":1511963986,
                "height":5.039519
                },
                {
                "epoch":1511964286,
                "height":5.111074
                },
                {
                "epoch":1511964586,
                "height":5.180613
                },
                {
                "epoch":1511964886,
                "height":5.248017
                },
                {
                "epoch":1511965186,
                "height":5.313167
                },
                {
                "epoch":1511965486,
                "height":5.375945
                },
                {
                "epoch":1511965786,
                "height":5.436232
                },
                {
                "epoch":1511966086,
                "height":5.493911
                },
                {
                "epoch":1511966386,
                "height":5.548867
                },
                {
                "epoch":1511966686,
                "height":5.600985
                },
                {
                "epoch":1511966986,
                "height":5.650153
                },
                {
                "epoch":1511967286,
                "height":5.696261
                },
                {
                "epoch":1511967586,
                "height":5.739202
                },
                {
                "epoch":1511967886,
                "height":5.778872
                },
                {
                "epoch":1511968186,
                "height":5.815169
                },
                {
                "epoch":1511968486,
                "height":5.847996
                },
                {
                "epoch":1511968786,
                "height":5.877261
                },
                {
                "epoch":1511969086,
                "height":5.902876
                },
                {
                "epoch":1511969386,
                "height":5.924756
                },
                {
                "epoch":1511969686,
                "height":5.942824
                },
                {
                "epoch":1511969986,
                "height":5.957009
                },
                {
                "epoch":1511970286,
                "height":5.967244
                },
                {
                "epoch":1511970586,
                "height":5.973470
                },
                {
                "epoch":1511970886,
                "height":5.975635
                },
                {
                "epoch":1511971186,
                "height":5.973694
                },
                {
                "epoch":1511971486,
                "height":5.967609
                },
                {
                "epoch":1511971786,
                "height":5.957353
                },
                {
                "epoch":1511972086,
                "height":5.942903
                },
                {
                "epoch":1511972386,
                "height":5.924247
                },
                {
                "epoch":1511972686,
                "height":5.901383
                },
                {
                "epoch":1511972986,
                "height":5.874315
                },
                {
                "epoch":1511973286,
                "height":5.843059
                },
                {
                "epoch":1511973586,
                "height":5.807639
                },
                {
                "epoch":1511973886,
                "height":5.768088
                },
                {
                "epoch":1511974186,
                "height":5.724451
                },
                {
                "epoch":1511974486,
                "height":5.676780
                },
                {
                "epoch":1511974786,
                "height":5.625138
                },
                {
                "epoch":1511975086,
                "height":5.569599
                },
                {
                "epoch":1511975386,
                "height":5.510244
                },
                {
                "epoch":1511975686,
                "height":5.447166
                },
                {
                "epoch":1511975986,
                "height":5.380465
                },
                {
                "epoch":1511976286,
                "height":5.310253
                },
                {
                "epoch":1511976586,
                "height":5.236649
                },
                {
                "epoch":1511976886,
                "height":5.159782
                },
                {
                "epoch":1511977186,
                "height":5.079787
                },
                {
                "epoch":1511977486,
                "height":4.996812
                },
                {
                "epoch":1511977786,
                "height":4.911008
                },
                {
                "epoch":1511978086,
                "height":4.822536
                },
                {
                "epoch":1511978386,
                "height":4.731563
                },
                {
                "epoch":1511978686,
                "height":4.638264
                },
                {
                "epoch":1511978986,
                "height":4.542818
                },
                {
                "epoch":1511979286,
                "height":4.445410
                },
                {
                "epoch":1511979586,
                "height":4.346232
                },
                {
                "epoch":1511979886,
                "height":4.245478
                },
                {
                "epoch":1511980186,
                "height":4.143346
                },
                {
                "epoch":1511980486,
                "height":4.040039
                },
                {
                "epoch":1511980786,
                "height":3.935761
                },
                {
                "epoch":1511981086,
                "height":3.830720
                },
                {
                "epoch":1511981386,
                "height":3.725121
                },
                {
                "epoch":1511981686,
                "height":3.619176
                },
                {
                "epoch":1511981986,
                "height":3.513092
                },
                {
                "epoch":1511982286,
                "height":3.407077
                },
                {
                "epoch":1511982586,
                "height":3.301338
                },
                {
                "epoch":1511982886,
                "height":3.196082
                },
                {
                "epoch":1511983186,
                "height":3.091510
                },
                {
                "epoch":1511983486,
                "height":2.987822
                },
                {
                "epoch":1511983786,
                "height":2.885215
                },
                {
                "epoch":1511984086,
                "height":2.783881
                },
                {
                "epoch":1511984386,
                "height":2.684006
                },
                {
                "epoch":1511984686,
                "height":2.585774
                },
                {
                "epoch":1511984986,
                "height":2.489360
                },
                {
                "epoch":1511985286,
                "height":2.394934
                },
                {
                "epoch":1511985586,
                "height":2.302661
                },
                {
                "epoch":1511985886,
                "height":2.212696
                },
                {
                "epoch":1511986186,
                "height":2.125188
                },
                {
                "epoch":1511986486,
                "height":2.040280
                },
                {
                "epoch":1511986786,
                "height":1.958103
                },
                {
                "epoch":1511987086,
                "height":1.878784
                },
                {
                "epoch":1511987386,
                "height":1.802438
                },
                {
                "epoch":1511987686,
                "height":1.729174
                },
                {
                "epoch":1511987986,
                "height":1.659091
                },
                {
                "epoch":1511988286,
                "height":1.592278
                },
                {
                "epoch":1511988586,
                "height":1.528817
                },
                {
                "epoch":1511988886,
                "height":1.468780
                },
                {
                "epoch":1511989186,
                "height":1.412230
                },
                {
                "epoch":1511989486,
                "height":1.359220
                },
                {
                "epoch":1511989786,
                "height":1.309797
                },
                {
                "epoch":1511990086,
                "height":1.263997
                },
                {
                "epoch":1511990386,
                "height":1.221848
                },
                {
                "epoch":1511990686,
                "height":1.183369
                },
                {
                "epoch":1511990986,
                "height":1.148571
                },
                {
                "epoch":1511991286,
                "height":1.117459
                },
                {
                "epoch":1511991586,
                "height":1.090028
                },
                {
                "epoch":1511991886,
                "height":1.066265
                },
                {
                "epoch":1511992186,
                "height":1.046154
                },
                {
                "epoch":1511992486,
                "height":1.029667
                },
                {
                "epoch":1511992786,
                "height":1.016773
                },
                {
                "epoch":1511993086,
                "height":1.007435
                },
                {
                "epoch":1511993386,
                "height":1.001608
                },
                {
                "epoch":1511993686,
                "height":0.999244
                },
                {
                "epoch":1511993986,
                "height":1.000288
                },
                {
                "epoch":1511994286,
                "height":1.004683
                },
                {
                "epoch":1511994586,
                "height":1.012366
                },
                {
                "epoch":1511994886,
                "height":1.023271
                },
                {
                "epoch":1511995186,
                "height":1.037326
                },
                {
                "epoch":1511995486,
                "height":1.054461
                },
                {
                "epoch":1511995786,
                "height":1.074598
                },
                {
                "epoch":1511996086,
                "height":1.097661
                },
                {
                "epoch":1511996386,
                "height":1.123569
                },
                {
                "epoch":1511996686,
                "height":1.152239
                },
                {
                "epoch":1511996986,
                "height":1.183590
                },
                {
                "epoch":1511997286,
                "height":1.217535
                },
                {
                "epoch":1511997586,
                "height":1.253991
                },
                {
                "epoch":1511997886,
                "height":1.292870
                },
                {
                "epoch":1511998186,
                "height":1.334086
                },
                {
                "epoch":1511998486,
                "height":1.377551
                },
                {
                "epoch":1511998786,
                "height":1.423178
                },
                {
                "epoch":1511999086,
                "height":1.470881
                },
                {
                "epoch":1511999386,
                "height":1.520570
                },
                {
                "epoch":1511999686,
                "height":1.572160
                },
                {
                "epoch":1511999986,
                "height":1.625563
                },
                {
                "epoch":1512000286,
                "height":1.680692
                },
                {
                "epoch":1512000586,
                "height":1.737459
                },
                {
                "epoch":1512000886,
                "height":1.795780
                },
                {
                "epoch":1512001186,
                "height":1.855566
                },
                {
                "epoch":1512001486,
                "height":1.916731
                },
                {
                "epoch":1512001786,
                "height":1.979190
                },
                {
                "epoch":1512002086,
                "height":2.042856
                },
                {
                "epoch":1512002386,
                "height":2.107641
                },
                {
                "epoch":1512002686,
                "height":2.173460
                },
                {
                "epoch":1512002986,
                "height":2.240224
                },
                {
                "epoch":1512003286,
                "height":2.307846
                },
                {
                "epoch":1512003586,
                "height":2.376238
                },
                {
                "epoch":1512003886,
                "height":2.445309
                },
                {
                "epoch":1512004186,
                "height":2.514970
                },
                {
                "epoch":1512004486,
                "height":2.585129
                },
                {
                "epoch":1512004786,
                "height":2.655694
                },
                {
                "epoch":1512005086,
                "height":2.726570
                },
                {
                "epoch":1512005386,
                "height":2.797662
                },
                {
                "epoch":1512005686,
                "height":2.868873
                },
                {
                "epoch":1512005986,
                "height":2.940104
                },
                {
                "epoch":1512006286,
                "height":3.011255
                },
                {
                "epoch":1512006586,
                "height":3.082223
                },
                {
                "epoch":1512006886,
                "height":3.152904
                },
                {
                "epoch":1512007186,
                "height":3.223193
                },
                {
                "epoch":1512007486,
                "height":3.292981
                },
                {
                "epoch":1512007786,
                "height":3.362158
                },
                {
                "epoch":1512008086,
                "height":3.430613
                },
                {
                "epoch":1512008386,
                "height":3.498233
                },
                {
                "epoch":1512008686,
                "height":3.564903
                },
                {
                "epoch":1512008986,
                "height":3.630506
                },
                {
                "epoch":1512009286,
                "height":3.694924
                },
                {
                "epoch":1512009586,
                "height":3.758040
                },
                {
                "epoch":1512009886,
                "height":3.819731
                },
                {
                "epoch":1512010186,
                "height":3.879879
                },
                {
                "epoch":1512010486,
                "height":3.938363
                },
                {
                "epoch":1512010786,
                "height":3.995060
                },
                {
                "epoch":1512011086,
                "height":4.049851
                },
                {
                "epoch":1512011386,
                "height":4.102614
                },
                {
                "epoch":1512011686,
                "height":4.153231
                },
                {
                "epoch":1512011986,
                "height":4.201583
                },
                {
                "epoch":1512012286,
                "height":4.247553
                },
                {
                "epoch":1512012586,
                "height":4.291028
                },
                {
                "epoch":1512012886,
                "height":4.331895
                },
                {
                "epoch":1512013186,
                "height":4.370045
                },
                {
                "epoch":1512013486,
                "height":4.405373
                },
                {
                "epoch":1512013786,
                "height":4.437778
                },
                {
                "epoch":1512014086,
                "height":4.467162
                },
                {
                "epoch":1512014386,
                "height":4.493433
                },
                {
                "epoch":1512014686,
                "height":4.516504
                },
                {
                "epoch":1512014986,
                "height":4.536294
                },
                {
                "epoch":1512015286,
                "height":4.552728
                },
                {
                "epoch":1512015586,
                "height":4.565738
                },
                {
                "epoch":1512015886,
                "height":4.575263
                },
                {
                "epoch":1512016186,
                "height":4.581249
                },
                {
                "epoch":1512016486,
                "height":4.583651
                },
                {
                "epoch":1512016786,
                "height":4.582432
                },
                {
                "epoch":1512017086,
                "height":4.577563
                },
                {
                "epoch":1512017386,
                "height":4.569024
                },
                {
                "epoch":1512017686,
                "height":4.556806
                },
                {
                "epoch":1512017986,
                "height":4.540908
                },
                {
                "epoch":1512018286,
                "height":4.521338
                },
                {
                "epoch":1512018586,
                "height":4.498117
                },
                {
                "epoch":1512018886,
                "height":4.471273
                },
                {
                "epoch":1512019186,
                "height":4.440845
                },
                {
                "epoch":1512019486,
                "height":4.406885
                },
                {
                "epoch":1512019786,
                "height":4.369451
                },
                {
                "epoch":1512020086,
                "height":4.328614
                },
                {
                "epoch":1512020386,
                "height":4.284455
                },
                {
                "epoch":1512020686,
                "height":4.237065
                },
                {
                "epoch":1512020986,
                "height":4.186545
                },
                {
                "epoch":1512021286,
                "height":4.133005
                },
                {
                "epoch":1512021586,
                "height":4.076566
                },
                {
                "epoch":1512021886,
                "height":4.017357
                },
                {
                "epoch":1512022186,
                "height":3.955517
                },
                {
                "epoch":1512022486,
                "height":3.891192
                },
                {
                "epoch":1512022786,
                "height":3.824539
                },
                {
                "epoch":1512023086,
                "height":3.755720
                },
                {
                "epoch":1512023386,
                "height":3.684906
                },
                {
                "epoch":1512023686,
                "height":3.612273
                },
                {
                "epoch":1512023986,
                "height":3.538005
                },
                {
                "epoch":1512024286,
                "height":3.462291
                },
                {
                "epoch":1512024586,
                "height":3.385324
                },
                {
                "epoch":1512024886,
                "height":3.307304
                },
                {
                "epoch":1512025186,
                "height":3.228431
                },
                {
                "epoch":1512025486,
                "height":3.148912
                },
                {
                "epoch":1512025786,
                "height":3.068954
                },
                {
                "epoch":1512026086,
                "height":2.988765
                },
                {
                "epoch":1512026386,
                "height":2.908558
                },
                {
                "epoch":1512026686,
                "height":2.828541
                },
                {
                "epoch":1512026986,
                "height":2.748926
                },
                {
                "epoch":1512027286,
                "height":2.669921
                },
                {
                "epoch":1512027586,
                "height":2.591734
                },
                {
                "epoch":1512027886,
                "height":2.514570
                },
                {
                "epoch":1512028186,
                "height":2.438631
                },
                {
                "epoch":1512028486,
                "height":2.364115
                },
                {
                "epoch":1512028786,
                "height":2.291216
                },
                {
                "epoch":1512029086,
                "height":2.220122
                },
                {
                "epoch":1512029386,
                "height":2.151017
                },
                {
                "epoch":1512029686,
                "height":2.084078
                },
                {
                "epoch":1512029986,
                "height":2.019475
                },
                {
                "epoch":1512030286,
                "height":1.957371
                },
                {
                "epoch":1512030586,
                "height":1.897923
                },
                {
                "epoch":1512030886,
                "height":1.841277
                },
                {
                "epoch":1512031186,
                "height":1.787573
                },
                {
                "epoch":1512031486,
                "height":1.736942
                },
                {
                "epoch":1512031786,
                "height":1.689505
                },
                {
                "epoch":1512032086,
                "height":1.645374
                },
                {
                "epoch":1512032386,
                "height":1.604652
                },
                {
                "epoch":1512032686,
                "height":1.567433
                },
                {
                "epoch":1512032986,
                "height":1.533800
                },
                {
                "epoch":1512033286,
                "height":1.503826
                },
                {
                "epoch":1512033586,
                "height":1.477575
                },
                {
                "epoch":1512033886,
                "height":1.455100
                },
                {
                "epoch":1512034186,
                "height":1.436446
                },
                {
                "epoch":1512034486,
                "height":1.421647
                },
                {
                "epoch":1512034786,
                "height":1.410727
                },
                {
                "epoch":1512035086,
                "height":1.403702
                },
                {
                "epoch":1512035386,
                "height":1.400577
                },
                {
                "epoch":1512035686,
                "height":1.401350
                },
                {
                "epoch":1512035986,
                "height":1.406007
                },
                {
                "epoch":1512036286,
                "height":1.414529
                },
                {
                "epoch":1512036586,
                "height":1.426887
                },
                {
                "epoch":1512036886,
                "height":1.443045
                },
                {
                "epoch":1512037186,
                "height":1.462958
                },
                {
                "epoch":1512037486,
                "height":1.486575
                },
                {
                "epoch":1512037786,
                "height":1.513839
                },
                {
                "epoch":1512038086,
                "height":1.544685
                },
                {
                "epoch":1512038386,
                "height":1.579044
                },
                {
                "epoch":1512038686,
                "height":1.616839
                },
                {
                "epoch":1512038986,
                "height":1.657990
                },
                {
                "epoch":1512039286,
                "height":1.702412
                },
                {
                "epoch":1512039586,
                "height":1.750015
                },
                {
                "epoch":1512039886,
                "height":1.800706
                },
                {
                "epoch":1512040186,
                "height":1.854387
                },
                {
                "epoch":1512040486,
                "height":1.910959
                },
                {
                "epoch":1512040786,
                "height":1.970318
                },
                {
                "epoch":1512041086,
                "height":2.032360
                },
                {
                "epoch":1512041386,
                "height":2.096978
                },
                {
                "epoch":1512041686,
                "height":2.164062
                },
                {
                "epoch":1512041986,
                "height":2.233503
                },
                {
                "epoch":1512042286,
                "height":2.305190
                },
                {
                "epoch":1512042586,
                "height":2.379009
                },
                {
                "epoch":1512042886,
                "height":2.454849
                },
                {
                "epoch":1512043186,
                "height":2.532596
                },
                {
                "epoch":1512043486,
                "height":2.612136
                },
                {
                "epoch":1512043786,
                "height":2.693356
                },
                {
                "epoch":1512044086,
                "height":2.776143
                },
                {
                "epoch":1512044386,
                "height":2.860382
                },
                {
                "epoch":1512044686,
                "height":2.945960
                },
                {
                "epoch":1512044986,
                "height":3.032765
                },
                {
                "epoch":1512045286,
                "height":3.120683
                },
                {
                "epoch":1512045586,
                "height":3.209602
                },
                {
                "epoch":1512045886,
                "height":3.299409
                },
                {
                "epoch":1512046186,
                "height":3.389992
                },
                {
                "epoch":1512046486,
                "height":3.481238
                },
                {
                "epoch":1512046786,
                "height":3.573037
                },
                {
                "epoch":1512047086,
                "height":3.665275
                },
                {
                "epoch":1512047386,
                "height":3.757840
                },
                {
                "epoch":1512047686,
                "height":3.850620
                },
                {
                "epoch":1512047986,
                "height":3.943502
                },
                {
                "epoch":1512048286,
                "height":4.036371
                },
                {
                "epoch":1512048586,
                "height":4.129115
                },
                {
                "epoch":1512048886,
                "height":4.221618
                },
                {
                "epoch":1512049186,
                "height":4.313764
                },
                {
                "epoch":1512049486,
                "height":4.405437
                },
                {
                "epoch":1512049786,
                "height":4.496518
                },
                {
                "epoch":1512050086,
                "height":4.586889
                },
                {
                "epoch":1512050386,
                "height":4.676429
                },
                {
                "epoch":1512050686,
                "height":4.765016
                },
                {
                "epoch":1512050986,
                "height":4.852527
                },
                {
                "epoch":1512051286,
                "height":4.938838
                },
                {
                "epoch":1512051586,
                "height":5.023822
                },
                {
                "epoch":1512051886,
                "height":5.107351
                },
                {
                "epoch":1512052186,
                "height":5.189297
                },
                {
                "epoch":1512052486,
                "height":5.269529
                },
                {
                "epoch":1512052786,
                "height":5.347916
                },
                {
                "epoch":1512053086,
                "height":5.424324
                },
                {
                "epoch":1512053386,
                "height":5.498620
                },
                {
                "epoch":1512053686,
                "height":5.570669
                },
                {
                "epoch":1512053986,
                "height":5.640337
                },
                {
                "epoch":1512054286,
                "height":5.707487
                },
                {
                "epoch":1512054586,
                "height":5.771983
                },
                {
                "epoch":1512054886,
                "height":5.833691
                },
                {
                "epoch":1512055186,
                "height":5.892475
                },
                {
                "epoch":1512055486,
                "height":5.948201
                },
                {
                "epoch":1512055786,
                "height":6.000737
                },
                {
                "epoch":1512056086,
                "height":6.049950
                },
                {
                "epoch":1512056386,
                "height":6.095712
                },
                {
                "epoch":1512056686,
                "height":6.137896
                },
                {
                "epoch":1512056986,
                "height":6.176378
                },
                {
                "epoch":1512057286,
                "height":6.211039
                },
                {
                "epoch":1512057586,
                "height":6.241762
                },
                {
                "epoch":1512057886,
                "height":6.268435
                },
                {
                "epoch":1512058186,
                "height":6.290951
                },
                {
                "epoch":1512058486,
                "height":6.309210
                },
                {
                "epoch":1512058786,
                "height":6.323116
                },
                {
                "epoch":1512059086,
                "height":6.332580
                },
                {
                "epoch":1512059386,
                "height":6.337520
                },
                {
                "epoch":1512059686,
                "height":6.337863
                },
                {
                "epoch":1512059986,
                "height":6.333542
                },
                {
                "epoch":1512060286,
                "height":6.324499
                },
                {
                "epoch":1512060586,
                "height":6.310687
                },
                {
                "epoch":1512060886,
                "height":6.292064
                },
                {
                "epoch":1512061186,
                "height":6.268603
                },
                {
                "epoch":1512061486,
                "height":6.240283
                },
                {
                "epoch":1512061786,
                "height":6.207096
                },
                {
                "epoch":1512062086,
                "height":6.169043
                },
                {
                "epoch":1512062386,
                "height":6.126139
                },
                {
                "epoch":1512062686,
                "height":6.078407
                },
                {
                "epoch":1512062986,
                "height":6.025884
                },
                {
                "epoch":1512063286,
                "height":5.968618
                },
                {
                "epoch":1512063586,
                "height":5.906667
                },
                {
                "epoch":1512063886,
                "height":5.840104
                },
                {
                "epoch":1512064186,
                "height":5.769012
                },
                {
                "epoch":1512064486,
                "height":5.693486
                },
                {
                "epoch":1512064786,
                "height":5.613633
                },
                {
                "epoch":1512065086,
                "height":5.529572
                },
                {
                "epoch":1512065386,
                "height":5.441432
                },
                {
                "epoch":1512065686,
                "height":5.349356
                },
                {
                "epoch":1512065986,
                "height":5.253495
                },
                {
                "epoch":1512066286,
                "height":5.154012
                },
                {
                "epoch":1512066586,
                "height":5.051081
                },
                {
                "epoch":1512066886,
                "height":4.944882
                },
                {
                "epoch":1512067186,
                "height":4.835610
                },
                {
                "epoch":1512067486,
                "height":4.723463
                },
                {
                "epoch":1512067786,
                "height":4.608652
                },
                {
                "epoch":1512068086,
                "height":4.491392
                },
                {
                "epoch":1512068386,
                "height":4.371908
                },
                {
                "epoch":1512068686,
                "height":4.250428
                },
                {
                "epoch":1512068986,
                "height":4.127188
                },
                {
                "epoch":1512069286,
                "height":4.002427
                },
                {
                "epoch":1512069586,
                "height":3.876391
                },
                {
                "epoch":1512069886,
                "height":3.749326
                },
                {
                "epoch":1512070186,
                "height":3.621483
                },
                {
                "epoch":1512070486,
                "height":3.493113
                },
                {
                "epoch":1512070786,
                "height":3.364468
                },
                {
                "epoch":1512071086,
                "height":3.235802
                },
                {
                "epoch":1512071386,
                "height":3.107367
                },
                {
                "epoch":1512071686,
                "height":2.979412
                },
                {
                "epoch":1512071986,
                "height":2.852186
                },
                {
                "epoch":1512072286,
                "height":2.725935
                },
                {
                "epoch":1512072586,
                "height":2.600899
                },
                {
                "epoch":1512072886,
                "height":2.477316
                },
                {
                "epoch":1512073186,
                "height":2.355416
                },
                {
                "epoch":1512073486,
                "height":2.235425
                },
                {
                "epoch":1512073786,
                "height":2.117562
                },
                {
                "epoch":1512074086,
                "height":2.002037
                },
                {
                "epoch":1512074386,
                "height":1.889056
                },
                {
                "epoch":1512074686,
                "height":1.778811
                },
                {
                "epoch":1512074986,
                "height":1.671491
                },
                {
                "epoch":1512075286,
                "height":1.567270
                },
                {
                "epoch":1512075586,
                "height":1.466317
                },
                {
                "epoch":1512075886,
                "height":1.368787
                },
                {
                "epoch":1512076186,
                "height":1.274827
                },
                {
                "epoch":1512076486,
                "height":1.184572
                },
                {
                "epoch":1512076786,
                "height":1.098146
                },
                {
                "epoch":1512077086,
                "height":1.015662
                },
                {
                "epoch":1512077386,
                "height":0.937223
                },
                {
                "epoch":1512077686,
                "height":0.862917
                },
                {
                "epoch":1512077986,
                "height":0.792825
                },
                {
                "epoch":1512078286,
                "height":0.727013
                },
                {
                "epoch":1512078586,
                "height":0.665539
                },
                {
                "epoch":1512078886,
                "height":0.608447
                },
                {
                "epoch":1512079186,
                "height":0.555772
                },
                {
                "epoch":1512079486,
                "height":0.507537
                },
                {
                "epoch":1512079786,
                "height":0.463756
                },
                {
                "epoch":1512080086,
                "height":0.424431
                },
                {
                "epoch":1512080386,
                "height":0.389556
                },
                {
                "epoch":1512080686,
                "height":0.359114
                },
                {
                "epoch":1512080986,
                "height":0.333080
                },
                {
                "epoch":1512081286,
                "height":0.311420
                },
                {
                "epoch":1512081586,
                "height":0.294092
                },
                {
                "epoch":1512081886,
                "height":0.281047
                },
                {
                "epoch":1512082186,
                "height":0.272227
                },
                {
                "epoch":1512082486,
                "height":0.267569
                },
                {
                "epoch":1512082786,
                "height":0.267004
                },
                {
                "epoch":1512083086,
                "height":0.270456
                },
                {
                "epoch":1512083386,
                "height":0.277845
                },
                {
                "epoch":1512083686,
                "height":0.289085
                },
                {
                "epoch":1512083986,
                "height":0.304089
                },
                {
                "epoch":1512084286,
                "height":0.322763
                },
                {
                "epoch":1512084586,
                "height":0.345011
                },
                {
                "epoch":1512084886,
                "height":0.370736
                },
                {
                "epoch":1512085186,
                "height":0.399838
                },
                {
                "epoch":1512085486,
                "height":0.432214
                },
                {
                "epoch":1512085786,
                "height":0.467761
                },
                {
                "epoch":1512086086,
                "height":0.506375
                },
                {
                "epoch":1512086386,
                "height":0.547953
                },
                {
                "epoch":1512086686,
                "height":0.592389
                },
                {
                "epoch":1512086986,
                "height":0.639578
                },
                {
                "epoch":1512087286,
                "height":0.689418
                },
                {
                "epoch":1512087586,
                "height":0.741804
                },
                {
                "epoch":1512087886,
                "height":0.796635
                },
                {
                "epoch":1512088186,
                "height":0.853808
                },
                {
                "epoch":1512088486,
                "height":0.913224
                },
                {
                "epoch":1512088786,
                "height":0.974784
                },
                {
                "epoch":1512089086,
                "height":1.038389
                },
                {
                "epoch":1512089386,
                "height":1.103945
                },
                {
                "epoch":1512089686,
                "height":1.171356
                },
                {
                "epoch":1512089986,
                "height":1.240528
                },
                {
                "epoch":1512090286,
                "height":1.311370
                },
                {
                "epoch":1512090586,
                "height":1.383789
                },
                {
                "epoch":1512090886,
                "height":1.457697
                },
                {
                "epoch":1512091186,
                "height":1.533003
                },
                {
                "epoch":1512091486,
                "height":1.609619
                },
                {
                "epoch":1512091786,
                "height":1.687458
                },
                {
                "epoch":1512092086,
                "height":1.766430
                },
                {
                "epoch":1512092386,
                "height":1.846449
                },
                {
                "epoch":1512092686,
                "height":1.927427
                },
                {
                "epoch":1512092986,
                "height":2.009276
                },
                {
                "epoch":1512093286,
                "height":2.091906
                },
                {
                "epoch":1512093586,
                "height":2.175228
                },
                {
                "epoch":1512093886,
                "height":2.259150
                },
                {
                "epoch":1512094186,
                "height":2.343582
                },
                {
                "epoch":1512094486,
                "height":2.428427
                },
                {
                "epoch":1512094786,
                "height":2.513591
                },
                {
                "epoch":1512095086,
                "height":2.598976
                },
                {
                "epoch":1512095386,
                "height":2.684482
                },
                {
                "epoch":1512095686,
                "height":2.770006
                },
                {
                "epoch":1512095986,
                "height":2.855443
                },
                {
                "epoch":1512096286,
                "height":2.940686
                },
                {
                "epoch":1512096586,
                "height":3.025624
                },
                {
                "epoch":1512096886,
                "height":3.110143
                },
                {
                "epoch":1512097186,
                "height":3.194128
                },
                {
                "epoch":1512097486,
                "height":3.277459
                },
                {
                "epoch":1512097786,
                "height":3.360015
                },
                {
                "epoch":1512098086,
                "height":3.441672
                },
                {
                "epoch":1512098386,
                "height":3.522301
                }
                ],
                "rawTideStats": [
                {
                "maxheight":6.337863,
                "minheight":0.267004
                }
                ]
        }
                ,       "webcams": [
                {
                "handle": "Inafog",
                "camid": "InafogCAM3",
                "camindex": "3",
                "assoc_station_id": "KCASANFR79",
                "link": "http://",
                "linktext": "",
                "cameratype": "Phycam 1080p",
                "organization": "",
                "neighborhood": "Mission",
                "zip": "94110",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.75410080",
                "lon": "-122.41188049",
                "updated": "2016-12-13 00:15:23",
                "updated_epoch": "",
                "downloaded": "2016-12-11 21:16:17",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Inafog/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Inafog/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Inafog/3/show.html"
                },
                {
                "handle": "wntdone",
                "camid": "wntdoneCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Traviso-Circle",
                "zip": "94550",
                "city": "Livermore",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.65081787",
                "lon": "-121.78880310",
                "updated": "2016-12-13 00:08:20",
                "updated_epoch": "",
                "downloaded": "2016-11-27 18:27:42",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/n/wntdone/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/n/wntdone/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/wntdone/1/show.html"
                },
                {
                "handle": "ampledata",
                "camid": "ampledataCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR302",
                "link": "http://www.twoeight.net",
                "linktext": "28 Blocks Later",
                "cameratype": "",
                "organization": "28 Blocks Later",
                "neighborhood": "Outer Sunset",
                "zip": "94122",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76002884",
                "lon": "-122.49762726",
                "updated": "2016-12-13 00:19:43",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:48",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/m/ampledata/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/m/ampledata/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/ampledata/1/show.html"
                },
                {
                "handle": "CastroCam",
                "camid": "CastroCamCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANFR114",
                "link": "https://castrocam.net",
                "linktext": "Harvey Milk Plaza pride flag from the CastroCam",
                "cameratype": "Axis 2120",
                "organization": "Harvey Milk Plaza",
                "neighborhood": "The Castro",
                "zip": "94114",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76245499",
                "lon": "-122.43499756",
                "updated": "2016-12-13 00:15:12",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:15:26",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/CastroCam/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM869",
                "camindex": "869",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVB11",
                "organization": "California Department of Transportation",
                "neighborhood": "S880 AT JCT-84",
                "zip": "94555",
                "city": "Fremont",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.56422043",
                "lon": "-122.03964233",
                "updated": "2016-12-13 00:18:50",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:50",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/869/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/869/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/869/show.html"
                },
                {
                "handle": "NeuroDoc",
                "camid": "NeuroDocCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPIEDM8",
                "link": "http://",
                "linktext": "",
                "cameratype": "BloomSky",
                "organization": "",
                "neighborhood": "Piedmont",
                "zip": "94611",
                "city": "Piedmont",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.83071136",
                "lon": "-122.23279572",
                "updated": "2016-12-13 00:19:24",
                "updated_epoch": "",
                "downloaded": "2016-11-03 18:57:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/e/NeuroDoc/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/e/NeuroDoc/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/NeuroDoc/1/show.html"
                },
                {
                "handle": "spacini",
                "camid": "spaciniCAM2",
                "camindex": "2",
                "assoc_station_id": "KCAPLEAS17",
                "link": "http://73.222.131.115:8081",
                "linktext": "Mobile cam",
                "cameratype": "Logitech HD 9000",
                "organization": "",
                "neighborhood": "Parkside-North",
                "zip": "94588",
                "city": "Pleasanton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.68523788",
                "lon": "-121.88873291",
                "updated": "2016-12-13 00:18:32",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:38:28",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/p/spacini/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/p/spacini/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/spacini/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM664",
                "camindex": "664",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-12",
                "organization": "California Department of Transportation",
                "neighborhood": "ALA-580-580 W/O FLYNN ROAD",
                "zip": "94551",
                "city": "Livermore",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.71854401",
                "lon": "-121.67071533",
                "updated": "2016-12-13 00:14:32",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:32",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/664/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/664/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/664/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM290",
                "camindex": "290",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVE73",
                "organization": "California Department of Transportation",
                "neighborhood": "NB US-101 at Spencer Ave",
                "zip": "94966",
                "city": "Sausalito",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85079575",
                "lon": "-122.48973083",
                "updated": "2016-12-13 00:18:01",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:35",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/290/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/290/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/290/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM865",
                "camindex": "865",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV976",
                "organization": "California Department of Transportation",
                "neighborhood": "EOF Suisun Valley Road",
                "zip": "94585",
                "city": "Birds Landing",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.22373199",
                "lon": "-122.12695312",
                "updated": "2016-12-13 00:17:00",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:14",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/865/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/865/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/865/show.html"
                },
                {
                "handle": "CastroCam",
                "camid": "CastroCamCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR114",
                "link": "https://castrocam.net",
                "linktext": "SF Skyline from the CastroCam",
                "cameratype": "",
                "organization": "The CastroCam.net",
                "neighborhood": "The Castro",
                "zip": "94114",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.75971603",
                "lon": "-122.43164062",
                "updated": "2016-12-13 00:12:49",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:13:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/CastroCam/1/show.html"
                },
                {
                "handle": "queenswayinn",
                "camid": "queenswayinnCAM1",
                "camindex": "1",
                "assoc_station_id": "KCANAPA26",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "",
                "zip": "94558",
                "city": "Napa",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.34102631",
                "lon": "-122.11499786",
                "updated": "2016-12-13 00:07:58",
                "updated_epoch": "",
                "downloaded": "2016-11-07 16:08:13",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/q/u/queenswayinn/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/q/u/queenswayinn/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/queenswayinn/1/show.html"
                },
                {
                "handle": "tpkirkp",
                "camid": "tpkirkpCAM1",
                "camindex": "1",
                "assoc_station_id": "KCACONCO16",
                "link": "",
                "linktext": "",
                "cameratype": "Trendnet TV-IP310pi",
                "organization": "K6TPK",
                "neighborhood": "North Concord",
                "zip": "94520",
                "city": "Clyde",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.02865982",
                "lon": "-122.03018188",
                "updated": "2016-12-13 00:10:42",
                "updated_epoch": "",
                "downloaded": "2016-11-08 19:46:57",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/p/tpkirkp/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/p/tpkirkp/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/tpkirkp/1/show.html"
                },
                {
                "handle": "admin",
                "camid": "adminCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR70",
                "link": "http://www.wunderground.com/",
                "linktext": "Visit the Underground Weather",
                "cameratype": "Axis 207W",
                "organization": "",
                "neighborhood": "Sunnyside",
                "zip": "94131",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.73237610",
                "lon": "-122.44290161",
                "updated": "2016-12-13 00:18:51",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:46:32",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/d/admin/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/d/admin/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/admin/1/show.html"
                },
                {
                "handle": "davidjyoung",
                "camid": "davidjyoungCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAFREMO55",
                "link": "http://davidjyoung.com/webcam/weather.jpg",
                "linktext": "",
                "cameratype": "",
                "organization": "homeowner",
                "neighborhood": "Antelope Hills",
                "zip": "94539",
                "city": "Fremont",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.52480316",
                "lon": "-121.93508148",
                "updated": "2016-12-12 15:59:31",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:10:14",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/a/davidjyoung/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/a/davidjyoung/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/davidjyoung/1/show.html"
                },
                {
                "handle": "tvnwun",
                "camid": "tvnwunCAM2",
                "camindex": "2",
                "assoc_station_id": "",
                "link": "http://arbor-studios.com/thirdavecam.html",
                "linktext": "",
                "cameratype": "Sony SNC-RZ30N",
                "organization": "",
                "neighborhood": "",
                "zip": "94404",
                "city": "Foster City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.57001114",
                "lon": "-122.27998352",
                "updated": "2016-12-13 00:19:58",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:03",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/tvnwun/2/show.html"
                },
                {
                "handle": "climbingbum71",
                "camid": "climbingbum71CAM3",
                "camindex": "3",
                "assoc_station_id": "KCABELMO5",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Belmont - Island Park",
                "zip": "94002",
                "city": "Belmont",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.53251266",
                "lon": "-122.26840973",
                "updated": "2016-12-13 00:16:24",
                "updated_epoch": "",
                "downloaded": "2016-12-11 14:48:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/l/climbingbum71/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/l/climbingbum71/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/climbingbum71/3/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM164",
                "camindex": "164",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "3-2",
                "organization": "California Department of Transportation",
                "neighborhood": "Florin Road",
                "zip": "95831",
                "city": "Sacramento",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.49535370",
                "lon": "-121.51647186",
                "updated": "2016-12-13 00:19:12",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:30",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/164/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/164/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/164/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1240",
                "camindex": "1240",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-45",
                "organization": "California Department of Transportation",
                "neighborhood": "E/O MT HOUSE PKWY RWIS",
                "zip": "95391",
                "city": "Tracy",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74106598",
                "lon": "-121.51351166",
                "updated": "2016-12-13 00:17:18",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:32",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1240/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1240/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1240/show.html"
                },
                {
                "handle": "TWMatt",
                "camid": "TWMattCAM1",
                "camindex": "1",
                "assoc_station_id": "KCARICHM10",
                "link": "http://www.tradewindssailing.com/members/weather.shtml",
                "linktext": "Tradewinds Weather",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Marina Bay",
                "zip": "94804",
                "city": "Richmond",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.91283035",
                "lon": "-122.34511566",
                "updated": "2016-12-13 00:16:03",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:24",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/w/TWMatt/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/w/TWMatt/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/TWMatt/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1141",
                "camindex": "1141",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV388",
                "organization": "California Department of Transportation",
                "neighborhood": "N1 JSO PRESIDIO TUNNEL",
                "zip": "94129",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.79137802",
                "lon": "-122.46964264",
                "updated": "2016-12-13 00:15:29",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:45",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1141/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1141/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1141/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM864",
                "camindex": "864",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV711",
                "organization": "California Department of Transportation",
                "neighborhood": "S880 At Paseo Grande OC",
                "zip": "94580",
                "city": "San Lorenzo",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.68111038",
                "lon": "-122.12460327",
                "updated": "2016-12-13 00:18:29",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:38",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/864/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/864/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/864/show.html"
                },
                {
                "handle": "NoeV",
                "camid": "NoeVCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR98",
                "link": "",
                "linktext": "View looking East toward Bernal Hill",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Noe Valley (South)",
                "zip": "94131",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74566269",
                "lon": "-122.43021393",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-11-16 04:40:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/o/NoeV/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/o/NoeV/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/NoeV/1/show.html"
                },
                {
                "handle": "pattonplace",
                "camid": "pattonplaceCAM3",
                "camindex": "3",
                "assoc_station_id": "KCAPACIF34",
                "link": "http://pattonplace.info/weather/MiniWeb",
                "linktext": "Linda Mar, pacifica Ca.",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Linda Mar, Pacifica, Ca",
                "zip": "94044",
                "city": "Pacifica",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.59320068",
                "lon": "-122.48574066",
                "updated": "2016-12-13 00:16:27",
                "updated_epoch": "",
                "downloaded": "2016-12-12 03:24:25",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/a/pattonplace/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/a/pattonplace/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/pattonplace/3/show.html"
                },
                {
                "handle": "iszatso",
                "camid": "iszatsoCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAOAKLA39",
                "link": "http://iszatso.net",
                "linktext": "Weather, Eastmont Hills II",
                "cameratype": "c525",
                "organization": "",
                "neighborhood": "Eastmont Hills",
                "zip": "94605",
                "city": "Oakland",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77137375",
                "lon": "-122.16301727",
                "updated": "2016-12-13 00:17:49",
                "updated_epoch": "",
                "downloaded": "2016-11-07 13:37:51",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/s/iszatso/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/s/iszatso/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/iszatso/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM868",
                "camindex": "868",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV519",
                "organization": "California Department of Transportation",
                "neighborhood": "W580 AT REGATTA BL",
                "zip": "94804",
                "city": "Richmond",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.92081833",
                "lon": "-122.33356476",
                "updated": "2016-12-13 00:19:19",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:25",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/868/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/868/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/868/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM272",
                "camindex": "272",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV516",
                "organization": "California Department of Transportation",
                "neighborhood": "West of Ashby Avenue",
                "zip": "94662",
                "city": "Emeryville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.84672165",
                "lon": "-122.29804230",
                "updated": "2016-12-13 00:16:12",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:12",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/272/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/272/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/272/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1151",
                "camindex": "1151",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVD34",
                "organization": "California Department of Transportation",
                "neighborhood": "E80 at Bay Bridge SAS Tower",
                "zip": "94130",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.81482697",
                "lon": "-122.35912323",
                "updated": "2016-12-13 00:19:38",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1151/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1151/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1151/show.html"
                },
                {
                "handle": "LUMINARIS",
                "camid": "LUMINARISCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANFR306",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "LUMINARIS",
                "neighborhood": "West Portal San Francisco",
                "zip": "94127",
                "city": "SAN FRANCISCO",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74007797",
                "lon": "-122.47079468",
                "updated": "2016-12-13 00:16:59",
                "updated_epoch": "",
                "downloaded": "2016-12-11 19:56:38",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/l/u/LUMINARIS/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/l/u/LUMINARIS/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/LUMINARIS/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM279",
                "camindex": "279",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVD01",
                "organization": "California Department of Transportation",
                "neighborhood": "WB I-80 at Fremont",
                "zip": "94105",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.78478241",
                "lon": "-122.39333344",
                "updated": "2016-12-13 00:14:49",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:50",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/279/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/279/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/279/show.html"
                },
                {
                "handle": "nickcalif",
                "camid": "nickcalifCAM1",
                "camindex": "1",
                "assoc_station_id": "KCARICHM29",
                "link": "http://",
                "linktext": "",
                "cameratype": "BloomSky",
                "organization": "",
                "neighborhood": "Richmond Annex",
                "zip": "94804",
                "city": "Richmond",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.90909576",
                "lon": "-122.30775452",
                "updated": "2016-12-13 00:19:44",
                "updated_epoch": "",
                "downloaded": "2016-11-16 00:04:27",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/i/nickcalif/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/n/i/nickcalif/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/nickcalif/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1351",
                "camindex": "1351",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-48",
                "organization": "California Department of Transportation",
                "neighborhood": "Brannan Island Rd RWIS",
                "zip": "95641",
                "city": "Isleton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.12820053",
                "lon": "-121.58556366",
                "updated": "2016-12-12 19:41:05",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:53",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1351/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1351/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1351/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1147",
                "camindex": "1147",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVA35",
                "organization": "California Department of Transportation",
                "neighborhood": "N680 at Jct 580",
                "zip": "94568",
                "city": "Dublin",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.70148087",
                "lon": "-121.92188263",
                "updated": "2016-12-13 00:18:42",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1147/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1147/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1147/show.html"
                },
                {
                "handle": "ajkimoto",
                "camid": "ajkimotoCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANBR19",
                "link": "",
                "linktext": "",
                "cameratype": "Bloomsky",
                "organization": "Renewable Energy Council",
                "neighborhood": "Crestmoor III",
                "zip": "94066",
                "city": "San Bruno",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.61776733",
                "lon": "-122.43671417",
                "updated": "2016-12-13 00:15:17",
                "updated_epoch": "",
                "downloaded": "2016-11-16 00:24:37",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/j/ajkimoto/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/j/ajkimoto/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/ajkimoto/1/show.html"
                },
                {
                "handle": "randykoch",
                "camid": "randykochCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASUNOL6",
                "link": "http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KCASUNOL6",
                "linktext": "Top of Welch Creek Rd",
                "cameratype": "NetCamXL 3",
                "organization": "",
                "neighborhood": "",
                "zip": "94586",
                "city": "Sunol",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.53305435",
                "lon": "-121.79457855",
                "updated": "2016-12-13 00:17:01",
                "updated_epoch": "",
                "downloaded": "2016-11-08 17:26:20",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/a/randykoch/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/a/randykoch/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/randykoch/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM276",
                "camindex": "276",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV121",
                "organization": "California Department of Transportation",
                "neighborhood": "WB I-80 at Ashby",
                "zip": "94662",
                "city": "Emeryville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85214996",
                "lon": "-122.29994965",
                "updated": "2016-12-13 00:18:41",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/276/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/276/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/276/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1150",
                "camindex": "1150",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVD33",
                "organization": "California Department of Transportation",
                "neighborhood": "W80 at Bay Bridge SAS Tower",
                "zip": "94130",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.81512451",
                "lon": "-122.35949707",
                "updated": "2016-12-13 00:16:55",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:27",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1150/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1150/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1150/show.html"
                },
                {
                "handle": "admin",
                "camid": "adminCAM3",
                "camindex": "3",
                "assoc_station_id": "KCASANFR650",
                "link": "http://www.wunderground.com",
                "linktext": "Wunderground",
                "cameratype": "BloomSky",
                "organization": "Weather Underground",
                "neighborhood": "SunnysideBloomSky",
                "zip": "94131",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/New_York",
                "lat": "37.73212051",
                "lon": "-122.44285583",
                "updated": "2016-12-13 00:12:02",
                "updated_epoch": "",
                "downloaded": "2016-10-31 03:23:57",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/d/admin/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/d/admin/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/admin/3/show.html"
                },
                {
                "handle": "ankitprasad",
                "camid": "ankitprasadCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR914",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "SoMa (7th and Brannan)",
                "zip": "94103",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77360535",
                "lon": "-122.40225220",
                "updated": "2016-12-13 00:16:30",
                "updated_epoch": "",
                "downloaded": "2016-12-02 21:59:18",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/n/ankitprasad/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/n/ankitprasad/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/ankitprasad/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1179",
                "camindex": "1179",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "3-1",
                "organization": "California Department of Transportation",
                "neighborhood": "Pocket Road",
                "zip": "95832",
                "city": "Sacramento",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.48112869",
                "lon": "-121.51052856",
                "updated": "2016-12-13 00:19:28",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:38",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1179/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1179/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1179/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1142",
                "camindex": "1142",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV398",
                "organization": "California Department of Transportation",
                "neighborhood": "E380 JWO JCT101",
                "zip": "94066",
                "city": "San Bruno",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.63375854",
                "lon": "-122.40777588",
                "updated": "2016-12-13 00:17:04",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:21:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1142/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1142/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1142/show.html"
                },
                {
                "handle": "tvnwun",
                "camid": "tvnwunCAM3",
                "camindex": "3",
                "assoc_station_id": "",
                "link": "http://arbor-studios.com/thirdavecam.html",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "",
                "zip": "94404",
                "city": "Foster City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.57001114",
                "lon": "-122.27998352",
                "updated": "2016-12-13 00:15:48",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:48",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/tvnwun/3/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1146",
                "camindex": "1146",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV709",
                "organization": "California Department of Transportation",
                "neighborhood": "N238 JSO ASHLAND AV",
                "zip": "94578",
                "city": "San Leandro",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.68948746",
                "lon": "-122.11741638",
                "updated": "2016-12-13 00:14:17",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:19",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1146/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1146/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1146/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1140",
                "camindex": "1140",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV216",
                "organization": "California Department of Transportation",
                "neighborhood": "S680 AT NORTH MAIN",
                "zip": "94596",
                "city": "Walnut Creek",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.91496277",
                "lon": "-122.06655121",
                "updated": "2016-12-13 00:17:05",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:25",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1140/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1140/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1140/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM284",
                "camindex": "284",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV812",
                "organization": "California Department of Transportation",
                "neighborhood": "SB I-680 at I-80",
                "zip": "94585",
                "city": "Birds Landing",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.20523453",
                "lon": "-122.13825226",
                "updated": "2016-12-13 00:18:15",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/284/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/284/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/284/show.html"
                },
                {
                "handle": "elagache",
                "camid": "elagacheCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAORIND10",
                "link": "http://www.canebas.org/Weather/",
                "linktext": "Canebas weather station",
                "cameratype": "Logitech Webcam Pro 9000",
                "organization": "",
                "neighborhood": "Hills near Del Rey",
                "zip": "94563",
                "city": "Orinda",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85111237",
                "lon": "-122.15499878",
                "updated": "2016-12-13 00:19:49",
                "updated_epoch": "",
                "downloaded": "2016-11-07 18:45:45",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/e/l/elagache/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/e/l/elagache/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/elagache/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM872",
                "camindex": "872",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVE03",
                "organization": "California Department of Transportation",
                "neighborhood": "San Mateo Bridge Substation 2",
                "zip": "94404",
                "city": "Foster City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.58441925",
                "lon": "-122.25087738",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:30",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/872/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/872/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/872/show.html"
                },
                {
                "handle": "wideasleep1",
                "camid": "wideasleep1CAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "None",
                "neighborhood": "Tiburon",
                "zip": "94920",
                "city": "Tiburon",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.87884140",
                "lon": "-122.46309662",
                "updated": "2016-12-13 00:10:27",
                "updated_epoch": "",
                "downloaded": "2016-11-08 19:56:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/i/wideasleep1/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/i/wideasleep1/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/wideasleep1/1/show.html"
                },
                {
                "handle": "svrotkat",
                "camid": "svrotkatCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABETHE5",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Stone Rd",
                "zip": "94511",
                "city": "Bethel Island",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.01372910",
                "lon": "-121.61679840",
                "updated": "2016-12-13 00:19:09",
                "updated_epoch": "",
                "downloaded": "2016-11-08 13:14:49",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/v/svrotkat/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/v/svrotkat/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/svrotkat/1/show.html"
                },
                {
                "handle": "philsweatherstation",
                "camid": "philsweatherstationCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANFR34",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Twin Peaks",
                "zip": "94131",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74875641",
                "lon": "-122.45358276",
                "updated": "2016-12-13 00:19:01",
                "updated_epoch": "",
                "downloaded": "2016-11-25 13:12:02",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/h/philsweatherstation/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/h/philsweatherstation/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/philsweatherstation/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM275",
                "camindex": "275",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVF05",
                "organization": "California Department of Transportation",
                "neighborhood": "South of Pine Valley Road UC",
                "zip": "94583",
                "city": "San Ramon",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.73674774",
                "lon": "-121.95062256",
                "updated": "2016-12-13 00:15:48",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:48",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/275/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/275/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/275/show.html"
                },
                {
                "handle": "VacavilleWeather",
                "camid": "VacavilleWeatherCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAVACAV1",
                "link": "http://vacavilleweather.com",
                "linktext": "Vacaville Weather",
                "cameratype": "Panasonic WV-NS202A",
                "organization": "",
                "neighborhood": "Brown\u0027s Valley",
                "zip": "95688",
                "city": "Vacaville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.39128876",
                "lon": "-121.97102356",
                "updated": "2016-12-13 00:04:12",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:19",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/v/a/VacavilleWeather/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/v/a/VacavilleWeather/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/VacavilleWeather/1/show.html"
                },
                {
                "handle": "botsmaker",
                "camid": "botsmakerCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://home.comcast.net/~botronics/hayward.html",
                "linktext": "Hayward CA Weathercam",
                "cameratype": "Vivitar 3350",
                "organization": "",
                "neighborhood": "SF Bay Area",
                "zip": "94544",
                "city": "Hayward",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.64954376",
                "lon": "-122.12226105",
                "updated": "2016-12-13 00:19:50",
                "updated_epoch": "",
                "downloaded": "2016-11-08 17:07:15",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/o/botsmaker/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/o/botsmaker/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/botsmaker/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1132",
                "camindex": "1132",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-38",
                "organization": "California Department of Transportation",
                "neighborhood": "E/O Mt House Pkwy",
                "zip": "95391",
                "city": "Tracy",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74111176",
                "lon": "-121.51338959",
                "updated": "2016-12-13 00:16:05",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:55",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1132/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1132/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1132/show.html"
                },
                {
                "handle": "modululite",
                "camid": "modululiteCAM2",
                "camindex": "2",
                "assoc_station_id": "KCABERKE66",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Alvarado Ridge",
                "zip": "94705",
                "city": "Berkeley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85989380",
                "lon": "-122.23697662",
                "updated": "2016-12-13 00:18:04",
                "updated_epoch": "",
                "downloaded": "2016-11-16 05:04:22",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/o/modululite/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/o/modululite/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/modululite/2/show.html"
                },
                {
                "handle": "barenjager",
                "camid": "barenjagerCAM1",
                "camindex": "1",
                "assoc_station_id": "KCADALYC1",
                "link": "",
                "linktext": "",
                "cameratype": "Edimax IC-3115W",
                "organization": "",
                "neighborhood": "Mussel Rock - overlooking Pacifica",
                "zip": "94015",
                "city": "Daly City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.66720200",
                "lon": "-122.48923492",
                "updated": "2016-12-13 00:19:01",
                "updated_epoch": "",
                "downloaded": "2016-12-03 15:44:39",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/a/barenjager/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/a/barenjager/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/barenjager/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM273",
                "camindex": "273",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVA49",
                "organization": "California Department of Transportation",
                "neighborhood": "Tassajara",
                "zip": "94588",
                "city": "Pleasanton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.70188141",
                "lon": "-121.87167358",
                "updated": "2016-12-13 00:19:53",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:56",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/273/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/273/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/273/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1274",
                "camindex": "1274",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVD32",
                "organization": "California Department of Transportation",
                "neighborhood": "W80 at Bay Bridge SAS Tower",
                "zip": "94130",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.81512451",
                "lon": "-122.35949707",
                "updated": "2016-12-13 00:18:03",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:09",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1274/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1274/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1274/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM286",
                "camindex": "286",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVB21",
                "organization": "California Department of Transportation",
                "neighborhood": "SB I-880 at 66th",
                "zip": "94622",
                "city": "Oakland",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.75385284",
                "lon": "-122.20860291",
                "updated": "2016-12-13 00:15:10",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:35",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/286/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/286/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/286/show.html"
                },
                {
                "handle": "shelt",
                "camid": "sheltCAM1",
                "camindex": "1",
                "assoc_station_id": "KCATIBUR4",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Mt Tiburon",
                "zip": "94920",
                "city": "Tiburon",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.88853455",
                "lon": "-122.45846558",
                "updated": "2016-12-13 00:17:34",
                "updated_epoch": "",
                "downloaded": "2016-11-08 16:54:19",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/h/shelt/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/h/shelt/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/shelt/1/show.html"
                },
                {
                "handle": "RichCastroWX",
                "camid": "RichCastroWXCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR366",
                "link": "http://wx.ci.net",
                "linktext": "Castro - Twin Peaks view",
                "cameratype": "",
                "organization": "InstantLinux.Net",
                "neighborhood": "The Castro (facing west)",
                "zip": "94114-20",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76181412",
                "lon": "-122.43022919",
                "updated": "2016-12-13 00:17:33",
                "updated_epoch": "",
                "downloaded": "2016-11-07 18:00:28",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/i/RichCastroWX/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/i/RichCastroWX/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/RichCastroWX/1/show.html"
                },
                {
                "handle": "RBWeatherCam",
                "camid": "RBWeatherCamCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPLEAS15",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Foothill High School",
                "zip": "94588",
                "city": "Pleasanton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.66904831",
                "lon": "-121.91646576",
                "updated": "2016-12-13 00:12:30",
                "updated_epoch": "",
                "downloaded": "2016-12-09 01:59:53",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/b/RBWeatherCam/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/b/RBWeatherCam/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/RBWeatherCam/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1149",
                "camindex": "1149",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVD10",
                "organization": "California Department of Transportation",
                "neighborhood": "W80 at SFOBB at Incline",
                "zip": "94623",
                "city": "Oakland",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.82445908",
                "lon": "-122.31742859",
                "updated": "2016-12-13 00:17:56",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:03",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1149/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1149/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1149/show.html"
                },
                {
                "handle": "mu2",
                "camid": "mu2CAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KCASANFR69",
                "linktext": "Fort Funston",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Fort Funston",
                "zip": "94132",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.71267319",
                "lon": "-122.49961853",
                "updated": "2016-12-13 00:16:58",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:20",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/u/mu2/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/u/mu2/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/mu2/1/show.html"
                },
                {
                "handle": "DBeck",
                "camid": "DBeckCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPLEAS8",
                "link": "http://",
                "linktext": "",
                "cameratype": "Pelco Dome",
                "organization": "",
                "neighborhood": "Mt. Diablo",
                "zip": "94523",
                "city": "Pleasant Hill",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.93607330",
                "lon": "-122.08470917",
                "updated": "2016-12-13 00:19:52",
                "updated_epoch": "",
                "downloaded": "2016-12-05 15:52:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/b/DBeck/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/b/DBeck/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/DBeck/1/show.html"
                },
                {
                "handle": "stuschmidt",
                "camid": "stuschmidtCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAHALFM36",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "El Granada",
                "zip": "94019",
                "city": "El Granada",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.50952530",
                "lon": "-122.46846008",
                "updated": "2016-12-13 00:18:25",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:26:19",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/t/stuschmidt/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/t/stuschmidt/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/stuschmidt/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1275",
                "camindex": "1275",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVXXX",
                "organization": "California Department of Transportation",
                "neighborhood": "E80 at Lower Deck - Sterling Onramp",
                "zip": "94105",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.78432083",
                "lon": "-122.39338684",
                "updated": "2016-12-13 00:18:13",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:21",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1275/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1275/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1275/show.html"
                },
                {
                "handle": "RichCastroWX",
                "camid": "RichCastroWXCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANFR366",
                "link": "http://",
                "linktext": "Castro - City view",
                "cameratype": "",
                "organization": "InstantLinux.Net",
                "neighborhood": "The Castro (facing northeast)",
                "zip": "94114",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76181412",
                "lon": "-122.43022919",
                "updated": "2016-12-13 00:18:27",
                "updated_epoch": "",
                "downloaded": "2016-11-08 18:57:18",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/i/RichCastroWX/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/i/RichCastroWX/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/RichCastroWX/2/show.html"
                },
                {
                "handle": "mspencer",
                "camid": "mspencerCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABRENT27",
                "link": "KCABRENT27",
                "linktext": "Brentwood-Mark\u0027s Weather Station",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Brentwood-Mark\u0027s Weather Station",
                "zip": "94513",
                "city": "Brentwood",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.89890671",
                "lon": "-121.72301483",
                "updated": "2016-12-13 00:18:23",
                "updated_epoch": "",
                "downloaded": "2016-11-08 05:00:46",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/s/mspencer/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/m/s/mspencer/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/mspencer/1/show.html"
                },
                {
                "handle": "propheadcity",
                "camid": "propheadcityCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR593",
                "link": "http://www.wunderground.com/personal-weather-station/dashboard?ID=KCASANFR593",
                "linktext": "Golden Gates Heights Weather",
                "cameratype": "AvertX",
                "organization": "",
                "neighborhood": "Golden Gate Heights",
                "zip": "94122",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.75694275",
                "lon": "-122.47083282",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-12-05 04:21:33",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/r/propheadcity/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/r/propheadcity/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/propheadcity/1/show.html"
                },
                {
                "handle": "DougMumma",
                "camid": "DougMummaCAM1",
                "camindex": "1",
                "assoc_station_id": "KCALIVER9",
                "link": "http://www.mumma.org/weather/webcam.jpg",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "",
                "zip": "94550",
                "city": "Livermore",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.65730667",
                "lon": "-121.78097534",
                "updated": "2016-12-13 00:19:46",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/o/DougMumma/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/o/DougMumma/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/DougMumma/1/show.html"
                },
                {
                "handle": "Ingineer",
                "camid": "IngineerCAM2",
                "camindex": "2",
                "assoc_station_id": "KCABERKE28",
                "link": "http://weather.mindfart.com/",
                "linktext": "West Berkeley Weather",
                "cameratype": "Ubiquiti",
                "organization": "Ingineerix Inc",
                "neighborhood": "West Berkeley (facing SF)",
                "zip": "94710",
                "city": "Berkeley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85169983",
                "lon": "-122.28775024",
                "updated": "2016-12-13 00:18:12",
                "updated_epoch": "",
                "downloaded": "2016-11-08 13:43:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Ingineer/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Ingineer/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Ingineer/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM278",
                "camindex": "278",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV304",
                "organization": "California Department of Transportation",
                "neighborhood": "EB I-80 at US-101",
                "zip": "94141",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77025223",
                "lon": "-122.40549469",
                "updated": "2016-12-13 00:17:32",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:50",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/278/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/278/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/278/show.html"
                },
                {
                "handle": "pilotFrankie",
                "camid": "pilotFrankieCAM1",
                "camindex": "1",
                "assoc_station_id": "KCALAFAY37",
                "link": "http://",
                "linktext": "Fravia Vineyard",
                "cameratype": "AmbientCam",
                "organization": "Fravia Vineyard",
                "neighborhood": "Hillside Terrace",
                "zip": "94549",
                "city": "lafayette",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.92102051",
                "lon": "-122.10666656",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-12-11 18:38:36",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/i/pilotFrankie/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/i/pilotFrankie/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/pilotFrankie/1/show.html"
                },
                {
                "handle": "WunderSF",
                "camid": "WunderSFCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR236",
                "link": "http://www.wunderground.com",
                "linktext": "Wunderground",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Wunderground HQ",
                "zip": "94108",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.79327011",
                "lon": "-122.40435028",
                "updated": "2016-12-13 00:19:21",
                "updated_epoch": "",
                "downloaded": "2016-11-16 14:07:53",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/u/WunderSF/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/u/WunderSF/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/WunderSF/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM873",
                "camindex": "873",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVE07",
                "organization": "California Department of Transportation",
                "neighborhood": "San Mateo Bridge Substation 6",
                "zip": "94557",
                "city": "Hayward",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.60815048",
                "lon": "-122.18407440",
                "updated": "2016-12-13 00:17:55",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:16",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/873/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/873/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/873/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM282",
                "camindex": "282",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV305",
                "organization": "California Department of Transportation",
                "neighborhood": "NB US-101 at Candlestick Park",
                "zip": "94160",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.71068192",
                "lon": "-122.39523315",
                "updated": "2016-12-13 00:17:21",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:45",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/282/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/282/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/282/show.html"
                },
                {
                "handle": "tripko",
                "camid": "tripkoCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "",
                "linktext": "Campobello-Blue Rock Hills",
                "cameratype": "Bloomsky",
                "organization": "BlueRock Hills Neighborhood",
                "neighborhood": "Blue Rock Hills",
                "zip": "94591",
                "city": "Vallejo",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.12776947",
                "lon": "-122.20717621",
                "updated": "2016-12-13 00:19:33",
                "updated_epoch": "",
                "downloaded": "2016-11-05 03:21:19",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/r/tripko/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/r/tripko/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/tripko/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM277",
                "camindex": "277",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV106",
                "organization": "California Department of Transportation",
                "neighborhood": "EB I-80 at 80/580/880 IC",
                "zip": "94608",
                "city": "Emeryville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.82844543",
                "lon": "-122.29394531",
                "updated": "2016-12-13 00:15:03",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:31",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/277/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/277/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/277/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1352",
                "camindex": "1352",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-49",
                "organization": "California Department of Transportation",
                "neighborhood": "Jct SR 160 RWIS",
                "zip": "94571",
                "city": "Rio Vista",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.15454865",
                "lon": "-121.67549133",
                "updated": "2016-12-12 23:45:15",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:21:55",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1352/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1352/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1352/show.html"
                },
                {
                "handle": "DRSTORM",
                "camid": "DRSTORMCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASAUSA5",
                "link": "http://",
                "linktext": "North Sausalito (Spring Street)",
                "cameratype": "",
                "organization": "Peronsal",
                "neighborhood": "North Sausalito",
                "zip": "94965",
                "city": "Sausalito",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.85930634",
                "lon": "-122.49730682",
                "updated": "2016-12-13 00:17:28",
                "updated_epoch": "",
                "downloaded": "2016-11-16 16:36:42",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/r/DRSTORM/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/r/DRSTORM/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/DRSTORM/1/show.html"
                },
                {
                "handle": "barenjager",
                "camid": "barenjagerCAM2",
                "camindex": "2",
                "assoc_station_id": "KCADALYC1",
                "link": "",
                "linktext": "",
                "cameratype": "Edimax IC-3115W",
                "organization": "",
                "neighborhood": "Mussel Rock Park",
                "zip": "94015",
                "city": "Daly City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.67049408",
                "lon": "-122.49199677",
                "updated": "2016-12-13 00:19:52",
                "updated_epoch": "",
                "downloaded": "2016-12-04 20:10:43",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/a/barenjager/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/a/barenjager/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/barenjager/2/show.html"
                },
                {
                "handle": "AllenRbrts",
                "camid": "AllenRbrtsCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPLEAS6",
                "link": "www.robertsroost.org",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Grey Eagle",
                "zip": "94566",
                "city": "Pleasanton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.65619659",
                "lon": "-121.84037018",
                "updated": "2016-12-13 00:14:00",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:15",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/l/AllenRbrts/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/l/AllenRbrts/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/AllenRbrts/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM653",
                "camindex": "653",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "10-1",
                "organization": "California Department of Transportation",
                "neighborhood": "SJ-205-HANSON ROAD OC",
                "zip": "95391",
                "city": "Tracy",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.74105072",
                "lon": "-121.51320648",
                "updated": "2016-12-13 00:05:49",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/653/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/653/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/653/show.html"
                },
                {
                "handle": "CastroCam",
                "camid": "CastroCamCAM3",
                "camindex": "3",
                "assoc_station_id": "KCASANFR114",
                "link": "https://castrocam.net",
                "linktext": "San Francisco City Hall Dome 7769 feet from the CastroCam",
                "cameratype": "Axis P1346",
                "organization": "Seen from the CastroCam",
                "neighborhood": "San Francisco Civic Center",
                "zip": "94102",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77843475",
                "lon": "-122.41821289",
                "updated": "2016-12-13 00:15:10",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:36",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/CastroCam/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/CastroCam/3/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM870",
                "camindex": "870",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVB24",
                "organization": "California Department of Transportation",
                "neighborhood": "S880 AT 29TH AV",
                "zip": "94601",
                "city": "Oakland",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77483749",
                "lon": "-122.23344421",
                "updated": "2016-12-12 05:07:39",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:03",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/870/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/870/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/870/show.html"
                },
                {
                "handle": "DougMumma",
                "camid": "DougMummaCAM2",
                "camindex": "2",
                "assoc_station_id": "",
                "link": "http://www.mumma.org",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "",
                "zip": "94550",
                "city": "Livermore",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.65738678",
                "lon": "-121.78093719",
                "updated": "2016-12-13 00:12:12",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:12:56",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/o/DougMumma/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/o/DougMumma/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/DougMumma/2/show.html"
                },
                {
                "handle": "tvnwun",
                "camid": "tvnwunCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://arbor-studios.com/thirdavecam.html",
                "linktext": "KCAFOSTE4",
                "cameratype": "Sony SNC-RZ30N",
                "organization": "",
                "neighborhood": "",
                "zip": "94404",
                "city": "Foster City",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.57001114",
                "lon": "-122.27998352",
                "updated": "2016-12-13 00:19:54",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:21",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/v/tvnwun/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/tvnwun/1/show.html"
                },
                {
                "handle": "Inafog",
                "camid": "InafogCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR79",
                "link": "http://",
                "linktext": "Looking West at Twin Peaks",
                "cameratype": "Panasonic HCM-735",
                "organization": "",
                "neighborhood": "Looking W toward Twin Peaks",
                "zip": "94110",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.75410080",
                "lon": "-122.41188049",
                "updated": "2016-12-13 00:18:19",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:30:12",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Inafog/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/i/n/Inafog/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Inafog/1/show.html"
                },
                {
                "handle": "PCRover",
                "camid": "PCRoverCAM3",
                "camindex": "3",
                "assoc_station_id": "KCAPACIF42",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Linda Mar Ridge, Pacifica, CA",
                "zip": "94044",
                "city": "Pacifica",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.58864212",
                "lon": "-122.48257446",
                "updated": "2016-12-13 00:15:27",
                "updated_epoch": "",
                "downloaded": "2016-11-08 06:32:45",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/c/PCRover/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/c/PCRover/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/PCRover/3/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM863",
                "camindex": "863",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TV102",
                "organization": "California Department of Transportation",
                "neighborhood": "W580 JWO 24 IC",
                "zip": "94609",
                "city": "Oakland",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.82538223",
                "lon": "-122.27289581",
                "updated": "2016-12-13 00:16:52",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:22",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/863/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/863/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/863/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM285",
                "camindex": "285",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVF11",
                "organization": "California Department of Transportation",
                "neighborhood": "I-680 at Greenbrook",
                "zip": "94526",
                "city": "Danville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.79347992",
                "lon": "-121.98325348",
                "updated": "2016-12-13 00:19:59",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:23",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/285/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/285/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/285/show.html"
                },
                {
                "handle": "HillHiker",
                "camid": "HillHikerCAM2",
                "camindex": "2",
                "assoc_station_id": "KCAPLEAS23",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Pleasant Hill, Diablo Valley Estates",
                "zip": "94523",
                "city": "Pleasant Hill",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.97560501",
                "lon": "-122.08995056",
                "updated": "2016-12-13 00:18:25",
                "updated_epoch": "",
                "downloaded": "2016-11-08 16:56:40",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/i/HillHiker/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/i/HillHiker/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/HillHiker/2/show.html"
                },
                {
                "handle": "chris12381",
                "camid": "chris12381CAM2",
                "camindex": "2",
                "assoc_station_id": "",
                "link": "http://www.sanmateocountyairports.org",
                "linktext": "",
                "cameratype": "TriVision 335",
                "organization": "San Mateo County - Half Moon Bay Airport",
                "neighborhood": "Half Moon Bay",
                "zip": "94019",
                "city": "Half Moon Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.51400757",
                "lon": "-122.49509430",
                "updated": "2016-12-13 00:13:41",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:13:40",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/chris12381/2/show.html"
                },
                {
                "handle": "chris12381",
                "camid": "chris12381CAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://www.sanmateocountyairports.org",
                "linktext": "",
                "cameratype": "TriVision 335",
                "organization": "San Mateo County - San Carlos Airport",
                "neighborhood": "San Carlos",
                "zip": "94070",
                "city": "San Carlos",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.51507568",
                "lon": "-122.25051880",
                "updated": "2016-12-13 00:15:49",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/chris12381/1/show.html"
                },
                {
                "handle": "andyknosp",
                "camid": "andyknospCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPLEAS84",
                "link": "http://",
                "linktext": "",
                "cameratype": "Cisco WVC80N",
                "organization": "",
                "neighborhood": "Kottinger Ranch",
                "zip": "94566",
                "city": "Pleasanton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.65050507",
                "lon": "-121.85385895",
                "updated": "2016-12-13 00:19:06",
                "updated_epoch": "",
                "downloaded": "2016-11-16 12:33:33",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/n/andyknosp/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/n/andyknosp/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/andyknosp/1/show.html"
                },
                {
                "handle": "chris12381",
                "camid": "chris12381CAM3",
                "camindex": "3",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "San Mateo County - Half Moon Bay Airport",
                "neighborhood": "Half Moon Bay",
                "zip": "94019",
                "city": "Half Moon BAy",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.51404953",
                "lon": "-122.49517059",
                "updated": "2016-12-13 00:17:24",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:25",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chris12381/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/chris12381/3/show.html"
                },
                {
                "handle": "astroturtle",
                "camid": "astroturtleCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR932",
                "link": "http://astroturtle.com",
                "linktext": "TurtleCam",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Inner Sunset",
                "zip": "94122",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76246262",
                "lon": "-122.47235870",
                "updated": "2016-12-12 23:56:53",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:30:30",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/s/astroturtle/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/s/astroturtle/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/astroturtle/1/show.html"
                },
                {
                "handle": "portofinor",
                "camid": "portofinorCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "www.RNTL.net",
                "linktext": "Mt. Diablo Cam",
                "cameratype": "",
                "organization": "",
                "neighborhood": "",
                "zip": "94526",
                "city": "Danville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.81860733",
                "lon": "-122.01505280",
                "updated": "2016-12-13 00:19:14",
                "updated_epoch": "",
                "downloaded": "2016-11-08 16:47:10",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/o/portofinor/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/o/portofinor/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/portofinor/1/show.html"
                },
                {
                "handle": "astroturtle",
                "camid": "astroturtleCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANFR932",
                "link": "http://astroturtle.com",
                "linktext": "astroturtle.com",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Inner Sunset",
                "zip": "94122",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76246262",
                "lon": "-122.47235870",
                "updated": "2016-12-13 00:14:47",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:15:26",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/s/astroturtle/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/s/astroturtle/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/astroturtle/2/show.html"
                },
                {
                "handle": "Hoytler",
                "camid": "HoytlerCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANLE20",
                "link": "https://twitter.com/droweather",
                "linktext": "Dro Weather Twitter",
                "cameratype": "Bloomsky2",
                "organization": "",
                "neighborhood": "Assumption Parish",
                "zip": "94577",
                "city": "San Leandro",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.71889496",
                "lon": "-122.14614868",
                "updated": "2016-12-13 00:16:26",
                "updated_epoch": "",
                "downloaded": "2016-11-08 16:25:11",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/o/Hoytler/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/o/Hoytler/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Hoytler/1/show.html"
                },
                {
                "handle": "x1962x",
                "camid": "x1962xCAM8",
                "camindex": "8",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Potrero Hill",
                "zip": "94107",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.76346207",
                "lon": "-122.39401245",
                "updated": "2016-12-13 00:12:50",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:22:09",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/x/1/x1962x/8/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/x/1/x1962x/8/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/x1962x/8/show.html"
                },
                {
                "handle": "davisdane",
                "camid": "davisdaneCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "",
                "linktext": "Pittsburg / Antioch HWY Camera",
                "cameratype": "DCS-2330L",
                "organization": "2000 Loveridge Road",
                "neighborhood": "North Park Plaza",
                "zip": "94565",
                "city": "Pittsburg",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.01476288",
                "lon": "-121.86080170",
                "updated": "2016-12-13 00:18:49",
                "updated_epoch": "",
                "downloaded": "2016-12-11 19:59:49",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/a/davisdane/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/a/davisdane/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/davisdane/1/show.html"
                },
                {
                "handle": "shackledtodesk",
                "camid": "shackledtodeskCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANCA52",
                "link": "",
                "linktext": "",
                "cameratype": "RaspberryPi",
                "organization": "",
                "neighborhood": "San Carlos",
                "zip": "94070",
                "city": "San Carlos",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.50699997",
                "lon": "-122.27999878",
                "updated": "2016-12-13 00:19:54",
                "updated_epoch": "",
                "downloaded": "0000-00-00 00:00:00",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/h/shackledtodesk/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/h/shackledtodesk/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/shackledtodesk/1/show.html"
                },
                {
                "handle": "seanw71",
                "camid": "seanw71CAM4",
                "camindex": "4",
                "assoc_station_id": "",
                "link": "http://www.santacruzharbor.org/harborCams/lightHousePoint.html",
                "linktext": "Santa Cruz Harbor",
                "cameratype": "",
                "organization": "Santa Cruz Harbor",
                "neighborhood": "Light House Point",
                "zip": "95062",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.96366501",
                "lon": "-122.00149536",
                "updated": "2016-12-13 00:02:59",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/e/seanw71/4/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/e/seanw71/4/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/seanw71/4/show.html"
                },
                {
                "handle": "SmallIsGood",
                "camid": "SmallIsGoodCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANTA649",
                "link": "",
                "linktext": "",
                "cameratype": "D-Link DCS-5029L",
                "organization": "",
                "neighborhood": "Quail Crossing",
                "zip": "95060",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.98424530",
                "lon": "-122.04387665",
                "updated": "2016-12-13 00:15:22",
                "updated_epoch": "",
                "downloaded": "2016-11-08 03:47:29",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/m/SmallIsGood/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/m/SmallIsGood/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/SmallIsGood/1/show.html"
                },
                {
                "handle": "dcayne",
                "camid": "dcayneCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANGR2",
                "link": "http://www.califcoast.net",
                "linktext": "CalifCoast.Net",
                "cameratype": "Sony SNC-RX570N",
                "organization": "Pinnacle Ranch",
                "neighborhood": "Stage Road",
                "zip": "94074",
                "city": "San Gregorio",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.29713058",
                "lon": "-122.38067627",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:28",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/c/dcayne/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/c/dcayne/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/dcayne/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM1153",
                "camindex": "1153",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "5-39",
                "organization": "California Department of Transportation",
                "neighborhood": "Morrissey Blvd",
                "zip": "95063",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.98859787",
                "lon": "-122.00662231",
                "updated": "2016-12-13 00:17:50",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:16",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1153/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/1153/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/1153/show.html"
                },
                {
                "handle": "SFSeaCliff",
                "camid": "SFSeaCliffCAM4",
                "camindex": "4",
                "assoc_station_id": "KCAPALOA23",
                "link": "http://kj6etn.palo-alto.ca.us",
                "linktext": "Near JLS",
                "cameratype": "Raspberry Pi",
                "organization": "",
                "neighborhood": "Fairmeadow",
                "zip": "94306",
                "city": "Palo Alto",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.41918564",
                "lon": "-122.11727905",
                "updated": "2016-12-13 00:14:20",
                "updated_epoch": "",
                "downloaded": "2016-12-12 06:29:49",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/4/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/4/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/SFSeaCliff/4/show.html"
                },
                {
                "handle": "SFSeaCliff",
                "camid": "SFSeaCliffCAM3",
                "camindex": "3",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "Dlink",
                "organization": "n/a",
                "neighborhood": "Fairmeadow",
                "zip": "94306",
                "city": "Palo Alto",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.41889572",
                "lon": "-122.11765289",
                "updated": "2016-12-13 00:10:38",
                "updated_epoch": "",
                "downloaded": "2016-12-10 02:59:06",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/3/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/3/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/SFSeaCliff/3/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM314",
                "camindex": "314",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "5-23",
                "organization": "California Department of Transportation",
                "neighborhood": "Ocean Street",
                "zip": "95061",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.98802185",
                "lon": "-122.02324677",
                "updated": "2016-12-13 00:19:15",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:31",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/314/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/314/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/314/show.html"
                },
                {
                "handle": "CloudIX",
                "camid": "CloudIXCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABOULD39",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Cloud9",
                "zip": "95006",
                "city": "Boulder Creek",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.22757721",
                "lon": "-122.15077972",
                "updated": "2016-12-13 00:13:46",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:14:09",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/l/CloudIX/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/l/CloudIX/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/CloudIX/1/show.html"
                },
                {
                "handle": "SFSeaCliff",
                "camid": "SFSeaCliffCAM2",
                "camindex": "2",
                "assoc_station_id": "KCAPALOA23",
                "link": "",
                "linktext": "Detail",
                "cameratype": "Raspberry Pi with 5x Lens",
                "organization": "",
                "neighborhood": "Fairmeadow",
                "zip": "94306",
                "city": "Palo Alto",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.41897964",
                "lon": "-122.11715698",
                "updated": "2016-12-13 00:10:30",
                "updated_epoch": "",
                "downloaded": "2016-12-05 04:22:05",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/f/SFSeaCliff/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/SFSeaCliff/2/show.html"
                },
                {
                "handle": "jshluker",
                "camid": "jshlukerCAM1",
                "camindex": "1",
                "assoc_station_id": "KCALOSAL2",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Near Village, Los Altos, CA",
                "zip": "94022",
                "city": "Los Altos",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.37152863",
                "lon": "-122.11380768",
                "updated": "2016-12-13 00:17:14",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:31",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/s/jshluker/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/s/jshluker/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/jshluker/1/show.html"
                },
                {
                "handle": "KCASARATOGA",
                "camid": "KCASARATOGACAM1",
                "camindex": "1",
                "assoc_station_id": "KCASARAT1",
                "link": "http://saratoga-weather.org/",
                "linktext": "Saratoga, CA Weather",
                "cameratype": "Panasonic BB-HCM735A",
                "organization": "Saratoga-Weather.org",
                "neighborhood": "Saratoga",
                "zip": "95070",
                "city": "Saratoga",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.27472305",
                "lon": "-122.02288055",
                "updated": "2016-12-13 00:19:31",
                "updated_epoch": "",
                "downloaded": "2016-11-30 21:44:05",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/k/c/KCASARATOGA/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/k/c/KCASARATOGA/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/KCASARATOGA/1/show.html"
                },
                {
                "handle": "ssnell",
                "camid": "ssnellCAM2",
                "camindex": "2",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Scotts Valley Heights",
                "zip": "95066",
                "city": "Scotts Valley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.04216766",
                "lon": "-122.02097321",
                "updated": "2016-12-13 00:19:11",
                "updated_epoch": "",
                "downloaded": "2016-12-11 01:23:56",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/s/ssnell/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/s/ssnell/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/ssnell/2/show.html"
                },
                {
                "handle": "wstefanc",
                "camid": "wstefancCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAWOODS4",
                "link": "http://",
                "linktext": "",
                "cameratype": "AirCam",
                "organization": "",
                "neighborhood": "",
                "zip": "94062",
                "city": "Woodside",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.46248627",
                "lon": "-122.34681702",
                "updated": "2016-12-13 00:18:33",
                "updated_epoch": "",
                "downloaded": "2016-11-21 15:25:46",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/s/wstefanc/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/s/wstefanc/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/wstefanc/1/show.html"
                },
                {
                "handle": "fbergholz",
                "camid": "fbergholzCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASCOTT28",
                "link": "http://fbergholz.com",
                "linktext": "",
                "cameratype": "",
                "organization": "Ferd\u0027s",
                "neighborhood": "Cross St Scotts Valley, CA",
                "zip": "95066",
                "city": "Scotts Valley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.05776596",
                "lon": "-122.00425720",
                "updated": "2016-12-13 00:17:11",
                "updated_epoch": "",
                "downloaded": "2016-12-12 23:29:08",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/b/fbergholz/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/b/fbergholz/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/fbergholz/1/show.html"
                },
                {
                "handle": "cwdwx",
                "camid": "cwdwxCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAHALFM28",
                "link": "http://adequatebird.com",
                "linktext": "Adequate Bird",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Canada Cove",
                "zip": "94019",
                "city": "Half Moon Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.42992401",
                "lon": "-122.43190765",
                "updated": "2016-12-13 00:15:48",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/w/cwdwx/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/w/cwdwx/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cwdwx/1/show.html"
                },
                {
                "handle": "gstev46412",
                "camid": "gstev46412CAM2",
                "camindex": "2",
                "assoc_station_id": "KCASCOTT33",
                "link": "http://www.wunderground.com/personal-weather-station/dashboard?ID=KCASCOTT33",
                "linktext": "Casa Way PWS",
                "cameratype": "Panasonic KX-HCM280",
                "organization": "e-ctrl",
                "neighborhood": "Casa Way",
                "zip": "95066",
                "city": "Scotts Valley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.06691360",
                "lon": "-122.01015472",
                "updated": "2016-12-12 23:37:00",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:16:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/g/s/gstev46412/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/g/s/gstev46412/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/gstev46412/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM312",
                "camindex": "312",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "5-21",
                "organization": "California Department of Transportation",
                "neighborhood": "Emeline Ave",
                "zip": "95063",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.98954391",
                "lon": "-122.01958466",
                "updated": "2016-12-13 00:16:57",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:17:13",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/312/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/312/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/312/show.html"
                },
                {
                "handle": "cwdwx",
                "camid": "cwdwxCAM2",
                "camindex": "2",
                "assoc_station_id": "KCAHALFM28",
                "link": "http://adequatebird.com",
                "linktext": "Adequate Bird",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Canada Cove",
                "zip": "94019",
                "city": "Half Moon Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.43000031",
                "lon": "-122.43000031",
                "updated": "2016-12-13 00:13:18",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:13:18",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/w/cwdwx/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/w/cwdwx/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cwdwx/2/show.html"
                },
                {
                "handle": "a434mhz",
                "camid": "a434mhzCAM2",
                "camindex": "2",
                "assoc_station_id": "",
                "link": "",
                "linktext": "",
                "cameratype": "Moonglow Allsky",
                "organization": "",
                "neighborhood": "Midtown",
                "zip": "94306",
                "city": "palo alto",
                "state": "CA",
                "country": "US",
                "tzname": "America/Tijuana",
                "lat": "37.42984009",
                "lon": "-122.12597656",
                "updated": "2016-12-12 02:04:16",
                "updated_epoch": "",
                "downloaded": "2016-12-06 20:15:22",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/4/a434mhz/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/4/a434mhz/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/a434mhz/2/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM313",
                "camindex": "313",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "5-22",
                "organization": "California Department of Transportation",
                "neighborhood": "Fishhook",
                "zip": "95063",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "36.99204636",
                "lon": "-122.01975250",
                "updated": "2016-12-13 00:19:05",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:31",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/313/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/313/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/313/show.html"
                },
                {
                "handle": "dbw77",
                "camid": "dbw77CAM2",
                "camindex": "2",
                "assoc_station_id": "KCALOSGA2",
                "link": "http://",
                "linktext": "",
                "cameratype": "FOSCAM FI8906W",
                "organization": "",
                "neighborhood": "",
                "zip": "95033",
                "city": "Los Gatos",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.16051865",
                "lon": "-122.01085663",
                "updated": "2016-12-13 00:13:15",
                "updated_epoch": "",
                "downloaded": "2016-03-17 18:03:15",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/b/dbw77/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/d/b/dbw77/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/dbw77/2/show.html"
                },
                {
                "handle": "BruceCrawford",
                "camid": "BruceCrawfordCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAHALFM26",
                "link": "http://www.FlyPCC.org",
                "linktext": "PCC",
                "cameratype": "",
                "organization": "PCC",
                "neighborhood": "PCC R/C Flying Field",
                "zip": "94019",
                "city": "Half Moon Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.41485596",
                "lon": "-122.42326355",
                "updated": "2016-12-13 00:15:31",
                "updated_epoch": "",
                "downloaded": "2016-11-30 19:37:59",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/r/BruceCrawford/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/r/BruceCrawford/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/BruceCrawford/1/show.html"
                },
                {
                "handle": "fly44d",
                "camid": "fly44dCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABOULD26",
                "link": "http://",
                "linktext": "Deck Cam",
                "cameratype": "Logitech C210",
                "organization": "",
                "neighborhood": "Wildwood",
                "zip": "95006",
                "city": "Boulder Creek",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.15303421",
                "lon": "-122.13677979",
                "updated": "2016-12-13 00:18:35",
                "updated_epoch": "",
                "downloaded": "2016-11-23 11:00:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/l/fly44d/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/l/fly44d/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/fly44d/1/show.html"
                },
                {
                "handle": "bkreid",
                "camid": "bkreidCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPALOA15",
                "link": "http://vegetablegroup.org/weather/",
                "linktext": "",
                "cameratype": "Toshiba IK-WB80A",
                "organization": "",
                "neighborhood": "Looking south",
                "zip": "94301",
                "city": "Palo Alto",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.43561554",
                "lon": "-122.14177704",
                "updated": "2016-12-13 00:10:29",
                "updated_epoch": "",
                "downloaded": "2016-11-08 15:02:27",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/k/bkreid/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/k/bkreid/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/bkreid/1/show.html"
                },
                {
                "handle": "bdairfield",
                "camid": "bdairfieldCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANTA382",
                "link": "http://www.dvq.com",
                "linktext": "DVQ",
                "cameratype": "",
                "organization": "DVQ",
                "neighborhood": "Bonny Doon Airport",
                "zip": "95060",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.07060623",
                "lon": "-122.12740326",
                "updated": "2016-12-13 00:15:50",
                "updated_epoch": "",
                "downloaded": "2016-11-08 10:03:39",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/d/bdairfield/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/d/bdairfield/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/bdairfield/1/show.html"
                },
                {
                "handle": "JoeChristy",
                "camid": "JoeChristyCAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "",
                "cameratype": "BloomSky",
                "organization": "",
                "neighborhood": "Bonny Doon",
                "zip": "95060",
                "city": "Santa Cruz",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.05436325",
                "lon": "-122.11503601",
                "updated": "2016-12-13 00:16:08",
                "updated_epoch": "",
                "downloaded": "2016-11-05 02:47:34",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/o/JoeChristy/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/o/JoeChristy/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/JoeChristy/1/show.html"
                },
                {
                "handle": "fastspider",
                "camid": "fastspiderCAM1",
                "camindex": "1",
                "assoc_station_id": "KCALOSGA22",
                "link": "",
                "linktext": "Skyline View",
                "cameratype": "",
                "organization": "Disorganized",
                "neighborhood": "Skyline",
                "zip": "95033",
                "city": "Los Gatos",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.20744705",
                "lon": "-122.04929352",
                "updated": "2016-12-13 00:18:41",
                "updated_epoch": "",
                "downloaded": "2016-11-27 16:36:50",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/a/fastspider/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/f/a/fastspider/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/fastspider/1/show.html"
                },
                {
                "handle": "Radius9999",
                "camid": "Radius9999CAM1",
                "camindex": "1",
                "assoc_station_id": "",
                "link": "http://",
                "linktext": "Lompico Cam",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Felton, Lompico, CA",
                "zip": "95018",
                "city": "Felton",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.10527802",
                "lon": "-122.05110931",
                "updated": "2016-12-13 00:18:51",
                "updated_epoch": "",
                "downloaded": "0000-00-00 00:00:00",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/a/Radius9999/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/a/Radius9999/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Radius9999/1/show.html"
                },
                {
                "handle": "cvogeo",
                "camid": "cvogeoCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPETAL9",
                "link": "http://",
                "linktext": "",
                "cameratype": "AmbientCAM",
                "organization": "na",
                "neighborhood": "East Petaluma, CA",
                "zip": "94954",
                "city": "Petaluma",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.25377655",
                "lon": "-122.62054443",
                "updated": "2016-12-13 00:19:08",
                "updated_epoch": "",
                "downloaded": "2016-11-08 15:27:44",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/v/cvogeo/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/v/cvogeo/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cvogeo/1/show.html"
                },
                {
                "handle": "SerenoDelMar",
                "camid": "SerenoDelMarCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABODEG3",
                "link": "http://home.comcast.net/~leonardweber/site/?/home/",
                "linktext": "West Facing Camera",
                "cameratype": "Logitech HD Pro C920",
                "organization": "",
                "neighborhood": "Sereno Del Mar",
                "zip": "94923",
                "city": "Bodega Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.38171768",
                "lon": "-123.07574463",
                "updated": "2016-12-13 00:19:18",
                "updated_epoch": "",
                "downloaded": "2016-11-30 21:10:30",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/e/SerenoDelMar/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/s/e/SerenoDelMar/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/SerenoDelMar/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM274",
                "camindex": "274",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVE83",
                "organization": "California Department of Transportation",
                "neighborhood": "I-580",
                "zip": "94915",
                "city": "San Rafael",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.96193314",
                "lon": "-122.50971222",
                "updated": "2016-12-13 00:18:14",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:18:42",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/274/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/274/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/274/show.html"
                },
                {
                "handle": "JWPAGE",
                "camid": "JWPAGECAM1",
                "camindex": "1",
                "assoc_station_id": "KCAPETAL45",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "MARINCC",
                "neighborhood": "Park Place",
                "zip": "94954",
                "city": "Petaluma",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.25960541",
                "lon": "-122.63813782",
                "updated": "2016-12-13 00:19:00",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:09:43",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/w/JWPAGE/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/w/JWPAGE/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/JWPAGE/1/show.html"
                },
                {
                "handle": "CVGonzalves",
                "camid": "CVGonzalvesCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAFORES45",
                "link": "http://",
                "linktext": "",
                "cameratype": "Bloomsky",
                "organization": "Russian River Vineyards",
                "neighborhood": "Forestville",
                "zip": "95436",
                "city": "Forestville",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.46545029",
                "lon": "-122.88332367",
                "updated": "2016-12-13 00:15:44",
                "updated_epoch": "",
                "downloaded": "2016-11-07 01:48:55",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/v/CVGonzalves/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/v/CVGonzalves/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/CVGonzalves/1/show.html"
                },
                {
                "handle": "thegn",
                "camid": "thegnCAM1",
                "camindex": "1",
                "assoc_station_id": "KCAFAIRF51",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "East Fairfax",
                "zip": "94930",
                "city": "fairfax",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.98697281",
                "lon": "-122.58450317",
                "updated": "2016-12-13 00:19:58",
                "updated_epoch": "",
                "downloaded": "2016-12-06 20:15:21",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/h/thegn/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/t/h/thegn/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/thegn/1/show.html"
                },
                {
                "handle": "WunderYakuza",
                "camid": "WunderYakuzaCAM8",
                "camindex": "8",
                "assoc_station_id": "",
                "link": "http://www.nps.gov/pore/index.htm",
                "linktext": "Point Reyes National Seashore",
                "cameratype": "Olympus",
                "organization": "National Park Service",
                "neighborhood": "Point Reyes",
                "zip": "94937",
                "city": "Point Reyes National Seashore",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.99670410",
                "lon": "-123.01919556",
                "updated": "2016-12-13 00:06:42",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:20:42",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/u/WunderYakuza/8/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/w/u/WunderYakuza/8/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/WunderYakuza/8/show.html"
                },
                {
                "handle": "lparkerwu66",
                "camid": "lparkerwu66CAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANTA708",
                "link": "http://",
                "linktext": "",
                "cameratype": "AMBIENTCAMHD",
                "organization": "",
                "neighborhood": "Oakmont, Oak Shadow",
                "zip": "95409",
                "city": "Santa Rosa",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.43581009",
                "lon": "-122.58711243",
                "updated": "2016-12-13 00:11:48",
                "updated_epoch": "",
                "downloaded": "2016-11-08 20:25:15",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/l/p/lparkerwu66/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/l/p/lparkerwu66/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/lparkerwu66/1/show.html"
                },
                {
                "handle": "gjpc",
                "camid": "gjpcCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR107",
                "link": "http://www.oceanbeach.org",
                "linktext": "GGNRA NW",
                "cameratype": "EdiMax 3030",
                "organization": "CircleSoft",
                "neighborhood": "GGNRA",
                "zip": "94121",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77195740",
                "lon": "-122.51078796",
                "updated": "2016-12-13 00:18:22",
                "updated_epoch": "",
                "downloaded": "2016-11-08 15:49:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/g/j/gjpc/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/g/j/gjpc/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/gjpc/1/show.html"
                },
                {
                "handle": "rd95401",
                "camid": "rd95401CAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANTA447",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Olivet & Piner Roads",
                "zip": "95401",
                "city": "Santa Rosa",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.46214676",
                "lon": "-122.81562805",
                "updated": "2016-12-13 00:19:32",
                "updated_epoch": "",
                "downloaded": "2016-11-30 22:11:20",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/d/rd95401/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/d/rd95401/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/rd95401/1/show.html"
                },
                {
                "handle": "Havec",
                "camid": "HavecCAM2",
                "camindex": "2",
                "assoc_station_id": "KCAMILLV27",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "Residence",
                "neighborhood": "Hillside Ave, Mill Valley, CA USA",
                "zip": "94941",
                "city": "Mill Valley",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.90957260",
                "lon": "-122.54431152",
                "updated": "2016-12-13 00:18:39",
                "updated_epoch": "",
                "downloaded": "2016-12-11 20:33:04",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/a/Havec/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/h/a/Havec/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/Havec/2/show.html"
                },
                {
                "handle": "jsntg",
                "camid": "jsntgCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANFR99",
                "link": "http://ob-kc.com",
                "linktext": "Ocean Beach Kelly\u0027s Cove",
                "cameratype": "Canon Powershot G3",
                "organization": "OB-KC.COM",
                "neighborhood": "Ocean Beach - Kelly\u0027s Cove",
                "zip": "94121",
                "city": "San Francisco",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.77495575",
                "lon": "-122.51038361",
                "updated": "2016-12-13 00:15:28",
                "updated_epoch": "",
                "downloaded": "2016-03-12 15:53:03",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/s/jsntg/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/j/s/jsntg/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/jsntg/1/show.html"
                },
                {
                "handle": "cadot1",
                "camid": "cadot1CAM874",
                "camindex": "874",
                "assoc_station_id": "",
                "link": "http://www.dot.ca.gov/dist3/departments/traffic/cameras/",
                "linktext": "California Department of Transportation",
                "cameratype": "4-TVE86",
                "organization": "California Department of Transportation",
                "neighborhood": "S101 AT IGNACIO BL",
                "zip": "94949",
                "city": "Novato",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.06616974",
                "lon": "-122.53706360",
                "updated": "2016-12-13 00:19:05",
                "updated_epoch": "",
                "downloaded": "2016-12-13 00:19:24",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/874/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/a/cadot1/874/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/cadot1/874/show.html"
                },
                {
                "handle": "KG6WYP",
                "camid": "KG6WYPCAM1",
                "camindex": "1",
                "assoc_station_id": "KCABODEG10",
                "link": "https://www.wunderground.com/webcams/KG6WYP/1/show.html",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "El Camino Bella",
                "zip": "94923",
                "city": "Bodega Bay",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.37779236",
                "lon": "-123.07748413",
                "updated": "2016-12-13 00:00:33",
                "updated_epoch": "",
                "downloaded": "2016-12-12 20:55:48",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/k/g/KG6WYP/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/k/g/KG6WYP/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/KG6WYP/1/show.html"
                },
                {
                "handle": "rmamer",
                "camid": "rmamerCAM1",
                "camindex": "1",
                "assoc_station_id": "KCASEBAS82",
                "link": "http://",
                "linktext": "",
                "cameratype": "bloomsky",
                "organization": "",
                "neighborhood": "Bloomfield Road Sebastopol, CA",
                "zip": "95472",
                "city": "Sebastopol",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.36548615",
                "lon": "-122.81896210",
                "updated": "2016-12-13 00:10:39",
                "updated_epoch": "",
                "downloaded": "2016-12-01 15:02:01",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/m/rmamer/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/m/rmamer/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/rmamer/1/show.html"
                },
                {
                "handle": "chrisrosa",
                "camid": "chrisrosaCAM1",
                "camindex": "1",
                "assoc_station_id": "KCANOVAT57",
                "link": "http://www.drivesaversdatarecovery.com",
                "linktext": "DriveSavers Data Recovery",
                "cameratype": "BloomSky",
                "organization": "DriveSavers, Inc.",
                "neighborhood": "Bel Marin Keys",
                "zip": "94949",
                "city": "Novato",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.07662964",
                "lon": "-122.53539276",
                "updated": "2016-12-13 00:17:06",
                "updated_epoch": "",
                "downloaded": "2016-11-08 02:26:24",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chrisrosa/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/c/h/chrisrosa/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/chrisrosa/1/show.html"
                },
                {
                "handle": "redhill16",
                "camid": "redhill16CAM1",
                "camindex": "1",
                "assoc_station_id": "KCASANAN1",
                "link": "http://www.bowkera.com",
                "linktext": "San Anselmo Red Hill Camera",
                "cameratype": "Insteon",
                "organization": "",
                "neighborhood": "Red Hill",
                "zip": "94960",
                "city": "San Anselmo",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "37.98078156",
                "lon": "-122.55938721",
                "updated": "2016-12-13 00:19:24",
                "updated_epoch": "",
                "downloaded": "2016-11-16 16:53:43",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/e/redhill16/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/r/e/redhill16/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/redhill16/1/show.html"
                },
                {
                "handle": "AR323",
                "camid": "AR323CAM2",
                "camindex": "2",
                "assoc_station_id": "KCASEBAS1",
                "link": "http://",
                "linktext": "",
                "cameratype": "Bloomsky",
                "organization": "Sebastopol Sky",
                "neighborhood": "Sebastopol, CA",
                "zip": "95472",
                "city": "Sebastopol",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.38024902",
                "lon": "-122.82374573",
                "updated": "2016-12-12 18:47:45",
                "updated_epoch": "",
                "downloaded": "2016-11-02 23:36:07",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/r/AR323/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/a/r/AR323/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/AR323/2/show.html"
                },
                {
                "handle": "pizote",
                "camid": "pizoteCAM2",
                "camindex": "2",
                "assoc_station_id": "KCASANTA586",
                "link": "http://",
                "linktext": "",
                "cameratype": "",
                "organization": "",
                "neighborhood": "Manor Estates",
                "zip": "95403",
                "city": "Santa Rosa",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.45536423",
                "lon": "-122.75846863",
                "updated": "2016-12-13 00:10:23",
                "updated_epoch": "",
                "downloaded": "2016-11-07 04:39:25",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/i/pizote/2/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/p/i/pizote/2/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/pizote/2/show.html"
                },
                {
                "handle": "bsears",
                "camid": "bsearsCAM1",
                "camindex": "1",
                "assoc_station_id": "KCACAZAD7",
                "link": "http://",
                "linktext": "",
                "cameratype": "Ambient Weather AMBIENTCAMHDA",
                "organization": "",
                "neighborhood": "Kidd Creek",
                "zip": "95421",
                "city": "Cazadero",
                "state": "CA",
                "country": "US",
                "tzname": "America/Los_Angeles",
                "lat": "38.49599838",
                "lon": "-123.06999969",
                "updated": "2016-12-13 00:19:18",
                "updated_epoch": "",
                "downloaded": "2016-11-08 19:58:26",
                "isrecent": "1",
                "CURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/s/bsears/1/current.jpg?t=1511753040",
                "WIDGETCURRENTIMAGEURL": "http://icons.wunderground.com/webcamramdisk/b/s/bsears/1/widget.jpg?t=1511753040",
                "CAMURL": "http://www.wunderground.com/webcams/bsears/1/show.html"
                }
        ]
}
`

	ar := &wu.ApiResponse{}
	if err := json.Unmarshal([]byte(text), &ar); err != nil {
		panic(err)
	}
	return ar
}

// Returns a new ApiResponse with a planner set.
func GetTestPlanner() *wu.ApiResponse{
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

	ar := &wu.ApiResponse{}
	if err := json.Unmarshal([]byte(text), &ar); err != nil {
		panic(err)
	}
	return ar
}
