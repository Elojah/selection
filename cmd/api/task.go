package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elojah/selection/pkg/task"
	"github.com/rs/zerolog/log"
)

// Tasks list all tasks for route /tasks.
func (h *Handler) Tasks(w http.ResponseWriter, r *http.Request) {
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
	logger := log.With().Str("route", "/task").Str("ids", idsParam).Str("method", "GET").Logger()

	var tasks []task.T
	var tags []task.Tags

	// #Check id parameter is not empty
	if idsParam == "" {
		// #Fetch all tasks
		var err error
		tasks, err = h.TaskStore.GetAllTasks(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve tasks")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}

		// #Fetch associated tags
		tags, err = h.TaskTagStore.GetAllTags(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve task tags")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	} else {
		// #Fetch tasks by id
		var err error
		ids := strings.Split(idsParam, ",")
		tasks, err = h.TaskStore.GetTasksByID(ctx, ids)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve task")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}

		// #Fetch associated tags
		tags, err = h.TaskTagStore.GetTagsByID(ctx, ids)
		if err != nil {
			logger.Error().Err(err).Msg("failed to retrieve task tags")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}

	dto := linkTaskTags(tasks, tags)

	// #Format and respond task
	raw, err := json.Marshal(dto)
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

func linkTaskTags(tasks []task.T, tags []task.Tags) []task.DTO {

	// #Convert tags into map to ease next access
	tagsMap := make(map[string]task.Tags, len(tags))
	for _, t := range tags {
		tagsMap[t.ID] = t
	}

	dto := make([]task.DTO, len(tasks))
	for i, t := range tasks {
		taskTags, ok := tagsMap[t.ID]
		if !ok {
			// no tags found for this task, raise an error ?
			taskTags.Tags = []string{}
		}
		dto[i] = task.DTO{
			Task: t,
			Tags: taskTags.Tags,
		}
	}

	return dto
}
