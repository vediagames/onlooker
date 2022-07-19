package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	leveldomain "github.com/vediagames/onlooker/domain/level"
)

// CreateLevel godoc
// @Summary  Creates level object
// @Produce  json
// @Tags     level
// @Accept   json
// @Param    body  body      createLevelRequest  true  "Create level"
// @Success  200   {object}  createLevelResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level [post]
func (c controller) CreateLevel(ctx *gin.Context) {
	var req createLevelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res, err := c.levelService.Create(ctx, leveldomain.CreateRequest{
		SessionUUID: req.SessionUUID,
		Level:       req.Level,
		ClientTime:  req.ClientTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createLevelResponse{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	})
}

type createLevelRequest struct {
	SessionUUID string    `json:"session_uuid"`
	Level       int       `json:"level"`
	ClientTime  time.Time `json:"client_time"`
}

type createLevelResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}

// HandleEventDeath godoc
// @Summary  Logs death of player in level
// @Produce  json
// @Tags     level
// @Accept   json
// @Param    body  body      handleEventDeathRequest  true  "Log death"
// @Success  200   {object}  handleEventDeathResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/event/death [post]
func (c controller) HandleEventDeath(ctx *gin.Context) {
	var req handleEventDeathRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res, err := c.levelService.LogDeath(ctx, leveldomain.LogDeathRequest{
		UUID:       req.UUID,
		ClientTime: req.ClientTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, handleEventDeathResponse{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	})
}

type handleEventDeathRequest struct {
	UUID       string    `json:"uuid"`
	ClientTime time.Time `json:"client_time"`
}

type handleEventDeathResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}

// HandleEventComplete godoc
// @Summary  Logs completion of level
// @Produce  json
// @Tags     level
// @Accept   json
// @Param    body  body      handleEventCompleteRequest  true  "Log completion"
// @Success  200   {object}  handleEventCompleteResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/event/complete [post]
func (c controller) HandleEventComplete(ctx *gin.Context) {
	var req handleEventCompleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res, err := c.levelService.LogComplete(ctx, leveldomain.LogCompleteRequest{
		UUID:        req.UUID,
		ClientTime:  req.ClientTime,
		Achievement: leveldomain.Achievement(req.Achievement),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, handleEventCompleteResponse{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	})
}

type handleEventCompleteRequest struct {
	UUID        string    `json:"uuid"`
	ClientTime  time.Time `json:"client_time"`
	Achievement string    `json:"achievement"`
}

type handleEventCompleteResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}

// HandleEventUseGrapplingHook godoc
// @Summary  Logs usage of grappling hook
// @Produce  json
// @Tags     level
// @Accept   json
// @Param    body  body      handleEventUseGrapplingHookRequest  true  "Log grappling hook usage"
// @Success  200   {object}  handleEventUseGrapplingHookResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/event/grappling-hook-usage [post]
func (c controller) HandleEventUseGrapplingHook(ctx *gin.Context) {
	var req handleEventUseGrapplingHookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res, err := c.levelService.LogGrapplingHookUsage(ctx, leveldomain.LogGrapplingHookUsageRequest{
		UUID:       req.UUID,
		ClientTime: req.ClientTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, handleEventUseGrapplingHookResponse{
		UUID:       res.UUID,
		ServerTime: res.ServerTime,
	})
}

type handleEventUseGrapplingHookRequest struct {
	UUID       string    `json:"uuid"`
	ClientTime time.Time `json:"client_time"`
}

type handleEventUseGrapplingHookResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}
