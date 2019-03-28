package srg

import (
	"context"
	"time"

	merrors "github.com/elojah/selection/pkg/errors"
	"github.com/elojah/selection/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	layout = "2006-01-02T15:04:05.000Z"
)

// GetUserAll implemented with mongodb.
func (s *Store) GetUserAll(ctx context.Context) ([]user.U, error) {
	return nil, merrors.ErrNotImplemented{}
}

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
	ID        string `json:"_id"`
	CreatedAt struct {
		Date string `json:"$date"`
	} `json:"createdAt"`
	UpdatedAt struct {
		Date string `json:"$date"`
	} `json:"updatedAt"`
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	Tags             []string `json:"tags"`
	TaskApplications []string `json:"taskApplications"`
}

// Domain converts a mongodb user into a domain user.
func (u *mongoUser) Domain() (user.U, error) {
	cat, err := time.Parse(layout, u.CreatedAt.Date)
	if err != nil {
		return user.U{}, err
	}

	uat, err := time.Parse(layout, u.UpdatedAt.Date)
	if err != nil {
		return user.U{}, err
	}

	return user.U{
		ID:               u.ID,
		CreatedAt:        cat,
		UpdatedAt:        uat,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Tags:             u.Tags,
		TaskApplications: u.TaskApplications,
	}, nil
}
