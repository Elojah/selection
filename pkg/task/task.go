package task

import "context"

// Store to retrieve task by ID.
type Store interface {
	GetAllTasks(context.Context) ([]T, error)
	GetTasksByID(context.Context, []string) ([]T, error)
}
