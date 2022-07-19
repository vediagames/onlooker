package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	domain "github.com/vediagames/onlooker/domain/level"
	"github.com/vediagames/onlooker/errutil"
)

type store struct {
	db *sqlx.DB
}

type Config struct {
	ConnectionString string
}

func (c Config) Validate() error {
	var err errutil.Error

	if c.ConnectionString == "" {
		err.Add(fmt.Errorf("connection string is empty"))
	}

	return err.Err()
}

func New(cfg Config) (domain.Store, error) {
	if ve := cfg.Validate(); ve != nil {
		return nil, fmt.Errorf("invalid config: %w", ve)
	}

	db, err := sqlx.Open("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &store{
		db: db,
	}, nil
}

type insertResult struct {
	UUID       string    `db:"uuid"`
	ServerTime time.Time `db:"server_time"`
}

func (s store) Insert(ctx context.Context, q domain.InsertQuery) (domain.InsertResult, error) {
	var res insertResult

	metadata, err := json.Marshal(q.Metadata)
	if err != nil {
		return domain.InsertResult{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	err = s.db.Get(&res, `
		INSERT INTO levels (session_uuid, client_time, server_time, level, metadata) 
		VALUES ($1, $2, now(), $3, $4)
		RETURNING uuid, server_time
	`, q.SessionUUID, q.ClientTime, q.Level, metadata)
	if err != nil {
		return domain.InsertResult{}, fmt.Errorf("failed to insert level: %v", err)
	}

	return domain.InsertResult{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	}, nil
}

var eventTableMap = map[domain.Event]string{
	domain.EventComplete:           "level_complete_events",
	domain.EventDeath:              "level_death_events",
	domain.EventGrapplingHookUsage: "level_grappling_hook_events",
}

func (s store) InsertEvent(ctx context.Context, q domain.InsertEventQuery) (domain.InsertEventResult, error) {
	var res insertResult

	metadata, err := json.Marshal(q.Metadata)
	if err != nil {
		return domain.InsertEventResult{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	sqlQuery := fmt.Sprintf(`
		INSERT INTO %s (level_uuid, client_time, server_time, metadata) 
		VALUES ($1, $2, now(), $3)
		RETURNING uuid, server_time
	`, eventTableMap[q.Event])

	err = s.db.Get(&res, sqlQuery, q.UUID, q.ClientTime, metadata)
	if err != nil {
		return domain.InsertEventResult{}, fmt.Errorf("failed to insert event: %v", err)
	}

	return domain.InsertEventResult{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	}, nil
}
