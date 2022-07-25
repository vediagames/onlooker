package service

import (
	"context"
	"fmt"

	domain "github.com/vediagames/onlooker/domain/level"
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

func (s service) LogDeath(ctx context.Context, req domain.LogDeathRequest) (domain.LogDeathResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("invalid request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:       req.UUID,
		Event:      domain.EventDeath,
		ClientTime: req.ClientTime,
	})
	if err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogDeathResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("invalid response: %w", err)
	}

	return res, nil
}

func (s service) LogComplete(ctx context.Context, req domain.LogCompleteRequest) (domain.LogCompleteResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("invalid request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:       req.UUID,
		Event:      domain.EventComplete,
		ClientTime: req.ClientTime,
		Metadata: map[string]interface{}{
			"achievement":     req.Achievement,
			"completion_time": req.CompletionTime,
		},
	})
	if err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogCompleteResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("invalid response: %w", err)
	}

	return res, nil
}

func (s service) LogGrapplingHookUsage(ctx context.Context, req domain.LogGrapplingHookUsageRequest) (domain.LogGrapplingHookUsageResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("invalid request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:       req.UUID,
		Event:      domain.EventGrapplingHookUsage,
		ClientTime: req.ClientTime,
	})
	if err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogGrapplingHookUsageResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("invalid response: %w", err)
	}

	return res, nil
}
