package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/elojah/selection/pkg/user/srg"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.MongoTimeout)*time.Second)
	defer cancel()

	mclient, err := MongoClient(ctx, c.MongoDBURL)
	if err != nil {
		log.Error().Err(err).Str("url", c.MongoDBURL).Msg("failed to init mongodb client")
		return
	}

	userStore := &srg.Store{}
	if err := userStore.Up(mclient); err != nil {
		log.Error().Err(err).Msg("failed to init mongodb user store")
		return
	}

	h := NewHandler(ctx)
	h.UserStore = userStore
	if err := h.Dial(c); err != nil {
		log.Error().Err(err).Msg("failed to init http handler")
		return
	}

	// TODO handle SIG*
	select {}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
