package main

import (
	"github.com/gin-gonic/gin"
	leveldomain "github.com/vediagames/onlooker/domain/level"
)

type Controller interface {
	Hello(ctx *gin.Context)
}

type controller struct {
	levelService leveldomain.Service
}

type ControllerConfig struct {
	LevelService leveldomain.Service
}

func NewController(cfg ControllerConfig) Controller {
	return &controller{levelService: cfg.LevelService}
}

type httpError struct {
	Code    int    `json:"code" example:"400"`
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
