package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vediagames/onlooker/docs"
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
	r := gin.Default()

	c := NewController(ControllerConfig{})

	v1 := r.Group("/api/v1")

	v1.GET("/hello", c.Hello)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%s", defaultPort))
}
