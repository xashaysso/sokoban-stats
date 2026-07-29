package services

import (
	"context"
	"stats/entities"
)

type StatsServiceInterface interface {
	GetLevelStats(ctx context.Context, levelID int) ([]entities.LevelStats, error)
	CreateLevelStatsEntry(ctx context.Context, params *entities.CreateLevelStatsParams) (entities.LevelStats, error)
}
