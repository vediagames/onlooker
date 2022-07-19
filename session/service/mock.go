package service

import (
	"context"

	sessiondomain "github.com/vediagames/onlooker/domain/session"
)

type mock struct{}

func NewMock() sessiondomain.Service {
	return &mock{}
}

func (m mock) Create(ctx context.Context, request sessiondomain.CreateRequest) (sessiondomain.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}
