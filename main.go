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
	})

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

//A Native Collection has not been disposed, resulting in a memory leak. Allocated from:
//Unity.Collections.NativeArray`1:.ctor(Byte[], Allocator) (at /home/bokken/buildslave/unity/build/Runtime/Export/NativeArray/NativeArray.cs:69)
//UnityEngine.Networking.UploadHandlerRaw:.ctor(Byte[]) (at /home/bokken/buildslave/unity/build/Modules/UnityWebRequest/Public/UploadHandler/UploadHandler.bindings.cs:95)
//Onlooker.Controller:post(jsoner, String) (at Assets/Scripts/OnlookerManager.cs:589)
//Onlooker.Controller:CreateLevel(CreateLevelInputDto) (at Assets/Scripts/OnlookerManager.cs:516)
//<NewLevel>d__17:MoveNext() (at Assets/Scripts/OnlookerManager.cs:82)
//System.Runtime.CompilerServices.AsyncVoidMethodBuilder:Start(<NewLevel>d__17&)
//OnlookerManager:NewLevel(Int32)
//<>c__DisplayClass24_0:<PlayLevel>b__0() (at Assets/Scripts/LevelManager.cs:118)
//CrazyGames.CrazyAds:completedAdRequest(CrazySDKEvent) (at Assets/CrazySDK/CrazyAds/Scripts/CrazyAds.cs:145)
//CrazyGames.CrazyAds:completedAdRequest() (at Assets/CrazySDK/CrazyAds/Scripts/CrazyAds.cs:137)
//CrazyGames.CrazyAds:EndSimulation() (at Assets/CrazySDK/CrazyAds/Scripts/CrazyAds.cs:131)
//CrazyGames.<InvokeRealtimeCoroutine>d__15:MoveNext() (at Assets/CrazySDK/CrazyAds/Scripts/CrazyAds.cs:93)
//UnityEngine.SetupCoroutine:InvokeMoveNext(IEnumerator, IntPtr) (at /home/bokken/buildslave/unity/build/Runtime/Export/Scripting/Coroutines.cs:17)
//
