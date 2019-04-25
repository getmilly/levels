package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//CalculatedLevel ...
type CalculatedLevel struct {
	ID                  primitive.ObjectID `json:"_" bson:"_id,omitempty"`
	PlayerID            string             `json:"player_id" bson:"player_id,omitempty"`
	TotalExperience     int                `json:"total_experience" bson:"total_experience,omitempty"`
	Level               int                `json:"level" bson:"level,omitempty"`
	ExperienceToUpgrade int                `json:"experience_to_upgrade" bson:"experience_to_upgrade,omitempty"`
}
