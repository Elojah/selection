package user

import "context"

// Store to retrieve users by all/batch/ID.
type Store interface {
	GetAllUsers(context.Context) ([]U, error)
	GetUsersByID(context.Context, []string) ([]U, error)
}
