package srg

import (
	"github.com/elojah/selection/pkg/task"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ task.Store = (*Store)(nil)
var _ task.TagStore = (*Store)(nil)

// Store implements task domain stores with mongodb.
type Store struct {
	*mongo.Client

	task *mongo.Collection
	tags *mongo.Collection
}

func (s *Store) Up(client *mongo.Client) error {
	s.Client = client

	s.task = s.Client.Database("main").Collection("tasks")
	s.tags = s.Client.Database("tags").Collection("tasksTags")
	return nil
}
