package srg

import (
	"context"

	"github.com/elojah/selection/pkg/user"
)

// GetUserAll implemented with mongodb.
func (s *Store) GetUserAll(ctx context.Context) ([]user.U, error) {
	return nil, nil
}

// GetUser implemented with mongodb.
func (s *Store) GetUser(ctx context.Context, id string) (user.U, error) {
	return user.U{}, nil
}
