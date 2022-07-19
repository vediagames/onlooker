package service

import (
	"context"

	leveldomain "github.com/vediagames/onlooker/domain/level"
)

type mock struct{}

func NewMock() leveldomain.Service {
	return &mock{}
}

func (m mock) Create(ctx context.Context, request leveldomain.CreateRequest) (leveldomain.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) LogDeath(ctx context.Context, request leveldomain.LogDeathRequest) (leveldomain.LogDeathResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) LogComplete(ctx context.Context, request leveldomain.LogCompleteRequest) (leveldomain.LogCompleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) LogGrapplingHookUsage(ctx context.Context, request leveldomain.LogGrapplingHookUsageRequest) (leveldomain.LogGrapplingHookUsageResponse, error) {
	//TODO implement me
	panic("implement me")
}
