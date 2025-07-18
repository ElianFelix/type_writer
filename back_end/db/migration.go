package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@type_writer-db-1:%s/%s?sslmode=disable", DB_USER, DB_PASS, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal("Error starting db connection for migrations ", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Error during Up migration application ", err)
	}
	log.Print("Migrations ran properly")
}
