package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vediagames/onlooker/controller"
	_ "github.com/vediagames/onlooker/docs"
	levelservice "github.com/vediagames/onlooker/level/service"
	levelpostgresql "github.com/vediagames/onlooker/level/store/postgresql"
	sessionservice "github.com/vediagames/onlooker/session/service"
	sessionpostgresql "github.com/vediagames/onlooker/session/store/postgresql"
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
	viper.SetEnvPrefix("ONLOOKER")

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

	if !viper.IsSet("PSQL_CONNECTION_STRING") {
		logger.Fatal().Msg("PSQL_CONNECTION_STRING is not set")
	}

	psqlConnString := viper.GetString("PSQL_CONNECTION_STRING")

	levelStore, err := levelpostgresql.New(levelpostgresql.Config{
		ConnectionString: psqlConnString,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create level store: %w", err)
	}

	levelService, err := levelservice.New(levelservice.Config{
		Store: levelStore,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create level service: %w", err)
	}

	sessionStore, err := sessionpostgresql.New(sessionpostgresql.Config{
		ConnectionString: psqlConnString,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create session store: %w", err)
	}

	sessionService, err := sessionservice.New(sessionservice.Config{
		Store: sessionStore,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create session service: %w", err)
	}

	c := controller.New(controller.Config{
		LevelService:   levelService,
		SessionService: sessionService,
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
