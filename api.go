package wunderground

type Response struct {
	Version string `json:"version,omitempty"`
	TOS     string `json:"termsOfService,omitempty"`
}
