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

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	viper.AutomaticEnv()

	viper.SetDefault("PORT", "8080")

	port := viper.GetString("PORT")

	viper.SetEnvPrefix("ONLOOKER")

	apiToken := viper.GetString("API_TOKEN")

	if !viper.IsSet("PSQL_CONNECTION_STRING") {
		logger.Fatal().Msg("PSQL_CONNECTION_STRING is not set")
	}

	psqlConnString := viper.GetString("PSQL_CONNECTION_STRING")

	levelStore, err := levelpostgresql.New(levelpostgresql.Config{
		ConnectionString: psqlConnString,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create level store: %s", err)
	}

	levelService, err := levelservice.New(levelservice.Config{
		Store: levelStore,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create level service: %s", err)
	}

	sessionStore, err := sessionpostgresql.New(sessionpostgresql.Config{
		ConnectionString: psqlConnString,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create session store: %s", err)
	}

	sessionService, err := sessionservice.New(sessionservice.Config{
		Store: sessionStore,
	})
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to create session service: %s", err)
	}

	c := controller.New(controller.Config{
		LevelService:   levelService,
		SessionService: sessionService,
	})

	r := gin.New()
	r.Use(loggerMiddleware(logger))
	r.Use(corsMiddleware())
	r.Use(authMiddleware(apiToken))
	r.Use(gin.Recovery())

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
		Str("port", port).
		Msgf("starting server on port %s", port)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Fatal().Err(err).Msgf("failed to run the server: %w", err)
	}
}

func authMiddleware(apiToken string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithError(401, fmt.Errorf("missing authorization header"))
			return
		}

		if auth != fmt.Sprintf("Bearer %s", apiToken) {
			ctx.AbortWithError(403, fmt.Errorf("invalid token"))
			return
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func loggerMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		l := logger.With().
			Str("method", ctx.Request.Method).
			Str("url", ctx.Request.RequestURI).
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
	}
}
