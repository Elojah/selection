package srg

import (
	"context"
	"time"

	"github.com/elojah/selection/pkg/task"
	multierror "github.com/hashicorp/go-multierror"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTasks implemented with mongodb.
func (s *Store) GetAllTasks(ctx context.Context) ([]task.T, error) {

	cur, err := s.task.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var tasks []task.T
	for cur.Next(ctx) {
		var result mongoTask
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, err)
			continue
		}
		tasks = append(tasks, result.Domain())
	}

	return tasks, merr.ErrorOrNil()
}

// GetTasksByID implemented with mongodb.
func (s *Store) GetTasksByID(ctx context.Context, ids []string) ([]task.T, error) {

	a := make(bson.A, len(ids))
	for i, id := range ids {
		a[i] = id
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: a}}}}
	cur, err := s.task.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var merr *multierror.Error
	var tasks []task.T
	for cur.Next(ctx) {
		var result mongoTask
		if err := cur.Decode(&result); err != nil {
			merr = multierror.Append(merr, err)
			continue
		}
		tasks = append(tasks, result.Domain())
	}

	return tasks, merr.ErrorOrNil()
}

type mongoInfo struct {
	TranslatedName     string `bson:"translatedName"`
	Description        string `bson:"description"`
	TranslatedCategory string `bson:"translatedCategory"`
}

type mongoTask struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Type struct {
		Category string    `bson:"category"`
		Key      string    `bson:"key"`
		Fr       mongoInfo `bson:"fr"`
		En       mongoInfo `bson:"en"`
	} `bson:"type"`
	Country string `bson:"country"`
	Pricing struct {
		Sider    float64 `bson:"sider"`
		Currency string  `bson:"currency"`
		Side     float64 `bson:"side"`
	}
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Description string    `bson:"description"`
	Applicants  []struct {
		ID      string   `bson:"applicantId"`
		Status  string   `bson:"status"`
		Answers []string `bson:"answers"`
	} `bson:"applicants"`
}

// Domain converts a mongodb task into a domain task.
func (t *mongoTask) Domain() task.T {
	applicants := make([]task.Applicants, len(t.Applicants))
	for i, app := range t.Applicants {
		applicants[i] = task.Applicants{
			ID:      app.ID,
			Status:  app.Status,
			Answers: app.Answers,
		}
	}
	return task.T{
		ID:   t.ID,
		Name: t.Name,
		Type: task.Type{
			Category: t.Type.Category,
			Key:      t.Type.Key,
			Fr: task.Info{
				TranslatedName:     t.Type.Fr.TranslatedName,
				Description:        t.Type.Fr.Description,
				TranslatedCategory: t.Type.Fr.TranslatedCategory,
			},
			En: task.Info{
				TranslatedName:     t.Type.En.TranslatedName,
				Description:        t.Type.En.Description,
				TranslatedCategory: t.Type.En.TranslatedCategory,
			},
		},
		Country: t.Country,
		Pricing: task.Price{
			Side:  t.Pricing.Side,
			Sider: t.Pricing.Sider,
		},
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Description: t.Description,
		Applicants:  applicants,
	}
}
