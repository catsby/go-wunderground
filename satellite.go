package wunderground

// Use to request the satellite feature in Service.Request
var FSatellite = "satellite"

type Satellite struct {
	ImageURL    string `json:"image_url"`
	ImageURLIr4 string `json:"image_url_ir4"`
	ImageURLVis string `json:"image_url_vis"`
}
