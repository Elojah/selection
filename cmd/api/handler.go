package main

import (
	"net/http"

	"github.com/coreos/fleet/log"
	"github.com/elojah/selection/pkg/user"
)

// Handler handles api routes.
type Handler struct {
	user.Store
}

// Dial starts the auth server.
func (h *Handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", h.Users)

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
