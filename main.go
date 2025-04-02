package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucamienert/flashcards/config"
	"github.com/lucamienert/flashcards/controllers"
	"github.com/lucamienert/flashcards/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

// @title API
// @version 1.0
// @description Backend
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	config.InitDB(&cfg)

	AuthController = controllers.NewAuthController(config.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(config.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", cfg.ClientOrigin}
	corsConfig.AllowCredentials = true

	r.Use(cors.New(corsConfig))

	router := r.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(r.Run(":" + cfg.ServerPort))
}
