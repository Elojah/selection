package main

import (
	"encoding/json"
	"net/http"
	"sort"

	merrors "github.com/elojah/selection/pkg/errors"
	"github.com/elojah/selection/pkg/task"
	"github.com/rs/zerolog/log"
)

// Scores calculates scores for a task.
func (h *Handler) Scores(w http.ResponseWriter, r *http.Request) {
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
	logger := log.With().Str("route", "/task/scores").Str("id", id).Str("method", "GET").Logger()

	// #Check id parameter is not empty
	if id == "" {
		logger.Error().Err(merrors.ErrMissingParam{Name: "id"}).Msg("invalid parameter")
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}

	// #Fetch task by id
	t, err := h.TaskStore.GetTask(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve task")
		http.Error(w, "store failure", http.StatusInternalServerError)
		return
	}

	// #Fetch associated tags
	tags, err := h.TaskTagStore.GetTags(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve tags")
		http.Error(w, "store failure", http.StatusInternalServerError)
		return
	}

	// #Call scorer service to retrieve matches
	resp, err := h.TaskScorer.Calculate(ctx, &task.ScorerRequest{TaskID: id})
	if err != nil {
		logger.Error().Err(err).Msg("failed to calculate scores")
		http.Error(w, "calculation failure", http.StatusInternalServerError)
		return
	}
	scores := task.Scores{
		TaskID:      id,
		Applicants:  resp.Scores,
		Description: t.Description,
		Country:     t.Country,
		Tags:        tags,
	}

	// #Sort applicants by score
	sort.Slice(scores.Applicants, func(i, j int) bool {
		return scores.Applicants[i].Score > scores.Applicants[j].Score
	})

	// #Format and respond task
	raw, err := json.Marshal(scores)
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
