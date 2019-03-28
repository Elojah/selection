package main

import (
	"encoding/json"
	"os"
)

// Config is the udp server structure config.
type Config struct {
	MongoURL     string `json:"mongo_url"`
	MongoTimeout int    `json:"mongo_timeout"`

	Address string `json:"address"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`

	Scorer string `json:"scorer"`
}

// Populate reads file as JSON and populate config.
func (c *Config) Populate(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewDecoder(f).Decode(c)
}
