package srg

import (
	"github.com/elojah/selection/pkg/task"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ task.Store = (*Store)(nil)

// Store implements task domain stores with mongodb.
type Store struct {
	*mongo.Client

	task *mongo.Collection
}

func (s *Store) Up(client *mongo.Client) error {
	s.Client = client

	s.task = s.Client.Database("main").Collection("tasks")
	return nil
}
