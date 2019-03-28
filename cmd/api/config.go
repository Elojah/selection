package main

import (
	"encoding/json"
	"os"
)

// Config is the udp server structure config.
type Config struct {
	MongoDBURL   string `json:"mongodb_url"`
	MongoTimeout int    `json:"mongodb_timeout"`

	Address string `json:"address"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
}

// Populate reads file as JSON and populate config.
func (c *Config) Populate(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewDecoder(f).Decode(c)
}
