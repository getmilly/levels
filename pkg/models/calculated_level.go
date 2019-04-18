package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//CalculatedLevel ...
type CalculatedLevel struct {
	ID                  primitive.ObjectID `json:"_" bson:"_id,omitempty"`
	PlayerID            string             `json:"player_id" bson:"player_id"`
	TotalExperience     int                `json:"total_experience" bson:"total_experience"`
	Level               int                `json:"level" bson:"level"`
	NextLevelExperience int                `json:"next_level_experience" bson:"next_level_experience"`
}
