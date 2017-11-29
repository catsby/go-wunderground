package wunderground

// Use to request the webcams feature in Service.Request
var FWebCams = "webcams"

type WebCams []struct {
	Handle                string `json:"handle"`
	Camid                 string `json:"camid"`
	Camindex              string `json:"camindex"`
	AssocStationID        string `json:"assoc_station_id"`
	Link                  string `json:"link"`
	Linktext              string `json:"linktext"`
	Cameratype            string `json:"cameratype"`
	Organization          string `json:"organization"`
	Neighborhood          string `json:"neighborhood"`
	Zip                   string `json:"zip"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Country               string `json:"country"`
	Tzname                string `json:"tzname"`
	Lat                   string `json:"lat"`
	Lon                   string `json:"lon"`
	Updated               string `json:"updated"`
	Downloaded            string `json:"downloaded"`
	Isrecent              string `json:"isrecent"`
	CURRENTIMAGEURL       string `json:"CURRENTIMAGEURL"`
	WIDGETCURRENTIMAGEURL string `json:"WIDGETCURRENTIMAGEURL"`
	CAMURL                string `json:"CAMURL"`
}
