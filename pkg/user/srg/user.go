package srg

import (
	"context"
	"time"

	"github.com/elojah/selection/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUser implemented with mongodb.
func (s *Store) GetUser(ctx context.Context, id string) (user.U, error) {
	var result mongoUser

	filter := bson.M{"_id": id}
	if err := s.user.FindOne(ctx, filter).Decode(&result); err != nil {
		return user.U{}, err
	}
	return result.Domain()
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
func (u *mongoUser) Domain() (user.U, error) {
	return user.U{
		ID:               u.ID,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Tags:             u.Tags,
		TaskApplications: u.TaskApplications,
	}, nil
}
