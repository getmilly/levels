package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Level ...
type Level struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GameID     string             `json:"game_id" bson:"game_id"`
	PlayerID   string             `json:"player_id" bson:"player_id"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	Experience int                `json:"experience" bson:"experience"`
	Reason     string             `json:"reason" bson:"reason"`
	UserID     string             `json:"-" bson:"user_id"`
}

//CalculatedLevel ...
type CalculatedLevel struct {
	ID                  primitive.ObjectID `json:"_" bson:"_id,omitempty"`
	PlayerID            string             `json:"player_id" bson:"player_id"`
	TotalExperience     int                `json:"total_experience" bson:"total_experience"`
	Level               int                `json:"level" bson:"level"`
	NextLevelExperience int                `json:"next_level_experience" bson:"next_level_experience"`
}
