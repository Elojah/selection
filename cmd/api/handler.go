package main

import (
	"context"
	"net/http"

	"github.com/elojah/selection/pkg/task"
	"github.com/elojah/selection/pkg/user"

	"github.com/rs/zerolog/log"
)

// Handler handles api routes.
type Handler struct {
	srv *http.Server
	ctx context.Context

	UserStore user.Store
	TaskStore task.Store
}

// NewHandler returns a handler with context.
func NewHandler(ctx context.Context) *Handler {
	return &Handler{ctx: ctx}
}

// Dial starts the api server.
func (h *Handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", h.Users)
	mux.HandleFunc("/task", h.Tasks)
	// mux.HandleFunc("/task/{id}/match", h.Match)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() {
		if err := h.srv.ListenAndServeTLS(c.Cert, c.Key); err != nil {
			log.Error().Err(err).Msg("failed to start server")
		}
	}()
	return nil
}
