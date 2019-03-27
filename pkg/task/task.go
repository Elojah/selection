package task

// Store to retrieve task by ID.
type Store interface {
	GetTask(id string) (T, error)
}
