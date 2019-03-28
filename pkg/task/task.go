package task

import "context"

// Store to retrieve task by ID.
type Store interface {
	GetTask(context.Context, string) (T, error)
}
