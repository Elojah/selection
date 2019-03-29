package srg

import (
	"context"

	"github.com/elojah/selection/pkg/task"
	multierror "github.com/hashicorp/go-multierror"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTags implemented with mongodb.
func (s *Store) GetAllTags(ctx context.Context) ([]task.Tags, error) {

	cur, err := s.tags.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var tags []task.Tags
	for cur.Next(ctx) {
		var result mongoTag
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, err)
			continue
		}
		tags = append(tags, result.Domain())
	}

	return tags, merr.ErrorOrNil()
}

// GetTagsByID implemented with mongodb.
func (s *Store) GetTagsByID(ctx context.Context, ids []string) ([]task.Tags, error) {

	a := make(bson.A, len(ids))
	for i, id := range ids {
		a[i] = id
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: a}}}}
	cur, err := s.tags.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var tasks []task.Tags
	for cur.Next(ctx) {
		var result mongoTag
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, err)
			continue
		}
		tasks = append(tasks, result.Domain())
	}

	return tasks, merr.ErrorOrNil()
}

type mongoTag struct {
	ID   string   `json:"_id"`
	Tags []string `json:"tags"`
}

func (t mongoTag) Domain() task.Tags {
	return task.Tags{
		ID:   t.ID,
		Tags: t.Tags,
	}
}
