package main

import (
	"context"
	"fmt"
	"os"
	"time"

	tasksrg "github.com/elojah/selection/pkg/task/srg"
	usersrg "github.com/elojah/selection/pkg/user/srg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func run(prog string, filename string) {

	// #Init logger
	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	// #Read config file
	var c Config
	if err := c.Populate(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to read config")
		return
	}

	// #Set initial context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.MongoTimeout)*time.Second)
	defer cancel()

	// #Init mongodb client
	mclient, err := MongoClient(ctx, c.MongoURL)
	if err != nil {
		log.Error().Err(err).Str("url", c.MongoURL).Msg("failed to init mongodb client")
		return
	}

	// #Init data stores
	userStore := &usersrg.Store{}
	if err := userStore.Up(mclient); err != nil {
		log.Error().Err(err).Msg("failed to init mongodb user store")
		return
	}

	taskStore := &tasksrg.Store{}
	if err := taskStore.Up(mclient); err != nil {
		log.Error().Err(err).Msg("failed to init mongodb task store")
		return
	}

	// #Init handler
	h := NewHandler(ctx)
	h.UserStore = userStore
	h.TaskStore = taskStore
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
