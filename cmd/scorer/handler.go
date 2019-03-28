package main

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/elojah/selection/pkg/task"
	"github.com/elojah/selection/pkg/user"
)

// Handler handles api routes.
type Handler struct {
	UserStore    user.Store
	TaskStore    task.Store
	TaskTagStore task.TagStore
}

// Up starts the score server.
func (h *Handler) Up(c Config) error {
	lis, err := net.Listen("tcp", c.Address)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	task.RegisterScorerServer(s, h)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Error().Err(err).Msg("failed to init server")
		}
	}()
	log.Info().Msg("server listening")
	return nil
}
