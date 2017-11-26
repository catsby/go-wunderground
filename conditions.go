package wunderground

import (
  "log"
)

type ConditionsResponse struct {
	Response            Response            `json:"response"`
	CurrentObservastion CurrentObservastion `json:"current_observation"`
}

type CurrentObservastion struct {
	DisplayLocation struct {
		Full string `json:"full"`
	} `json:"display_location"`
}

func (co *ConditionsResponse) LocationName() string {
	return co.CurrentObservastion.DisplayLocation.Full
}

func (c *Service) Conditions(query interface{}) (*ConditionsResponse, error) {
	cr := &ConditionsResponse{}
	err := c.request("conditions", query, cr)

	if err != nil {
		log.Fatal("something wrong in the request: ", err)
	}

	return cr, err
}

