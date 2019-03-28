package main

import (
	"encoding/json"
	"net/http"

	merrors "github.com/elojah/selection/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Users list all users for route /users.
func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := h.ctx()
	defer cancel()

	id := r.URL.Query().Get("id")
	logger := log.With().Str("route", "/user").Str("id", id).Str("method", "GET").Logger()

	// #Check id parameter is not empty
	if id == "" {
		logger.Error().Err(merrors.ErrMissingParam{Name: "id"}).Msg("invalid parameter")
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}

	// #Retrieve user by id
	u, err := h.UserStore.GetUser(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve user")
		http.Error(w, "store failure", http.StatusInternalServerError)
		return
	}

	// #Format and respond user
	raw, err := json.Marshal(u)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal response")
		http.Error(w, "formatting failure", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		http.Error(w, "writing failure", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	logger.Info().Msg("success")
}
