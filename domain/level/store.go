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

func (q InsertQuery) Validate() error {
	var err errutil.Error

	if q.SessionUUID == "" {
		err.Add(fmt.Errorf("session uuid must be set"))
	}

	if q.Level < 0 {
		err.Add(fmt.Errorf("level must be above 0"))
	}

	if q.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type InsertResult struct {
	UUID       string
	ServerTime time.Time
}

func (r InsertResult) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
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

	if q.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if ve := q.Event.Validate(); ve != nil {
		err.Add(ve)
	}

	if q.StopWatchTime.IsZero() {
		err.Add(fmt.Errorf("stop watch time must be set"))
	}

	if q.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type InsertEventResult struct {
	ID         int
	ServerTime time.Time
}

func (r InsertEventResult) Validate() error {
	var err errutil.Error

	if r.ID < 0 {
		err.Add(fmt.Errorf("id must be above 0"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
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
