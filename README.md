# Sokoban Stats Service

A lightweight Go REST API service for receiving gameplay statistics and managing leaderboards for the Sokoban client.

**Client Repository:** [sokoban-client](https://github.com/xashaysso/sokoban-cw)

## Tech Stack
* **Language:** Go
* **Database:** PostgreSQL (`pgxpool`, `goose`)
* **Deployment:** Docker / Docker Compose

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/levels/:level_id` | Record level completion to leaderboard |
| `GET` | `/levels/:level_id` | Get level leaderboard|
| `GET` | `/player/:username/records` | Get player level records |

## Quick Start

```bash
git clone https://github.com/xashaysso/sokoban-stats
cd sokoban-stats
docker-compose up -d --build
