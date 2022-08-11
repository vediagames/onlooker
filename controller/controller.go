package controller

import (
	"github.com/gin-gonic/gin"
	leveldomain "github.com/vediagames/onlooker/domain/level"
	sessiondomain "github.com/vediagames/onlooker/domain/session"
)

type Controller interface {
	Hello(ctx *gin.Context)
	CreateSession(ctx *gin.Context)
	CreateLevel(ctx *gin.Context)
	HandleEventDeath(ctx *gin.Context)
	HandleEventsDeath(ctx *gin.Context)
	HandleEventComplete(ctx *gin.Context)
	HandleEventsComplete(ctx *gin.Context)
	HandleEventUseGrapplingHook(ctx *gin.Context)
	HandleEventsUseGrapplingHook(ctx *gin.Context)
}

type key string

func (k key) String() string {
	return string(k)
}

const (
	KeyRealIP = key("Real-IP")
)

type controller struct {
	levelService   leveldomain.Service
	sessionService sessiondomain.Service
}

type Config struct {
	LevelService   leveldomain.Service
	SessionService sessiondomain.Service
}

func New(cfg Config) Controller {
	return &controller{
		levelService:   cfg.LevelService,
		sessionService: cfg.SessionService,
	}
}

type httpError struct {
	Message string `json:"message" example:"status bad request"`
}

type helloResponse struct {
	Message string `json:"message" example:"Hello world!"`
}

// Hello godoc
// @Summary      Hello World
// @Description  Hello World
// @Produce      json
// @Success      200  {object}  helloResponse
// @Failure      400  {object}  httpError
// @Failure      404  {object}  httpError
// @Failure      500  {object}  httpError
// @Router       /hello [get]
func (c controller) Hello(ctx *gin.Context) {
	ctx.JSON(200, helloResponse{Message: "Hello world!"})
}
