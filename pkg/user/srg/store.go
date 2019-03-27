package srg

import (
	"github.com/elojah/selection/pkg/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ user.Store = (*Store)(nil)

// Store implements user domain stores with mongodb.
type Store struct {
	*mongo.Client

	user *mongo.Collection
}

func (s *Store) Up(client *mongo.Client) error {
	s.Client = client

	s.user = s.Client.Database("main").Collection("users")
	return nil
}
