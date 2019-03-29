package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/elojah/selection/pkg/task"
	"github.com/elojah/selection/pkg/user"
)

// Handler handles api routes.
type Handler struct {
	srv *http.Server
	ctx func() (context.Context, context.CancelFunc)

	client *grpc.ClientConn

	UserStore    user.Store
	TaskStore    task.Store
	TaskTagStore task.TagStore
	TaskScorer   task.ScorerClient
}

// Up starts the api server.
func (h *Handler) Up(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", h.Users)
	mux.HandleFunc("/task", h.Tasks)
	mux.HandleFunc("/task/scores", h.Scores)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() {
		if err := h.srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("failed to init server")
		}
	}()
	log.Info().Msg("server listening")

	var err error
	h.client, err = grpc.Dial(c.Scorer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	h.TaskScorer = task.NewScorerClient(h.client)

	return nil
}

// Down closes handler open connections.
func (h *Handler) Down() error {
	return h.client.Close()
}
