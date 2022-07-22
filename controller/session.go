package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	sessiondomain "github.com/vediagames/onlooker/domain/session"
)

// CreateSession godoc
// @Summary  Creates session object
// @Produce  json
// @Tags     session
// @Accept   json
// @Param    body  body      createSessionRequest  true  "Create session"
// @Success  200   {object}  createSessionResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /session [post]
func (c controller) CreateSession(ctx *gin.Context) {
	var req createSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	ip := ctx.ClientIP()
	if ip == "" {
		ip = "Not found"
	}

	res, err := c.sessionService.Create(ctx, sessiondomain.CreateRequest{
		ClientTime: req.ClientTime,
		IP:         ip,
		URL:        req.URL,
		Timezone:   req.Timezone,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createSessionResponse{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	})
}

type createSessionRequest struct {
	ClientTime time.Time `json:"client_time"`
	IP         string    `json:"ip"`
	URL        string    `json:"url"`
	Timezone   string    `json:"timezone"`
}

type createSessionResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}
