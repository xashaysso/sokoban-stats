package entities

import (
	"time"
)

type LevelStats struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	LevelID   int       `json:"level_id"`
	Steps     int       `json:"steps"`
	Time      int       `json:"time"` // <seconds>
	CreatedAt time.Time `json:"created_at"`
}

type CreateLevelStatsParams struct {
	Username string `json:"username"`
	LevelID  int    `json:"level_id"`
	Steps    int    `json:"steps"`
	Time     int    `json:"time"` // <seconds>
}
