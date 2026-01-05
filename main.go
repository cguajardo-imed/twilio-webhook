package main

import (
	"log"

	"twh/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	// app.Get("/", controllers.Incoming)
	app.Post("/", controllers.Incoming)

	log.Fatal(app.Listen(":3000"))
}
