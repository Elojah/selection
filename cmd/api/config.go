package main

import (
	"encoding/json"
	"os"
)

// Config is the udp server structure config.
type Config struct {
	MongoTimeout int    `json:"mongo_timeout"`
	MongoDBURL   string `json:"mongodb_url"`
}

// Populate reads file as JSON and populate config.
func (c *Config) Populate(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewDecoder(f).Decode(c)
}
