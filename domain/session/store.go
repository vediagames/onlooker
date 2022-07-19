package session

import (
	"context"
	"time"
)

type Store interface {
	Insert(context.Context, InsertQuery) (InsertResult, error)
}

type InsertQuery struct {
	ClientTime time.Time
	IP         string
	URL        string
	Timezone   string
	Metadata   map[string]interface{}
}

type InsertResult struct {
	UUID       string
	ServerTime time.Time
}
