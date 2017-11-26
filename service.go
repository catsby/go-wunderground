package wunderground

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	API_URL = "https://api.wunderground.com/api/"
	GIVE_ATTRIBUTION = true
)

// A very minimal abstraction around http.Client.Get that acts as a service fetch
type WUFetch interface {
	Get(url string) (resp *http.Response, err error)
}

// Service represents your API.
type Service struct {
	ApiKey string
	client WUFetch
}

// NewService creates a Service using the given, if none is provided
// it uses http.DefaultClient.
func NewService(key string) *Service {
	if GIVE_ATTRIBUTION {
		fmt.Print("Data provided by The Weather Underground (wunderground.com)")
		GIVE_ATTRIBUTION = false
	}

	return &Service{
		ApiKey: key,
		client: http.DefaultClient,
	}
}

func (c *Service) request(path string, query *Query) (*ApiResponse, error) {
	qs := fmt.Sprintf("/%s/q/%s", path, query)
	resp, err := c.client.Get(API_URL + c.ApiKey + qs)
	if err != nil {
		log.Fatal("whoops in request:", err)
	}

	defer resp.Body.Close()

	ar := &ApiResponse{}
	err = json.NewDecoder(resp.Body).Decode(ar)
	return ar, err
}
