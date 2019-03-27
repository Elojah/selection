package user

import "context"

// Store to retrieve users by all/batch/ID.
type Store interface {
	GetUserAll(context.Context) ([]U, error)
	GetUser(context.Context, string) (U, error)
}
