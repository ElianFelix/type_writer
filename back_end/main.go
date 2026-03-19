package main

import (
	"fmt"
	"os"
	"type_writer_api/controllers"
	"type_writer_api/helpers"
	local_middleware "type_writer_api/middleware"
	"type_writer_api/providers/activities"
	"type_writer_api/providers/scores"
	"type_writer_api/providers/texts"
	"type_writer_api/providers/users"
	"type_writer_api/services/activites"
	"type_writer_api/services/scores"
	"type_writer_api/services/texts"
	"type_writer_api/services/users"
	"type_writer_api/structures"

	"log/slog"

	"github.com/casbin/casbin/v3"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	ENV := os.Getenv("ENV")

	// Create a slog logger, which:
	//   - Logs to stdout.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e := echo.New()

	// Middleware
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=type_writer-db-1 user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York", DB_USER, DB_PASS, DB_NAME, DB_PORT),
	}), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Error initializing DB")
		e.Close()
	}

	if ENV == "INTEGRATION" {
		e.Logger.Debug("Loading test fixtures")
		err := helpers.LoadFixturesIntoDB(db, "testing/fixtures", true)
		if err != nil {
			e.Logger.Fatal("Error loading fixtures into DB\t", err)
		}
		e.Logger.Debug("Finished loading test fixtures")
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
	authController := controllers.NewAuthController(usersService)

	// Secure route group setup
	s := e.Group("")

	ce, err := casbin.NewEnforcer("config/authorization/auth_model.conf", "config/authorization/auth_policy.csv")
	if err != nil {
		e.Logger.Fatal("Error loading authorization enfocer", err)
	}

	s.Use(
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: 
				func(ctx echo.Context) jwt.Claims {
					return new(structures.JwtCustomClaims)
				},
			SigningKey: []byte("super_secret"),
		}),
	)
	s.Use(local_middleware.CasbinMiddleware(ce))

	// API routes
	//
	// Authentication routes
	e.POST("/login", authController.Login)

	// User routes
	e.GET("/users", userController.GetUsers)
	e.GET("/users/:user_id", userController.GetUser)
	e.POST("/users", userController.CreateUser)
	// Secure routes
	s.PUT("/users/:user_id", userController.UpdateUser)
	s.DELETE("/users/:user_id", userController.DeleteUser)

	// Text routes
	e.GET("/texts", textController.GetTexts)
	e.GET("/texts/:text_id", textController.GetText)
	// Secure routes
	s.POST("/texts", textController.CreateText)
	s.PUT("/texts/:text_id", textController.UpdateText)
	s.DELETE("/texts/:text_id", textController.DeleteText)

	// Activity routes
	e.GET("/activities", activityController.GetActivities)
	e.GET("/activities/:activity_id", activityController.GetActivity)
	// Secure routes
	s.POST("/activities", activityController.CreateActivity)
	s.PUT("/activities/:activity_id", activityController.UpdateActivity)
	s.DELETE("/activities/:activity_id", activityController.DeleteActivity)

	// Score routes
	e.GET("/scores", scoreController.GetScores)
	e.GET("/scores/:score_id", scoreController.GetScore)
	// Secure routes
	s.POST("/scores", scoreController.CreateScore)
	s.PUT("/scores/:score_id", scoreController.UpdateScore)
	s.DELETE("/scores/:score_id", scoreController.DeleteScore)

	// Server start
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", API_PORT)))
}
