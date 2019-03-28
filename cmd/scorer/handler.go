package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/elojah/selection/pkg/task"
	"github.com/elojah/selection/pkg/user"
)

// Handler handles api routes.
type Handler struct {
	ctx context.Context

	UserStore    user.Store
	TaskStore    task.Store
	TaskTagStore task.TagStore
}

// NewHandler returns a handler with context.
func NewHandler(ctx context.Context) *Handler {
	return &Handler{ctx: ctx}
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
