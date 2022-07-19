package store

import (
	"context"

	domain "github.com/vediagames/onlooker/domain/level"
)

type mock struct{}

func NewMock() domain.Store {
	return &mock{}
}

func (s mock) Insert(ctx context.Context, q domain.InsertQuery) (domain.InsertResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s mock) InsertEvent(ctx context.Context, q domain.InsertEventQuery) (domain.InsertEventResult, error) {
	//TODO implement me
	panic("implement me")
}
