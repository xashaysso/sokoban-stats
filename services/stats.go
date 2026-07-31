package services

import (
	"context"
	"stats/db/repositories"
	"stats/entities"
)

type StatsService struct {
	StatsRepo repositories.LevelStatsRepo
}

func NewStatsService(statsRepo repositories.LevelStatsRepo) *StatsService {
	return &StatsService{
		StatsRepo: statsRepo,
	}
}

func (s *StatsService) GetLevelStats(ctx context.Context, levelID int) ([]entities.LevelStats, error) {
	if levelID < 0 {
		return nil, ErrInvalidLevelID
	}
	levelStats, err := s.StatsRepo.GetLevelStats(ctx, levelID)
	if err != nil {
		return nil, err
	}
	return levelStats, nil
}

func (s *StatsService) CreateLevelStatsEntry(ctx context.Context, params *entities.CreateLevelStatsParams) (entities.LevelStats, error) {
	if params.LevelID < 0 {
		return entities.LevelStats{}, ErrInvalidLevelID
	}
	if len(params.Username) < 3 {
		return entities.LevelStats{}, ErrInvalidUsername
	}
	if params.Steps <= 0 {
		return entities.LevelStats{}, ErrInvalidSteps
	}
	if params.Time <= 0 {
		return entities.LevelStats{}, ErrInvalidTime
	}
	levelStats, err := s.StatsRepo.CreateLevelStatsEntry(ctx, params)
	if err != nil {
		return entities.LevelStats{}, err
	}
	return levelStats, nil
}
