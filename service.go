package wunderground

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const API_URL = "https://api.wunderground.com/api/"

// Service represents your API.
type Service struct {
	ApiKey string
	client *http.Client
}

// NewService creates a Service using the given, if none is provided
// it uses http.DefaultClient.
func NewService(key string) *Service {
	return &Service{
		ApiKey: key,
		client: http.DefaultClient,
	}
}

func (c *Service) request(path string, query, v interface{}) error {
	qs := fmt.Sprintf("/%s/q/%v.json", path, query)
	resp, err := c.client.Get(API_URL + c.ApiKey + qs)
	if err != nil {
		log.Fatal("whoops in request:", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return err
}
