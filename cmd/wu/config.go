package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User-persistent configuration
type Config struct {
	Key     string
	Station string
	Degrees string
}

func ParseConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open config (%s): %s", file, err)
	}

	c := &Config{}
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %s", err)
	}

	return c, nil
}
