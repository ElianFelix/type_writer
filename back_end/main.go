package main

import (
	"fmt"
	"net/http"
	"os"
	"type_writer_api/controllers"
	"type_writer_api/providers/activities"
	"type_writer_api/providers/scores"
	"type_writer_api/providers/texts"
	"type_writer_api/providers/users"
	"type_writer_api/services/activites"
	"type_writer_api/services/scores"
	"type_writer_api/services/texts"
	"type_writer_api/services/users"

	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	API_PORT := os.Getenv("API_PORT")

	// Create a slog logger, which:
	//   - Logs to stdout.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e := echo.New()

	// Middleware
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=type_writer-db-1 user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York", DB_USER, DB_PASS, DB_NAME, DB_PORT),
	}), &gorm.Config{})

	if err != nil {
		e.Logger.Fatal("Error initializing DB")
		e.Close()
	}

	// Providers
	usersProvider := users_provider.NewUsersProvider(db)
	textsProvider := texts_provider.NewTextsProvider(db)
	activitiesProvider := activities_provider.NewActivitiesProvider(db)
	scoresProvider := scores_provider.NewScoresProvider(db)

	// Services
	usersService := users_service.NewUsersService(usersProvider)
	textsService := texts_service.NewTextsService(textsProvider)
	activitiesService := activities_service.NewActivitiesService(activitiesProvider)
	scoresService := scores_service.NewScoresService(scoresProvider)

	// Controllers
	userController := controllers.NewUsersController(usersService)
	textController := controllers.NewTextsController(textsService)
	activityController := controllers.NewActivitiesController(activitiesService)
	scoreController := controllers.NewScoresController(scoresService)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	// User routes
	e.GET("/users", userController.GetUsers)
	e.GET("/users/:user_id", userController.GetUser)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:user_id", userController.UpdateUser)
	e.DELETE("/users/:user_id", userController.DeleteUser)

	// Text routes
	e.GET("/texts", textController.GetTexts)
	e.GET("/texts/:text_id", textController.GetText)
	e.POST("/texts", textController.CreateText)
	e.PUT("/texts/:text_id", textController.UpdateText)
	e.DELETE("/texts/:text_id", textController.DeleteText)

	// Activity routes
	e.GET("/activities", activityController.GetActivities)
	e.GET("/activities/:activity_id", activityController.GetActivity)
	e.POST("/activities", activityController.CreateActivity)
	e.PUT("/activities/:activity_id", activityController.UpdateActivity)
	e.DELETE("/activities/:activity_id", activityController.DeleteActivity)

	// Score routes
	e.GET("/scores", scoreController.GetScores)
	e.GET("/scores/:score_id", scoreController.GetScore)
	e.POST("/scores", scoreController.CreateScore)
	e.PUT("/scores/:score_id", scoreController.UpdateScore)
	e.DELETE("/scores/:score_id", scoreController.DeleteScore)

	// Server start
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", API_PORT)))
}
