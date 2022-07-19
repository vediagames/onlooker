package level

import (
	"context"
	"fmt"
	"time"

	"github.com/vediagames/onlooker/errutil"
)

type Service interface {
	Create(context.Context, CreateRequest) (CreateResponse, error)
	LogDeath(context.Context, LogDeathRequest) (LogDeathResponse, error)
	LogComplete(context.Context, LogCompleteRequest) (LogCompleteResponse, error)
	LogGrapplingHookUsage(context.Context, LogGrapplingHookUsageRequest) (LogGrapplingHookUsageResponse, error)
}

type CreateRequest struct {
	SessionUUID string
	Level       int
	ClientTime  time.Time
	Metadata    map[string]interface{}
}

func (r CreateRequest) Validate() error {
	var err errutil.Error

	if r.SessionUUID == "" {
		err.Add(fmt.Errorf("session uuid must be set"))
	}

	if r.Level < 0 {
		err.Add(fmt.Errorf("level must be above 0"))
	}

	if r.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type CreateResponse struct {
	UUID       string
	ServerTime time.Time
}

func (r CreateResponse) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
}

type LogDeathRequest struct {
	UUID       string
	ClientTime time.Time
}

func (r LogDeathRequest) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type LogDeathResponse struct {
	UUID       string
	ServerTime time.Time
}

func (r LogDeathResponse) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
}

type LogCompleteRequest struct {
	UUID        string
	ClientTime  time.Time
	Achievement Achievement
}

func (r LogCompleteRequest) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type LogCompleteResponse struct {
	UUID       string
	ServerTime time.Time
}

func (r LogCompleteResponse) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
}

type LogGrapplingHookUsageRequest struct {
	UUID       string
	ClientTime time.Time
}

func (r LogGrapplingHookUsageRequest) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	return err.Err()
}

type LogGrapplingHookUsageResponse struct {
	UUID       string
	ServerTime time.Time
}

func (r LogGrapplingHookUsageResponse) Validate() error {
	var err errutil.Error

	if r.UUID == "" {
		err.Add(fmt.Errorf("uuid must be set"))
	}

	if r.ServerTime.IsZero() {
		err.Add(fmt.Errorf("server time must be set"))
	}

	return err.Err()
}

type Achievement string

const (
	AchievementThreeStars Achievement = "three_stars"
	AchievementTwoStars   Achievement = "two_stars"
	AchievementOneStar    Achievement = "one_star"
)
