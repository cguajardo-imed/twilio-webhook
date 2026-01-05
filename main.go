package main

import (
	"log"

	"twh/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Define routes
	app.Get("/", controllers.Incoming)
	app.Post("/", controllers.Incoming)

	log.Fatal(app.Listen(":3000"))
}
