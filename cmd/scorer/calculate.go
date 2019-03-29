package main

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/elojah/selection/pkg/errors"
	"github.com/elojah/selection/pkg/task"
	"github.com/elojah/selection/pkg/user"
)

// Calculate handle score calculation route.
func (h *Handler) Calculate(ctx context.Context, r *task.ScorerRequest) (*task.ScorerReply, error) {

	id := r.TaskID

	// #Fetch associated task
	ts, err := h.TaskStore.GetTasksByID(ctx, []string{id})
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, errors.ErrNotFound{Collection: "task", Index: id}
	}
	t := ts[0]

	// #Fetch associated tags
	tags, err := h.TaskTagStore.GetTagsByID(ctx, []string{id})
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, errors.ErrNotFound{Collection: "tags", Index: id}
	}
	taskTags := tags[0]

	// #Fetch all applicants users
	ids := make([]string, len(t.Applicants))
	for i, app := range t.Applicants {
		ids[i] = app.ID
	}
	users, err := h.UserStore.GetUsersByID(ctx, ids)
	if err != nil {
		return nil, err
	}
	// convert users into map to ease access
	usersMap := make(map[string]user.U)
	for _, u := range users {
		usersMap[u.ID] = u
	}

	// #Calculate match percentage with all applicants
	var result *multierror.Error
	reply := task.ScorerReply{
		TaskID: id,
		Scores: make([]task.Score, len(t.Applicants)),
	}
	for i, applicant := range t.Applicants {
		u, ok := usersMap[applicant.ID]
		if !ok {
			result = multierror.Append(result, errors.ErrNotFound{Collection: "users", Index: applicant.ID})
			continue
		}
		reply.Scores[i] = task.Score{
			SiderID:   applicant.ID,
			Score:     matchTags(u.Tags, taskTags.Tags),
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	}

	return &reply, result.ErrorOrNil()
}

func matchTags(userTags []string, taskTags []string) float64 {

	// #Transform user tags into map
	tagsMap := make(map[string]struct{}, len(userTags))
	for _, tag := range userTags {
		tagsMap[tag] = struct{}{}
	}

	// #Count number of matches
	var count int
	for _, tag := range taskTags {
		if _, ok := tagsMap[tag]; ok {
			count++
		}
	}

	// Return count of matches / total tags of task
	return float64(count) / float64(len(taskTags))
}
