-- +goose Up
CREATE TABLE IF NOT EXISTS level_stats (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    level_id INT NOT NULL,
    steps INT NOT NULL,
    time INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS level_stats;
