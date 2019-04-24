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
