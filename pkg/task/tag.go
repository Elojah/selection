package task

// Tags represents all tags for a task.
type Tags []string

// TagStore to retrieves task tags by task ID.
type TagStore interface {
	GetTags(id string) (Tags, error)
}
