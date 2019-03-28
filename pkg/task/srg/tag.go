package srg

import (
	"context"

	"github.com/elojah/selection/pkg/task"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTags implemented with mongodb.
func (s *Store) GetTags(ctx context.Context, id string) (task.Tags, error) {
	var result mongoTag

	filter := bson.M{"_id": id}
	if err := s.tags.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return task.Tags(result.Tags), nil
}

type mongoTag struct {
	ID   string   `json:"_id"`
	Tags []string `json:"tags"`
}
