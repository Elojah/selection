package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	var c Config
	if err := c.Populate(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to read config")
		return
	}

	ctx, f := context.WithTimeout(context.Background(), time.Duration(c.MongoTimeout)*time.Second)
	if f != nil {
		f()
		return
	}

	mclient, err := MongoClient(ctx, c.MongoDBURL)
	if err != nil {
		log.Error().Err(err).Str("url", c.MongoDBURL).Msg("failed to start mongodb")
		return
	}

	_ = mclient
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
