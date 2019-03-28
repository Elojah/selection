package user

import "context"

// Store to retrieve users by all/batch/ID.
type Store interface {
	GetUser(context.Context, string) (U, error)
}
