package repositories

import (
	"context"
	"stats/entities"
)

type LevelStatsRepo interface {
	GetLevelStats(ctx context.Context, levelID int) ([]entities.LevelStats, error)
	CreateLevelStatsEntry(ctx context.Context, params *entities.CreateLevelStatsParams) (entities.LevelStats, error)
	GetPlayerRecords(ctx context.Context, username string) ([]entities.PlayerRecord, error)
}
