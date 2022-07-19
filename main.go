package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vediagames/onlooker/controller"
	_ "github.com/vediagames/onlooker/docs"
	levelservice "github.com/vediagames/onlooker/level/service"
	sessionservice "github.com/vediagames/onlooker/session/service"
)

// @title        Onlooker Rest API
// @version      0.1.0
// @description  Lorem something lol. Just todo.

// @contact.name   Vedia Games
// @contact.url    https://vediagames.com/contact
// @contact.email  info@vediagames.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description                 Token to access the API.

const (
	defaultPort = "8080"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	viper.AutomaticEnv()

	r := gin.New()

	if err := r.SetTrustedProxies(nil); err != nil {
		logger.Fatal().Err(err).Msgf("failed to set trusted proxies: %w", err)
	}

	r.Use(gin.Recovery())
	r.Use(func(ctx *gin.Context) {
		start := time.Now()

		l := logger.With().
			Str("method", ctx.Request.Method).
			Str("client_ip", ctx.ClientIP()).
			Logger()

		ctx.Request = ctx.Request.WithContext(
			l.WithContext(ctx.Request.Context()),
		)

		ctx.Next()

		l = l.With().
			Int("code", ctx.Writer.Status()).
			Int("size", ctx.Writer.Size()).
			Logger()

		if ctx.Err() != nil {
			l.Error().
				Err(ctx.Err()).
				Msgf("failed request: %w", ctx.Err())
		}

		l.Info().TimeDiff("latency", time.Now(), start).Msg("finished request")
	})

	//levelService := levelservice.New(levelservice.Config{})
	//
	//sessionService := sessionservice.New(sessionservice.Config{})

	c := controller.New(controller.Config{
		LevelService:   levelservice.NewMock(),
		SessionService: sessionservice.NewMock(),
	})

	v1 := r.Group("/api/v1")

	v1.GET("/hello", c.Hello)

	session := v1.Group("/session")
	session.POST("/", c.CreateSession)

	level := v1.Group("/level")
	level.POST("/", c.CreateLevel)

	levelEvent := level.Group("/event")
	levelEvent.POST("/death", c.HandleEventDeath)
	levelEvent.POST("/complete", c.HandleEventComplete)
	levelEvent.POST("/grappling-hook-usage", c.HandleEventUseGrapplingHook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info().
		Str("protocol", "http").
		Str("port", defaultPort).
		Msgf("starting server on port %s", defaultPort)

	if err := r.Run(fmt.Sprintf(":%s", defaultPort)); err != nil {
		logger.Fatal().Err(err).Msgf("failed to run the server: %w", err)
	}
}
