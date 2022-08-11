package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	leveldomain "github.com/vediagames/onlooker/domain/level"
)

// CreateLevel godoc
// @Summary  Creates level object
// @Produce  json
// @Tags     level, create
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

	res, err := c.levelService.Create(ctx.Request.Context(), leveldomain.CreateRequest{
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
// @Tags     level, event, death
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

	res, err := c.levelService.LogDeath(ctx.Request.Context(), leveldomain.LogDeathRequest{
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
// @Tags     level, event, complete
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

	res, err := c.levelService.LogComplete(ctx.Request.Context(), leveldomain.LogCompleteRequest{
		UUID:           req.UUID,
		ClientTime:     req.ClientTime,
		Achievement:    leveldomain.Achievement(req.Achievement),
		CompletionTime: time.Duration(req.CompletionTimeSeconds),
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
	UUID                  string    `json:"uuid"`
	ClientTime            time.Time `json:"client_time"`
	Achievement           string    `json:"achievement"`
	CompletionTimeSeconds int       `json:"completion_time_seconds"`
}

type handleEventCompleteResponse struct {
	UUID       string    `json:"uuid"`
	ServerTime time.Time `json:"server_time"`
}

// HandleEventUseGrapplingHook godoc
// @Summary  Logs usage of grappling hook
// @Produce  json
// @Tags     level, grappling hook, event
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

	res, err := c.levelService.LogGrapplingHookUsage(ctx.Request.Context(), leveldomain.LogGrapplingHookUsageRequest{
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

// HandleEventsComplete godoc
// @Summary  Logs completion of level
// @Produce  json
// @Tags     level, events, complete
// @Accept   json
// @Param    body  body      handleEventsCompleteRequest  true  "Log completion"
// @Success  200   {object}  handleEventsCompleteResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/events/complete [post]
func (c controller) HandleEventsComplete(ctx *gin.Context) {
	var req handleEventsCompleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res := handleEventsCompleteResponse{
		Responses: make([]handleEventCompleteResponse, len(req.Requests)),
	}

	zerolog.Ctx(ctx.Request.Context()).Info().Msgf("inserting %d events", len(req.Requests))

	for _, r := range req.Requests {
		logRes, err := c.levelService.LogComplete(ctx.Request.Context(), leveldomain.LogCompleteRequest{
			UUID:           r.UUID,
			ClientTime:     r.ClientTime,
			Achievement:    leveldomain.Achievement(r.Achievement),
			CompletionTime: time.Duration(r.CompletionTimeSeconds),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
			return
		}

		res.Responses = append(res.Responses, handleEventCompleteResponse{
			UUID:       logRes.UUID,
			ServerTime: logRes.ServerTime,
		})
	}

	ctx.JSON(http.StatusOK, res)
}

type handleEventsCompleteRequest struct {
	Requests []handleEventCompleteRequest `json:"requests"`
}
type handleEventsCompleteResponse struct {
	Responses []handleEventCompleteResponse `json:"responses"`
}

// HandleEventsDeath godoc
// @Summary  Logs death of level
// @Produce  json
// @Tags     level, death, events
// @Accept   json
// @Param    body  body      handleEventsDeathRequest  true  "Log death"
// @Success  200   {object}  handleEventsDeathResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/events/death [post]
func (c controller) HandleEventsDeath(ctx *gin.Context) {
	var req handleEventsDeathRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res := handleEventsDeathResponse{
		Responses: make([]handleEventDeathResponse, len(req.Requests)),
	}

	zerolog.Ctx(ctx.Request.Context()).Info().Msgf("inserting %d events", len(req.Requests))

	for _, r := range req.Requests {
		logRes, err := c.levelService.LogDeath(ctx.Request.Context(), leveldomain.LogDeathRequest{
			UUID:       r.UUID,
			ClientTime: r.ClientTime,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
			return
		}

		res.Responses = append(res.Responses, handleEventDeathResponse{
			UUID:       logRes.UUID,
			ServerTime: logRes.ServerTime,
		})
	}

	ctx.JSON(http.StatusOK, res)
}

type handleEventsDeathRequest struct {
	Requests []handleEventDeathRequest `json:"requests"`
}
type handleEventsDeathResponse struct {
	Responses []handleEventDeathResponse `json:"responses"`
}

// HandleEventsUseGrapplingHook godoc
// @Summary  Logs usage of grappling hook
// @Produce  json
// @Tags     level, grappling hook, events
// @Accept   json
// @Param    body  body      handleEventsUseGrapplingHookRequest  true  "Log grappling hook usage"
// @Success  200   {object}  handleEventsUseGrapplingHookResponse
// @Failure  400   {object}  httpError
// @Failure  404   {object}  httpError
// @Failure  500   {object}  httpError
// @Router   /level/events/grappling-hook-usage [post]
func (c controller) HandleEventsUseGrapplingHook(ctx *gin.Context) {
	var req handleEventsUseGrapplingHookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
		return
	}

	res := handleEventsUseGrapplingHookResponse{
		Responses: make([]handleEventUseGrapplingHookResponse, len(req.Requests)),
	}

	zerolog.Ctx(ctx.Request.Context()).Info().Msgf("inserting %d events", len(req.Requests))

	for _, r := range req.Requests {
		logRes, err := c.levelService.LogGrapplingHookUsage(ctx.Request.Context(), leveldomain.LogGrapplingHookUsageRequest{
			UUID:       r.UUID,
			ClientTime: r.ClientTime,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
			return
		}

		res.Responses = append(res.Responses, handleEventUseGrapplingHookResponse{
			UUID:       logRes.UUID,
			ServerTime: logRes.ServerTime,
		})
	}

	ctx.JSON(http.StatusOK, res)
}

type handleEventsUseGrapplingHookRequest struct {
	Requests []handleEventUseGrapplingHookRequest `json:"requests"`
}
type handleEventsUseGrapplingHookResponse struct {
	Responses []handleEventUseGrapplingHookResponse `json:"responses"`
}
