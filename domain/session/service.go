package session

import (
	"context"
	"fmt"
	"time"

	"github.com/vediagames/onlooker/errutil"
)

type Service interface {
	Create(context.Context, CreateRequest) (CreateResponse, error)
}

type CreateRequest struct {
	ClientTime time.Time
	IP         string
	URL        string
	Timezone   string
}

func (r CreateRequest) Validate() error {
	var err errutil.Error

	if r.ClientTime.IsZero() {
		err.Add(fmt.Errorf("client time must be set"))
	}

	if r.IP == "" {
		err.Add(fmt.Errorf("ip must be set"))
	}

	if r.URL == "" {
		err.Add(fmt.Errorf("url must be set"))
	}

	if r.Timezone == "" {
		err.Add(fmt.Errorf("timezone must be set"))
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
