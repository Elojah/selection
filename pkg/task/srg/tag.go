package srg

import (
	"context"

	"github.com/elojah/selection/pkg/task"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTags implemented with mongodb.
func (s *Store) GetTags(ctx context.Context, id string) (task.Tags, error) {
	var result task.Tags

	filter := bson.M{"_id": id}
	if err := s.tags.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
