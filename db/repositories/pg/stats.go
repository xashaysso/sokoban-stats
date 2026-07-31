package pg

import (
	"context"
	"stats/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgStatsRepo struct {
	Pool *pgxpool.Pool
}

func NewStatsRepo(p *pgxpool.Pool) *PgStatsRepo {
	return &PgStatsRepo{
		Pool: p,
	}
}

func (r *PgStatsRepo) GetLevelStats(ctx context.Context, levelID int) ([]entities.LevelStats, error) {
	var levelStatsList []entities.LevelStats
	rows, err := r.Pool.Query(ctx, `SELECT id, username, level_id, steps, time, created_at FROM level_stats WHERE level_id = $1 ORDER BY steps, time ASC LIMIT 10`, levelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var levelStats entities.LevelStats
		err := rows.Scan(&levelStats.ID, &levelStats.Username, &levelStats.LevelID, &levelStats.Steps, &levelStats.Time, &levelStats.CreatedAt)
		if err != nil {
			return nil, err
		}

		levelStatsList = append(levelStatsList, levelStats)
	}

	return levelStatsList, nil
}

func (r *PgStatsRepo) CreateLevelStatsEntry(ctx context.Context, params *entities.CreateLevelStatsParams) (entities.LevelStats, error) {
	var StatsEntry entities.LevelStats
	row := r.Pool.QueryRow(ctx, `INSERT INTO level_stats (username, level_id, steps, time) VALUES ($1, $2, $3, $4) RETURNING id, username, level_id, steps, time, created_at`, params.Username, params.LevelID, params.Steps, params.Time)
	if err := row.Scan(&StatsEntry.ID, &StatsEntry.Username, &StatsEntry.LevelID, &StatsEntry.Steps, &StatsEntry.Time, &StatsEntry.CreatedAt); err != nil {
		return entities.LevelStats{}, err
	}

	return StatsEntry, nil
}

func (r *PgStatsRepo) GetPlayerRecords(ctx context.Context, username string) ([]entities.PlayerRecord, error) {
	var PlayerRecordsList []entities.PlayerRecord
	rows, err := r.Pool.Query(ctx, `SELECT DISTINCT ON (level_id) level_id, steps, time FROM level_stats WHERE username = $1 ORDER BY level_id, steps ASC, time ASC`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var playerRecord entities.PlayerRecord
		err := rows.Scan(&playerRecord.LevelID, &playerRecord.Steps, &playerRecord.Time)
		if err != nil {
			return nil, err
		}

		PlayerRecordsList = append(PlayerRecordsList, playerRecord)
	}

	return PlayerRecordsList, nil
}
