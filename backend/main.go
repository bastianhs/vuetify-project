package main

import (
	"github.com/bastianhs/vuetify-project/backend/config"
	"github.com/bastianhs/vuetify-project/backend/handlers"
	"github.com/bastianhs/vuetify-project/backend/models"
	"github.com/bastianhs/vuetify-project/backend/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Connect to database
	config.ConnectDB()

	// Migrate the schema
	models.MigrateUsers(config.DB)

	// Initialize repositories and handlers
	userRepo := repository.NewUserRepository(config.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	// Routes
	api := e.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUser)
		api.GET("/users", userHandler.GetAllUsers)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
