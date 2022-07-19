package service

import (
	"context"
	"fmt"

	domain "github.com/vediagames/onlooker/domain/session"
	"github.com/vediagames/onlooker/errutil"
)

type service struct {
	store domain.Store
}

type Config struct {
	Store domain.Store
}

func (c Config) Validate() error {
	var err errutil.Error

	if c.Store == nil {
		err.Add(fmt.Errorf("store is empty"))
	}

	return err.Err()
}

func New(cfg Config) (domain.Service, error) {
	if ve := cfg.Validate(); ve != nil {
		return nil, fmt.Errorf("invalid config: %w", ve)
	}

	return &service{
		store: cfg.Store,
	}, nil
}

func (s service) Create(ctx context.Context, req domain.CreateRequest) (domain.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.CreateResponse{}, fmt.Errorf("invalid request: %w", err)
	}

	newRes, err := s.store.Insert(ctx, domain.InsertQuery(req))
	if err != nil {
		return domain.CreateResponse{}, fmt.Errorf("failed to insert: %w", err)
	}

	res := domain.CreateResponse(newRes)

	if err := res.Validate(); err != nil {
		return domain.CreateResponse{}, fmt.Errorf("invalid response: %w", err)
	}

	return res, nil
}
