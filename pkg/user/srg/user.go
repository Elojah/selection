package srg

import (
	"context"
	"time"

	"github.com/elojah/selection/pkg/errors"
	"github.com/elojah/selection/pkg/user"
	multierror "github.com/hashicorp/go-multierror"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUser implemented with mongodb.
func (s *Store) GetUser(ctx context.Context, id string) (user.U, error) {
	var result mongoUser

	filter := bson.M{"_id": id}
	if err := s.user.FindOne(ctx, filter).Decode(&result); err != nil {
		return user.U{}, err
	}
	return result.Domain(), nil
}

// GetUsers implemented with mongodb.
func (s *Store) GetUsers(ctx context.Context, ids []string) ([]user.U, error) {

	filter := make(bson.D, len(ids))
	for i, id := range ids {
		filter[i] = bson.E{Key: "_id", Value: id}
	}

	cur, err := s.user.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var users []user.U
	for cur.Next(ctx) {
		var result mongoUser
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, errors.ErrNotFound{Collection: "users", Index: "unknown"})
			continue
		}
		users = append(users, result.Domain())
	}

	return users, merr.ErrorOrNil()
}

type mongoUser struct {
	ID               string    `bson:"_id"`
	CreatedAt        time.Time `bson:"createdAt"`
	UpdatedAt        time.Time `bson:"updatedAt"`
	FirstName        string    `bson:"firstName"`
	LastName         string    `bson:"lastName"`
	Tags             []string  `bson:"tags"`
	TaskApplications []string  `bson:"taskApplications"`
}

// Domain converts a mongodb user into a domain user.
func (u *mongoUser) Domain() user.U {
	return user.U{
		ID:               u.ID,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Tags:             u.Tags,
		TaskApplications: u.TaskApplications,
	}
}
