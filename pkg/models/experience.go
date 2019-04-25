package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Experience ...
type Experience struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GameID     string             `json:"game_id" bson:"game_id,omitempty"`
	PlayerID   string             `json:"player_id" bson:"player_id,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	Experience int                `json:"experience" bson:"experience,omitempty"`
	Reason     string             `json:"reason" bson:"reason,omitempty"`
	UserID     string             `json:"-" bson:"user_id,omitempty"`
}

//SumExperienceResult ...
type SumExperienceResult struct {
	TotalExperience int `bson:"total_experience,omitempty"`
}
