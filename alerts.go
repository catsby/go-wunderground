package wunderground

// Use to request the alerts feature in Service.Request
var FAlerts = "alerts"

type Alerts []struct {
	Type         string `json:"type"`
	Description  string `json:"description"`
	Date         string `json:"date"`
	DateEpoch    string `json:"date_epoch"`
	Expires      string `json:"expires"`
	ExpiresEpoch string `json:"expires_epoch"`
	Message      string `json:"message"`
	Phenomena    string `json:"phenomena"`
	Significance string `json:"significance"`
	ZONES        []struct {
		State string `json:"state"`
		ZONE  string `json:"ZONE"`
	} `json:"ZONES,omitempty"`
	StormBased struct {
		Vertices []struct {
			Lat string `json:"lat"`
			Lon string `json:"lon"`
		} `json:"vertices"`
		VertexCount int `json:"Vertex_count"`
		StormInfo   struct {
			TimeEpoch   int     `json:"time_epoch"`
			MotionDeg   int     `json:"Motion_deg"`
			MotionSpd   int     `json:"Motion_spd"`
			PositionLat float64 `json:"position_lat"`
			PositionLon float64 `json:"position_lon"`
		} `json:"stormInfo"`
	} `json:"StormBased,omitempty"`
	WtypeMeteoalarm            string `json:"wtype_meteoalarm,omitempty"`
	WtypeMeteoalarmName        string `json:"wtype_meteoalarm_name,omitempty"`
	LevelMeteoalarm            string `json:"level_meteoalarm,omitempty"`
	LevelMeteoalarmName        string `json:"level_meteoalarm_name,omitempty"`
	LevelMeteoalarmDescription string `json:"level_meteoalarm_description,omitempty"`
	Attribution                string `json:"attribution,omitempty"`
}
