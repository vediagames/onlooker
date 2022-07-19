package level

import (
	"context"
	"fmt"
	"time"

	"github.com/vediagames/onlooker/errutil"
)

type Store interface {
	Insert(context.Context, InsertQuery) (InsertResult, error)
	InsertEvent(context.Context, InsertEventQuery) (InsertEventResult, error)
}

type InsertQuery struct {
	SessionUUID string
	Level       int
	ClientTime  time.Time
}

type InsertResult struct {
	UUID       string
	ServerTime time.Time
}

type InsertEventQuery struct {
	UUID          string
	Event         Event
	StopWatchTime time.Time
	ClientTime    time.Time
	Metadata      map[string]interface{}
}

func (q InsertEventQuery) Validate() error {
	var err errutil.Error

	if ve := q.Event.Validate(); ve != nil {
		err.Add(ve)
	}

	return err.Err()
}

type InsertEventResult struct {
	ID         int
	ServerTime time.Time
}

type Event string

func (e Event) Validate() error {
	switch e {
	case EventDeath, EventComplete, EventGrapplingHookUsage:
		return nil
	default:
		return fmt.Errorf("invalid event: %q", e)
	}
}

const (
	EventDeath              Event = "death"
	EventComplete           Event = "complete"
	EventGrapplingHookUsage Event = "grappling_hook_usage"
)
