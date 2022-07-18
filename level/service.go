package level

import (
	"context"
	"fmt"

	domain "github.com/vediagames/onlooker/domain/level"
)

type service struct {
	store domain.Store
}

type Config struct {
}

func New(cfg Config) domain.Service {
	return &service{}
}

func (s service) Create(ctx context.Context, req domain.CreateRequest) (domain.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.CreateResponse{}, fmt.Errorf("failed to validate request: %w", err)
	}

	newRes, err := s.store.Insert(ctx, domain.InsertQuery(req))
	if err != nil {
		return domain.CreateResponse{}, fmt.Errorf("failed to insert: %w", err)
	}

	res := domain.CreateResponse(newRes)

	if err := res.Validate(); err != nil {
		return domain.CreateResponse{}, fmt.Errorf("failed to validate response: %w", err)
	}

	return res, nil
}

func (s service) LogDeath(ctx context.Context, req domain.LogDeathRequest) (domain.LogDeathResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("failed to validate request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:          req.UUID,
		Event:         domain.EventDeath,
		StopWatchTime: req.StopwatchTime,
		ClientTime:    req.ClientTime,
	})
	if err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogDeathResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogDeathResponse{}, fmt.Errorf("failed to validate response: %w", err)
	}

	return res, nil
}

func (s service) LogComplete(ctx context.Context, req domain.LogCompleteRequest) (domain.LogCompleteResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("failed to validate request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:          req.UUID,
		Event:         domain.EventComplete,
		StopWatchTime: req.StopwatchTime,
		ClientTime:    req.ClientTime,
	})
	if err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogCompleteResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogCompleteResponse{}, fmt.Errorf("failed to validate response: %w", err)
	}

	return res, nil
}

func (s service) LogGrapplingHookUsage(ctx context.Context, req domain.LogGrapplingHookUsageRequest) (domain.LogGrapplingHookUsageResponse, error) {
	if err := req.Validate(); err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("failed to validate request: %w", err)
	}

	insertRes, err := s.store.InsertEvent(ctx, domain.InsertEventQuery{
		UUID:          req.UUID,
		Event:         domain.EventGrapplingHookUsage,
		StopWatchTime: req.StopwatchTime,
		ClientTime:    req.ClientTime,
	})
	if err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("failed to insert event: %w", err)
	}

	res := domain.LogGrapplingHookUsageResponse(insertRes)

	if err := res.Validate(); err != nil {
		return domain.LogGrapplingHookUsageResponse{}, fmt.Errorf("failed to validate response: %w", err)
	}

	return res, nil
}
