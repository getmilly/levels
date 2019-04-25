package repositories

import (
	"context"

	"github.com/getmilly/levels/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//ExperienceHistory ...
type ExperienceHistory interface {
	Add(experience *models.Experience) error
	Sum(playerID string) (*models.SumExperienceResult, error)
}

type experienceHistory struct {
	collection *mongo.Collection
}

//NewExperienceHistory ...
func NewExperienceHistory(client *mongo.Client) ExperienceHistory {
	return &experienceHistory{
		collection: client.Database("levels").Collection("experience_history"),
	}
}

func (repository *experienceHistory) Add(experience *models.Experience) error {
	_, err := repository.collection.InsertOne(context.Background(), experience)

	return err
}

func (repository *experienceHistory) Sum(playerID string) (*models.SumExperienceResult, error) {
	result := new(models.SumExperienceResult)
	pipeline := make([]bson.M, 2)
	sum := bson.M{
		"$group": bson.M{
			"_id": "$player_id",
			"total_experience": bson.M{
				"$sum": "$experience",
			},
		},
	}

	match := bson.M{
		"$match": bson.M{
			"_id": playerID,
		},
	}

	pipeline[0] = sum
	pipeline[1] = match

	cursor, err := repository.collection.Aggregate(context.Background(), pipeline)

	if err != nil {
		return nil, err
	}

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
