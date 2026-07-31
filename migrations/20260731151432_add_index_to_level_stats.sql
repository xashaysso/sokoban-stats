-- +goose Up
CREATE INDEX IF NOT EXISTS idx_stats_user_level_score ON level_stats(username, level_id, steps, time);

-- +goose Down
DROP INDEX IF EXISTS idx_stats_user_level_score;
