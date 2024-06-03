package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/nutikuli/internProject_backend/pkg/datasource"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
		log.Info("Defaulting to port", port)
	}

	// routes definition

	db := datasource.DbConnection()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	datasource.InitRoute(db, app)
	app.Static("/public/image", "./public/image")

	log.Info("Listening on port", port)
	log.Info(app.Listen(":" + port))
}
