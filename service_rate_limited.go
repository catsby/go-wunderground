package wunderground

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"time"
)

var _ WUFetch = &RateLimitedClient{}

var logger = log.New(os.Stderr, "RateLimitedClient", log.LstdFlags)

// Provides a rate-limited WUFetch
type RateLimitedClient struct {
	client WUFetch
	lim    *rate.Limiter
}

func (c *RateLimitedClient) Get(url string) (resp *http.Response, err error) {
	r := c.lim.Reserve()
	if !r.OK() {
		// Not allowed to act! Did you remember to set lim.burst to be > 0 ?
		return nil, fmt.Errorf("request rejected by limiter")
	}
	if r.Delay() > (time.Duration(2*(1/c.lim.Limit())) * time.Second) {
		logger.Printf("INFO: Request delayed %f", r.Delay())
	}
	time.Sleep(r.Delay())

	return c.client.Get(url)
}

// Convenience function to create a service for a Developer API key
func NewDevLimitedService(key string) *Service {
	return NewRateLimitedService(key, 10.0/60)
}

// Create a rate-limited service object. limit defines the rate-per-second.
// See https://www.wunderground.com/weather/api/d/pricing.html
func NewRateLimitedService(key string, limit float64) *Service {
	service := NewService(key)
	service.client = &RateLimitedClient{
		client: service.client,
		// I don't know the accuracy of the API limit so choose a conservative burst.
		lim: rate.NewLimiter(rate.Limit(limit), 3 /* burst */),
	}
	return service
}
