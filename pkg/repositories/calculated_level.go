package repositories

import (
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
