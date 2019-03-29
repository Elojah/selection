package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elojah/selection/pkg/user"
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

	idsParam := r.URL.Query().Get("ids")
	logger := log.With().Str("route", "/user").Str("ids", idsParam).Str("method", "GET").Logger()

	var users []user.U

	// #Check if id parameter is set
	if idsParam == "" {
		// #Fetch all users
		var err error
		users, err = h.UserStore.GetAllUsers(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve users")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	} else {
		// #Fetch users by id
		var err error
		ids := strings.Split(idsParam, ",")
		users, err = h.UserStore.GetUsersByID(ctx, ids)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve user")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}

	// #If users is nil, respond with an empty array instead of null
	if users == nil {
		users = []user.U{}
	}

	// #Format and respond user
	raw, err := json.Marshal(users)
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
