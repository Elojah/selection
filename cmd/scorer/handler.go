package main

import (
	"net"

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
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
