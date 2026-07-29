package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func RunMigrations(dbURL string) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		slog.Error("failed to open db connection for migrations", slog.Any("err", err))
		os.Exit(1)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("failed to set goose dialect", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("running database migrations...")
	if err := goose.Up(db, "migrations"); err != nil {
		slog.Error("migration up failed", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("database migrations applied successfully")
}

func Serve() *pgxpool.Pool {
	ctx := context.Background()

	DB_URL := os.Getenv("DB_URL")

	var pool *pgxpool.Pool
	var err error

	const max_attempts = 5
	const delay = 4 * time.Second

	for attempt := 1; attempt <= max_attempts; attempt++ {
		pool, err = pgxpool.New(ctx, DB_URL)
		if err == nil {
			err = pool.Ping(ctx)
		}
		if err == nil {
			slog.Info("connected to db successfully")
			break
		}

		slog.Warn("database is not ready, retrying...", slog.Any("err", err), slog.Duration("wait_time", delay))
		if pool != nil {
			pool.Close()
		}

		if attempt == max_attempts {
			slog.Error("couldn't connect to db after max attempts. Exiting...")
			os.Exit(1)
		}

		time.Sleep(delay)
	}

	RunMigrations(DB_URL)

	return pool
}
