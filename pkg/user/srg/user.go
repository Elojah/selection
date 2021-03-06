package srg

import (
	"context"
	"time"

	"github.com/elojah/selection/pkg/user"
	multierror "github.com/hashicorp/go-multierror"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers implemented with mongodb.
func (s *Store) GetAllUsers(ctx context.Context) ([]user.U, error) {

	cur, err := s.user.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var users []user.U
	for cur.Next(ctx) {
		var result mongoUser
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, err)
			continue
		}
		users = append(users, result.Domain())
	}

	return users, merr.ErrorOrNil()
}

// GetUsersByID implemented with mongodb.
func (s *Store) GetUsersByID(ctx context.Context, ids []string) ([]user.U, error) {

	a := make(bson.A, len(ids))
	for i, id := range ids {
		a[i] = id
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: a}}}}
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
			merr = multierror.Append(merr, err)
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
func (u mongoUser) Domain() user.U {
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
