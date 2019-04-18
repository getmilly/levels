package messages

import "time"

//RewardReached ...
type RewardReached struct {
	Kind      string            `json:"kind"`
	MissionID string            `json:"mission_id"`
	TaskID    string            `json:"task_id"`
	GameID    string            `json:"game_id"`
	CreatedAt time.Time         `json:"created_at"`
	Reason    string            `json:"reason"`
	Data      *RewardExperience `json:"data"`
}

//RewardExperience ...
type RewardExperience struct {
	Experience int `json:"experience"`
}
