package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/getmilly/levels/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//CalculatedLevelRepository ...
type CalculatedLevelRepository interface {
	Set(level *models.CalculatedLevel) error
	Get(playerID string) (*models.CalculatedLevel, error)
}

type calculatedLevelRepository struct {
	collection *mongo.Collection
}

//NewCalculatedLevelRepository ...
func NewCalculatedLevelRepository(client *mongo.Client) CalculatedLevelRepository {
	collection := client.Database("levels").Collection("calculated_level")
	return &calculatedLevelRepository{
		collection: collection,
	}
}

func (repository *calculatedLevelRepository) Set(level *models.CalculatedLevel) error {
	ctx := context.Background()

	result := repository.collection.FindOneAndReplace(ctx, bson.M{
		"player_id": level.PlayerID,
	}, level)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			_, err := repository.collection.InsertOne(ctx, level)
			return err
		}

		return result.Err()
	}

	return nil
}

func (repository *calculatedLevelRepository) Get(playerID string) (*models.CalculatedLevel, error) {
	ctx := context.Background()

	level := new(models.CalculatedLevel)

	result := repository.collection.FindOne(ctx, bson.M{
		"player_id": playerID,
	})

	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(level)

	if err == mongo.ErrNoDocuments {
		return level, nil
	}

	return level, err
}
