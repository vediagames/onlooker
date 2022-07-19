package store

import (
	"context"

	domain "github.com/vediagames/onlooker/domain/session"
)

type mock struct{}

func New() (domain.Store, error) {
	return &mock{}, nil
}

func (s mock) Insert(ctx context.Context, query domain.InsertQuery) (domain.InsertResult, error) {
	//TODO implement me
	panic("implement me")
}
