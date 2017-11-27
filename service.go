package wunderground

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ApiURL          = "https://api.wunderground.com/api/"
	GiveAttribution = true
	LogErrors       = true
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
	if GiveAttribution {
		fmt.Println("Data provided by The Weather Underground (wunderground.com)")
		GiveAttribution = false
	}

	return &Service{
		ApiKey: key,
		client: http.DefaultClient,
	}
}

func (c *Service) Request(features []string, query *Query) (*ApiResponse, error) {
	if len(features) == 0 {
		return nil, fmt.Errorf("must request at least one feature")
	}

	return c.request(strings.Join(features, "/"), query)
}

func (c *Service) request(path string, query *Query) (*ApiResponse, error) {
	qs := fmt.Sprintf("/%s/q/%s", path, query)
	rawurl := ApiURL + c.ApiKey + qs
	logDebug(rawurl)
	resp, err := c.client.Get(rawurl)
	if err != nil {
		return nil, fmt.Errorf("whoops in request: ", err)
	}

	var body io.Reader = resp.Body
	defer resp.Body.Close()

	var debugOut bytes.Buffer
	if *verbosity > 1 {
		body = io.TeeReader(body, &debugOut)
		defer logDebug(&debugOut)
	}

	ar := &ApiResponse{}
	err = json.NewDecoder(body).Decode(ar)
	if err != nil {
		return nil, fmt.Errorf("error decoding: ", err)
	}

	if LogErrors && ar.Response.Error != nil {
		logger.Printf("ERROR: API response: %s (%s)", ar.Response.Error.Description, ar.Response.Error.Type)
	}

	return ar, nil
}
