package srg

import (
	"context"
	"time"

	"github.com/elojah/selection/pkg/task"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTask implemented with mongodb.
func (s *Store) GetTask(ctx context.Context, id string) (task.T, error) {
	var result mongoTask

	filter := bson.M{"_id": id}
	if err := s.task.FindOne(ctx, filter).Decode(&result); err != nil {
		return task.T{}, err
	}
	return result.Domain()
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
func (t *mongoTask) Domain() (task.T, error) {
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
	}, nil
}
