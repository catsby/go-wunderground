package wunderground

type Wunderground interface {
  Conditions(query interface{}) (*ConditionsResponse, error)
  Forecast(query interface{}) (*ApiResponse, error)
}

