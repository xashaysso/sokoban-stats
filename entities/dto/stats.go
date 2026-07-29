package dto

import "stats/entities"

type CreateLevelStatsRequest struct {
	Username string `json:"username"`
	Steps    int    `json:"steps"`
	Time     int    `json:"time"` // <seconds>
}

func (req *CreateLevelStatsRequest) ToParams(levelID int) *entities.CreateLevelStatsParams {
	return &entities.CreateLevelStatsParams{
		Username: req.Username,
		LevelID:  levelID,
		Steps:    req.Steps,
		Time:     req.Time,
	}
}
