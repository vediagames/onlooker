package postgresql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	domain "github.com/vediagames/onlooker/domain/session"
	"github.com/vediagames/onlooker/errutil"
)

type store struct {
	db *sqlx.DB
}

type Config struct {
	DB *sqlx.DB
}

func (c Config) Validate() error {
	var err errutil.Error

	if c.DB == nil {
		err.Add(fmt.Errorf("DB is empty"))
	}

	return err.Err()
}

func New(cfg Config) (domain.Store, error) {
	if ve := cfg.Validate(); ve != nil {
		return nil, fmt.Errorf("invalid config: %w", ve)
	}

	return &store{
		db: cfg.DB,
	}, nil
}

func (s store) Insert(ctx context.Context, query domain.InsertQuery) (domain.InsertResult, error) {
	//TODO implement me
	panic("implement me")
}
