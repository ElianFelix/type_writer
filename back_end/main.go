package main

import (
	"fmt"
	"net/http"
	"os"
	"type_writer_api/controllers"
	"type_writer_api/providers"
	"type_writer_api/services"

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
		DSN: fmt.Sprintf("host=back_end-db-1 user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York", DB_USER, DB_PASS, DB_NAME, DB_PORT),
	}), &gorm.Config{})

	if err != nil {
		e.Logger.Fatal("Error initializing DB")
		e.Close()
	}

	// Providers
	userProvider := providers.NewUserProvider(db)

	// Services
	userService := services.NewUserService(userProvider)

	// Controllers
	userController := controllers.NewUserController(userService)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	// User routes
	e.GET("/users", userController.GetUsers)

	// Server start
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", API_PORT)))
}
