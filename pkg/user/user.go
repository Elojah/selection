package user

// Store to retrieve users by all/batch/ID.
type Store interface {
	GetUserAll() ([]U, error)
	GetUser(id string) (U, error)
}
