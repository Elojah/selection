package task

import "context"

// TagStore to retrieves task tags by task ID.
type TagStore interface {
	GetAllTags(context.Context) ([]Tags, error)
	GetTagsByID(context.Context, []string) ([]Tags, error)
}
